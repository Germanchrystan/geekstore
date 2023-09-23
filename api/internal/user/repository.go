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

// ===================================================================================================//
type UserInterface interface {
	AddAddress(ctx context.Context, input dto.InputAddress_Dto, user_id int) (domain.Address, error)
	RemoveAddress(ctx context.Context, input dto.RemoveAddress_Dto, user_id int) (string, error)

	AddCreditCard(ctx context.Context, input dto.InputCreditCard_Dto, user_id int) (dto.DisplayCreditCard_Dto, error)
	RemoveCreditCard(ctx context.Context, input dto.RemoveCreditCard_Dto, user_id int) (string, error)

	ToggleProductWhishlist(ctx context.Context, user_id, product_id int) error

	AddProductToCart(ctx context.Context, user_id, stock_id int, quantity int, price float32) (string, error)
	RemoveProductFromCart(ctx context.Context, order_id int) error

	IncreaseProductInCart(ctx context.Context, order_id int) error
	DecreaseProductInCart(ctx context.Context, order_id int) error
}

// ===================================================================================================//
type repository struct {
	db *sql.DB
}

// ===================================================================================================//
func NewRepository(db *sql.DB) UserInterface {
	return &repository{
		db: db,
	}
}

// ===================================================================================================//
func (r *repository) AddAddress(ctx context.Context, input dto.InputAddress_Dto, user_id int) (domain.Address, error) {
	// Creating Address
	// address_id := uuid.New().String()
	newAddress := domain.Address{
		//Id:           address_id,
		Street:       input.State,
		StreetNumber: input.StreetNumber,
		State:        input.State,
		Country:      input.Country,
		Zipcode:      input.Zipcode,
	}
	addressQuery := "INSERT INTO addresses(\"street\",\"street_number\",\"state\", \"country\", \"zipcode\") values ($1,$2,$3,$4,$5)"
	stmt, err := r.db.Prepare(addressQuery)
	if err != nil {
		return domain.Address{}, errors.New("An error occured")
	}
	_, err = stmt.Exec(
		newAddress.Id,
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
	addressUserQuery := "INSERT INTO address_user(\"id\", \"user_id\", \"address_id\") VALUES($1,$2,$3)"
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

// ===================================================================================================//
func (r *repository) AddCreditCard(ctx context.Context, input dto.InputCreditCard_Dto, user_id int) (dto.DisplayCreditCard_Dto, error) {
	// Hashing all credit card data
	toBeHashedData := fmt.Sprintf("%s%s%d", input.Code, input.ExpiryDate, input.SecurityCode)
	hashedData, _ := bcrypt.GenerateFromPassword([]byte(toBeHashedData), 10)
	// Getting last numbers of the code
	lastCodeNumbersString := string(input.Code)[12:15]
	lastCodeNumbersInt, _ := strconv.Atoi(lastCodeNumbersString)
	// Creating new credit card struct
	newCreditCard := domain.CreditCard{
		// ID:              uuid.New().String(),
		UserID:          user_id,
		HashedData:      string(hashedData),
		LastCodeNumbers: lastCodeNumbersInt,
	}
	creditCardQuery := "INSERT INTO credit_cards(\"user_id\",\"hashed_data\", \"last_code_numbers\") VALUES($1,$2,$3)"

	stmt, err := r.db.Prepare(creditCardQuery)
	if err != nil {
		return dto.DisplayCreditCard_Dto{}, errors.New("An error occured")
	}
	_, err = stmt.Exec(
		newCreditCard.Id,
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

// ===================================================================================================//
func (r *repository) RemoveAddress(ctx context.Context, input dto.RemoveAddress_Dto, user_id int) (string, error) {
	addressQuery := "DELETE FROM addresses WHERE id=$1"
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

// ===================================================================================================//
func (r *repository) RemoveCreditCard(ctx context.Context, input dto.RemoveCreditCard_Dto, user_id string) (string, error) {
	query := "DELETE FROM credit_cards WHERE id=$1"
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

func (r *repository) ToggleProductWhishlist(ctx context.Context, user_id, product_id int) error {
	var whishlist_id string
	selectQuery := "SELECT _id FROM whishlists WHERE user_id=$1;"
	row := r.db.QueryRow(selectQuery, user_id)
	_ = row.Scan(&whishlist_id)

	var count int
	findProductInWhishlistQuery := "SELECT COUNT(*) FROM whishlists_products WHERE whishlist_id=$1 AND product_id=$2"
	row = r.db.QueryRow(findProductInWhishlistQuery, whishlist_id, product_id)
	_ = row.Scan(&count)

	if count == 0 {
		insertQuery := "INSERT INTO whishlists_products(\"_id\", \"whishlist_id\", \"product_id\") VALUES($1,$2,$3)"
		stmt, err := r.db.Prepare(insertQuery)
		if err != nil {
			return errors.New("An error occured while adding the product")
		}
		_, err = stmt.Exec(uuid.New().String(), whishlist_id, product_id)
		if err != nil {
			return errors.New("An error occured while adding the product")
		}
		return nil
	} else {
		deleteQuery := "DELETE FROM whishlists_products WHERE whishlist_id=$1 AND product_id=$2"
		stmt, err := r.db.Prepare(deleteQuery)
		if err != nil {
			return errors.New("An error occured while deleting the product")
		}
		_, err = stmt.Exec(whishlist_id, product_id)
		if err != nil {
			return errors.New("An error occured while deleting the product")
		}
		return nil
	}
}

// ===================================================================================================//
func (r *repository) AddProductToCart(ctx context.Context, user_id, stock_id int, quantity int, price float32) (string, error) {
	cart_id, err := r.GetOrCreateActiveCart(ctx, user_id)
	if err != nil {
		return "", errors.New("Error while retrieving the user's cart.")
	}

	// Checking if stock is already in cart
	var order_id string
	checkStockQuery := "SELECT _id FROM orders WHERE cart_id=$1 AND stock_id=$2;"
	orderRow := r.db.QueryRow(checkStockQuery, cart_id, stock_id)
	err = orderRow.Scan(&order_id)

	if err == nil && order_id != "" {
		// Increase order quantity
		increaseQuery := "UPDATE orders SET quantity = quantity + $1 WHERE _id = $2;"
		stmt, err := r.db.Prepare(increaseQuery)
		if err != nil {
			return "", errors.New("Error while increasing the order's quantity.")
		}
		res, err := stmt.Exec(quantity, order_id)
		if err != nil {
			return "", errors.New("Error while increasing the order's quantity.")
		}
		affected, err := res.RowsAffected()
		if affected != 1 {
			return "", errors.New("More than one row affected while increasing the order's quantity.")
		}
		return order_id, nil
	}

	// Add order to cart
	newOrder_id := uuid.New().String()
	insertQuery := "INSERT INTO orders (\"id\",\"stock_id\", \"cart_id\", \"quantity\", \"price\") VALUES ($1,$2,$3,$4,$5)"
	stmt, err := r.db.Prepare(insertQuery)
	if err != nil {
		return "", errors.New("Error while creating order.")
	}
	_, err = stmt.Exec(newOrder_id, stock_id, cart_id, quantity, price)
	if err != nil {
		return "", errors.New("Error while creating order.")
	}
	return newOrder_id, nil
}

// ===================================================================================================//
func (r *repository) RemoveProductFromCart(ctx context.Context, order_id int) error {
	deleteQuery := "DELETE FROM orders WHERE id=$1"
	stmt, err := r.db.Prepare(deleteQuery)
	if err != nil {
		return errors.New("Error while deleting order.")
	}
	_, err = stmt.Exec(order_id)
	if err != nil {
		return errors.New("Error while deleting order.")
	}
	return nil
}

// ===================================================================================================//
func (r *repository) IncreaseProductInCart(ctx context.Context, order_id int) error {
	// Increase order quantity
	increaseQuery := "UPDATE orders SET quantity = quantity + 1 WHERE id = $1;"
	stmt, err := r.db.Prepare(increaseQuery)
	if err != nil {
		return errors.New("Error while increasing the order's quantity.")
	}
	res, err := stmt.Exec(order_id)
	if err != nil {
		return errors.New("Error while increasing the order's quantity.")
	}
	affected, err := res.RowsAffected()
	if affected != 1 {
		return errors.New("More than one row affected while increasing the order's quantity.")
	}
	return nil
}

// ===================================================================================================//
func (r *repository) DecreaseProductInCart(ctx context.Context, order_id int) error {
	increaseQuery := "UPDATE orders SET quantity = quantity + 1 WHERE id = $1;"
	stmt, err := r.db.Prepare(increaseQuery)
	if err != nil {
		return errors.New("Error while decreasing the order's quantity.")
	}
	res, err := stmt.Exec(order_id)
	if err != nil {
		return errors.New("Error while decreasing the order's quantity.")
	}
	affected, err := res.RowsAffected()
	if affected != 1 {
		return errors.New("More than one row affected while decreasing the order's quantity.")
	}
	return nil
}

//===================================================================================================//

func (r *repository) GetOrCreateActiveCart(ctx context.Context, user_id int) (string, error) {
	query := "SELECT (\"id\") FROM carts WHERE user_id=$1 AND state=\"active\";"
	row := r.db.QueryRow(query, user_id)

	var cart_id string
	err := row.Scan(&cart_id)

	// If active cart doesn't exist for user, create a new one
	if err != nil || cart_id == "" {
		new_cart_id := uuid.New().String()
		query = "INSERT INTO carts (\"id\", \"user_id\", \"state\", \"total\") VALUES ($1,$2,$3,$4);"
		stmt, err := r.db.Prepare(query)
		if err != nil {
			return "", errors.New("There was an error while creating a new cart")
		}
		_, err = stmt.Exec(new_cart_id, user_id, "active", 0)
		if err != nil {
			return "", errors.New("There was an error while creating a new cart")
		}
		return new_cart_id, nil
	}
	return cart_id, nil
}

//===================================================================================================//
