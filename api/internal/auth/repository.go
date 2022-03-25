package auth

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/Germanchrystan/GeekStore/api/internal/domain"
	"github.com/Germanchrystan/GeekStore/api/internal/dto"
)

//===================================================================//
type AuthRepository interface {
	Login(ctx context.Context, loginReq dto.Login_Dto, isEmail bool) (dto.Session_Dto, error)
	Register(ctx context.Context, registerReq dto.Register_Dto) (string, error)
	//-----------------------------------------------------//
	ActivateUser(ctx context.Context, req dto.AdminUserAction_Dto) error
	BanUser(ctx context.Context, req dto.AdminUserAction_Dto) error
	MakeUserAdmin(ctx context.Context, req dto.AdminUserAction_Dto) error
	//-----------------------------------------------------//
}

//===================================================================//
type repository struct {
	db *sql.DB
}

//===================================================================//
func NewRepository(db *sql.DB) AuthRepository {
	return &repository{
		db: db,
	}
}

//===================================================================//
func (r *repository) Login(ctx context.Context, loginReq dto.Login_Dto, isEmail bool) (dto.Session_Dto, error) {
	// Checking if first value is the email or the username
	var firstValue string
	if isEmail {
		firstValue = "email"
	} else {
		firstValue = "username"
	}
	// Querying user
	query := fmt.Sprintf("SELECT * FROM users WHERE %s=? AND password=?", firstValue)
	row := r.db.QueryRow(query, loginReq.EmailOrUsername, loginReq.Password)

	user := domain.User{}
	err := row.Scan(&user.ID, &user.Username, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.IsActive, &user.IsAdmin, &user.IsBanned)
	if err != nil {
		return dto.Session_Dto{}, errors.New("Wrong Credentials")
	}
	if !user.IsActive {
		return dto.Session_Dto{}, errors.New("You must activate this account first")
	}
	if user.IsBanned {
		return dto.Session_Dto{}, errors.New("This user is banned")
	}
	sessionDto := dto.Session_Dto{}

	// Retrieving addresses
	var addresses []domain.Address
	addressesQuery := "SELECT * FROM addresses INNER JOIN addresses_users ON addresses_users.user_id = ?"
	aRows, err := r.db.Query(addressesQuery)
	if err == nil {
		a := domain.Address{}
		_ = aRows.Scan(&a.Street, &a.StreetNumber, &a.State, &a.Country, &a.Zipcode)
		addresses = append(addresses, a)
	}

	// Retrieving credit cards
	var creditCards []dto.DisplayCreditCard_DTO
	creditCardsQuery := "SELECT last_code_number FROM credit_cards WHERE user_id=?"
	ccRows, err := r.db.Query(creditCardsQuery)
	if err == nil {
		for ccRows.Next() {
			cc := dto.DisplayCreditCard_DTO{}
			_ = ccRows.Scan(&cc.LastCodeNumbers)
			creditCards = append(creditCards, cc)
		}
	}

	// Creating session
	newSession := domain.Session{
		ID:        uuid.New().String(),
		UserID:    user.ID,
		CreatedAt: time.Now().GoString(),
	}
	sessionQuery := "INSET INTO sessions(id, user_id, created_at) VALUES (?, ?, ?)"
	stmt, err := r.db.Prepare(sessionQuery)
	if err != nil {
		return dto.Session_Dto{}, errors.New("Unable to create session")
	}
	_, err = stmt.Exec(newSession.ID, newSession.UserID, newSession.CreatedAt)
	if err != nil {
		return dto.Session_Dto{}, errors.New("Unable to create session")
	}

	sessionDto.User = user
	sessionDto.Session = newSession
	sessionDto.Adresses = addresses
	sessionDto.CreditCards = creditCards

	return sessionDto, nil
}

func (r *repository) Register(ctx context.Context, registerDto dto.Register_Dto) (string, error) {
	id := uuid.New().String()
	query := "INSERT INTO users(id, username, firstname, lastname, email, password, is_active, is_admin, is_banned) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?) "
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return "", errors.New("There was an error while communicating to the database")
	}
	_, err = stmt.Exec(id, registerDto.Username, registerDto.FirstName, registerDto.LastName, registerDto.Email, registerDto.Password, false, true, false)
	if err != nil {
		return "", errors.New("There was an error while executing the command")
	}
	return id, errors.New("Couldn't Register")
}

func (r *repository) ActivateUser(ctx context.Context, req dto.AdminUserAction_Dto) error {
	return nil

}

func (r *repository) BanUser(ctx context.Context, req dto.AdminUserAction_Dto) error {
	return nil
}

func (r *repository) MakeUserAdmin(ctx context.Context, req dto.AdminUserAction_Dto) error {
	return nil
}

//===================================================================//
