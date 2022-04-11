package user

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"

	"github.com/Germanchrystan/GeekStore/api/internal/domain"
	"github.com/Germanchrystan/GeekStore/api/internal/dto"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

//===================================================================================================//
type UserRepository interface {
	AddAddress(ctx context.Context, input dto.InputAddress_Dto, user_id string) (domain.Address, error)
	RemoveAddress(ctx context.Context, input dto.RemoveAddress_Dto, user_id string) (string, error)

	AddCreditCard(ctx context.Context, input dto.InputCreditCard_Dto, user_id string) (dto.DisplayCreditCard_Dto, error)
	RemoveCreditCard(ctx context.Context, input dto.RemoveCreditCard_Dto, user_id string) (string, error)

	ToggleProductWhishlist(ctx context.Context, input string) error

	AddProductToCart(ctx context.Context, user_id, stock_id string, quantity int) error
	RemoveProductFromCart(ctx context.Context, user_id, stock_id string) error

	IncreaseProductInCart(ctx context.Context, user_id, stock_id string) error
	DecreaseProductInCart(ctx context.Context, user_id, stock_id string) error
}

//===================================================================================================//
type repository struct {
	db *sql.DB
}

//===================================================================================================//
func NewRepository(db *sql.DB) UserRepository {
	return &repository{
		db: db,
	}
}

//===================================================================================================//
func (r *repository) AddAddress(ctx context.Context, input dto.InputAddress_Dto, user_id string) (domain.Address, error) {
	// Creating Address
	address_id := uuid.New().String()
	newAddress := domain.Address{
		ID:           address_id,
		Street:       input.State,
		StreetNumber: input.StreetNumber,
		State:        input.State,
		Country:      input.Country,
		Zipcode:      input.Zipcode,
	}
	addressQuery := "INSERT INTO addresses(\"_id\", \"street\",\"street_number\",\"state\", \"country\", \"zipcode\") values ($1,$2,$3,$4,$5)"
	stmt, err := r.db.Prepare(addressQuery)
	if err != nil {
		return domain.Address{}, errors.New("An error occured")
	}
	_, err = stmt.Exec(
		newAddress.ID,
		newAddress.Street,
		newAddress.StreetNumber,
		newAddress.State, input.Country,
		newAddress.Zipcode,
	)

	if err != nil {
		return domain.Address{}, errors.New("An error occured")
	}
	//Creating relation between address and user
	address_user_id := uuid.New().String()
	addressUserQuery := "INSERT INTO address_user(\"_id\", \"user_id\", \"address_id\") VALUES($1,$2,$3)"
	stmt, err = r.db.Prepare(addressUserQuery)
	if err != nil {
		return domain.Address{}, errors.New("An error occured")
	}
	_, err = stmt.Exec(address_user_id, user_id, address_id)
	if err != nil {
		return domain.Address{}, errors.New("An error occured")
	}
	return newAddress, nil
}

//===================================================================================================//
func (r *repository) AddCreditCard(ctx context.Context, input dto.InputCreditCard_Dto, user_id string) (dto.DisplayCreditCard_Dto, error) {
	// Hashing all credit card data
	toBeHashedData := fmt.Sprintf("%s%s%d", input.Code, input.ExpiryDate, input.SecurityCode)
	hashedData, _ := bcrypt.GenerateFromPassword([]byte(toBeHashedData), 10)
	// Getting last numbers of the code
	lastCodeNumbersString := string(input.Code)[12:15]
	lastCodeNumbersInt, _ := strconv.Atoi(lastCodeNumbersString)
	// Creating new credit card struct
	newCreditCard := domain.CreditCard{
		ID:              uuid.New().String(),
		UserID:          user_id,
		HashedData:      string(hashedData),
		LastCodeNumbers: lastCodeNumbersInt,
	}
	creditCardQuery := "INSERT INTO credit_cards(\"_id\",\"user_id\",\"hashed_data\", \"last_code_numbers\") VALUES($1,$2,$3)"

	stmt, err := r.db.Prepare(creditCardQuery)
	if err != nil {
		return dto.DisplayCreditCard_Dto{}, errors.New("An error occured")
	}
	_, err = stmt.Exec(
		newCreditCard.ID,
		newCreditCard.UserID,
		newCreditCard.HashedData,
		newCreditCard.LastCodeNumbers,
	)

	if err != nil {
		return dto.DisplayCreditCard_Dto{}, errors.New("An error occured")
	}
	// Returning Dto with no error
	return dto.DisplayCreditCard_Dto{
		LastCodeNumbers: lastCodeNumbersInt,
	}, nil
}

//===================================================================================================//
func (r *repository) RemoveAddress(ctx context.Context, input dto.RemoveAddress_Dto, user_id string) (string, error) {
	addressQuery := "DELETE FROM addresses WHERE _id=$1"
	stmt, err := r.db.Prepare(addressQuery)
	if err != nil {
		return "", errors.New("And error occured")
	}
	res, err := stmt.Exec(input.AddressID)
	if err != nil {
		return "", errors.New("And error occured")
	}
	affect, err := res.RowsAffected()
	if affect == 0 {
		return "", errors.New("Address not found")
	} else if affect == 1 {
		addressUserQuery := "DELETE FROM addresses_users WHERE user_id=$1 AND address_id=$2"
		stmt, err = r.db.Prepare(addressUserQuery)
		if err != nil {
			return "", errors.New("And error occured")
		}
		affect, err = res.RowsAffected()
		if affect == 1 {
			return input.AddressID, nil
		}
	}
	return "", errors.New("Something weird happenned")
}

//===================================================================================================//
func (r *repository) RemoveCreditCard(ctx context.Context, input dto.RemoveCreditCard_Dto, user_id string) (string, error) {
	query := "DELETE FROM credit_cards WHERE _id=$1"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return "", errors.New("And error occured")
	}
	res, err := stmt.Exec(input.CreditCardID)
	if err != nil {
		return "", errors.New("And error occured")
	}
	affect, err := res.RowsAffected()
	if affect == 0 {
		return "", errors.New("Address not found")
	} else if affect == 1 {
		return input.CreditCardID, nil
	}
	return "", errors.New("An error occured")
}

//===================================================================================================//

func (r *repository) ToggleProductWhishlist(ctx context.Context, input string) error {

}

//===================================================================================================//
func (r *repository) AddProductToCart(ctx context.Context, user_id, stock_id string, quantity int) error {

}

//===================================================================================================//
func (r *repository) RemoveProductFromCart(ctx context.Context, user_id, stock_id string) error {

}

//===================================================================================================//
func (r *repository) IncreaseProductInCart(ctx context.Context, user_id, stock_id string) error {

}

//===================================================================================================//
func (r *repository) DecreaseProductInCart(ctx context.Context, user_id, stock_id string) error {

}

//===================================================================================================//
