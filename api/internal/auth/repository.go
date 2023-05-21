package auth

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"github.com/Germanchrystan/GeekStore/api/internal/domain"
	"github.com/Germanchrystan/GeekStore/api/internal/dto"
)

// ===================================================================================================//
type AuthRepository interface {
	Login(ctx context.Context, loginReq dto.Login_Dto, isEmail bool) (dto.Session_Dto, error)
	Register(ctx context.Context, registerReq dto.Register_Dto) (string, error)
	//-----------------------------------------------------//
	ActivateUser(ctx context.Context, user_id string) error
	//-----------------------------------------------------//
	IsUserUnique(ctx context.Context, email string, username string) error
	BanUser(ctx context.Context, user_id string) error
	ToggleUserAdmin(ctx context.Context, user_id string) (bool, error)
}

// ===================================================================================================//
type repository struct {
	db *sql.DB
}

// ===================================================================================================//
func NewRepository(db *sql.DB) AuthRepository {
	return &repository{
		db: db,
	}
}

// ===================================================================================================//
func (r *repository) Login(ctx context.Context, loginReq dto.Login_Dto, isEmail bool) (dto.Session_Dto, error) {
	// Checking if first value is the email or the username
	var firstValue string
	if isEmail {
		firstValue = "email"
	} else {
		firstValue = "username"
	}

	// Querying user by first value
	query := fmt.Sprintf("SELECT * FROM users WHERE %s=$1;", firstValue)

	row := r.db.QueryRow(query, loginReq.EmailOrUsername)

	// Retrieving User by first value
	user := domain.User{}
	err := row.Scan(&user.ID, &user.Username, &user.FirstName, &user.LastName, &user.Email, &user.IsActive, &user.HashedPassword, &user.IsAdmin, &user.IsBanned)
	if err != nil {
		return dto.Session_Dto{}, errors.New("Wrong Credentials")
	}

	// Checking password
	passwordError := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(loginReq.Password))
	if passwordError != nil {
		return dto.Session_Dto{}, errors.New("Wrong Credentials")
	}

	// Checking booleans
	if !user.IsActive {
		return dto.Session_Dto{}, errors.New("You must activate this account first")
	}
	if user.IsBanned {
		return dto.Session_Dto{}, errors.New("This user is banned")
	}
	sessionDto := dto.Session_Dto{}

	//--------------------------------------------------------------------------------//
	/* All this retrieving should be worked around with concurrence once it is tested */
	// Retrieving addresses
	addresses, err := r.GetAddressesByUserId(user)
	if err != nil {
		return dto.Session_Dto{}, errors.New("Error retrieving addresses")
	}
	// Retrieving credit cards
	creditCards, err := r.GetCreditCardsByUserId(user)
	if err != nil {
		return dto.Session_Dto{}, errors.New("Error retrieving credit cards")
	}
	// Create session
	newSession, err := r.CreateSession(user)

	sessionDto.User = user
	sessionDto.Session = newSession
	sessionDto.Adresses = addresses
	sessionDto.CreditCards = creditCards

	return sessionDto, nil
}

// ===================================================================================================//
func (r *repository) GetCreditCardsByUserId(user domain.User) ([]dto.DisplayCreditCard_Dto, error) {
	// Retrieving credit cards
	var creditCards []dto.DisplayCreditCard_Dto
	creditCardsQuery := "SELECT last_code_number FROM credit_cards WHERE user_id=$1;"
	ccRows, err := r.db.Query(creditCardsQuery, user.ID)
	if err == nil {
		for ccRows.Next() {
			cc := dto.DisplayCreditCard_Dto{}
			_ = ccRows.Scan(&cc.LastCodeNumbers)
			creditCards = append(creditCards, cc)
		}
		return creditCards, nil
	}
	return nil, err
}

// ===================================================================================================//

func (r *repository) GetAddressesByUserId(user domain.User) ([]domain.Address, error) {
	var addresses []domain.Address
	addressesQuery := "SELECT * FROM addresses INNER JOIN addresses_users ON addresses_users.user_id = $1;"
	aRows, err := r.db.Query(addressesQuery, user.ID)
	if err == nil {
		a := domain.Address{}
		_ = aRows.Scan(&a.Street, &a.StreetNumber, &a.State, &a.Country, &a.Zipcode)
		addresses = append(addresses, a)
		return addresses, nil
	}
	return nil, err
}

// ===================================================================================================//
func (r *repository) CreateSession(user domain.User) (domain.Session, error) {
	// Creating session
	newSession := domain.Session{
		ID:        uuid.New().String(),
		UserID:    user.ID,
		CreatedAt: time.Now().GoString(),
	}
	sessionQuery := "INSERT INTO sessions(\"_id\", \"user_id\", \"created_at\") VALUES ($1, $2, $3);"
	stmt, err := r.db.Prepare(sessionQuery)
	if err != nil {
		return domain.Session{}, errors.New("Unable to create session")
	}
	_, err = stmt.Exec(newSession.ID, newSession.UserID, newSession.CreatedAt)
	if err != nil {
		return domain.Session{}, errors.New("Unable to create session")
	}
	return newSession, nil
}

// ===================================================================================================//
func (r *repository) Register(ctx context.Context, registerDto dto.Register_Dto) (string, error) {
	// Creating new user
	id := uuid.New().String()
	query := "INSERT INTO users(\"_id\", \"username\", \"firstname\", \"lastname\", \"email\", \"hashed_password\", \"is_active\", \"is_admin\", \"is_banned\") VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9);"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return "", errors.New("There was an error while communicating to the database")
	}
	_, err = stmt.Exec(id, registerDto.Username, registerDto.FirstName, registerDto.LastName, registerDto.Email, registerDto.Password, false, true, false)
	if err != nil {
		return "", errors.New("There was an error while executing the command")
	}
	return id, err
}

// ===================================================================================================//
func (r *repository) ActivateUser(ctx context.Context, user_id string) error {
	query := "UPDATE users SET is_active=true WHERE _id=$1;"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(user_id)
	if err != nil {
		return err
	}

	return nil

}

//===================================================================================================//

func (r *repository) IsUserUnique(ctx context.Context, email string, username string) error {
	// Check email is not repeated
	var amountUsers int

	checkUniqueEmailQuery := "SELECT COUNT(*) FROM users WHERE email = $1;"
	row := r.db.QueryRow(checkUniqueEmailQuery, email)
	_ = row.Scan(&amountUsers)
	if amountUsers > 0 {
		return errors.New("E-mail already taken")
	}

	// Check username is not repeated
	checkUniqueUsernameQuery := "SELECT COUNT(*) FROM users WHERE username = $1;"
	row = r.db.QueryRow(checkUniqueUsernameQuery, username)
	_ = row.Scan(&amountUsers)
	if amountUsers > 0 {
		return errors.New("Username already taken")
	}
	return nil
}

// ===================================================================================================//
func (r *repository) BanUser(ctx context.Context, user_id string) error {
	query := "UPDATE users SET is_banned=true WHERE _id=$1;"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(user_id)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) ToggleUserAdmin(ctx context.Context, user_id string) (bool, error) {
	var isAlreadyAdmin bool
	row := r.db.QueryRow("SELECT is_admin FROM users WHERE _id=$1;", user_id)
	err := row.Scan(&isAlreadyAdmin)
	if err != nil {
		return false, err
	}

	query := "UPDATE users SET is_admin=$1 WHERE _id=$2;"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return false, err
	}

	_, err = stmt.Exec(!isAlreadyAdmin, user_id)
	if err != nil {
		return false, err
	}

	return !isAlreadyAdmin, nil

}
