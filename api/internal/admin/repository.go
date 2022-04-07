package admin

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Germanchrystan/GeekStore/api/internal/domain"
	"github.com/google/uuid"
)

//===================================================================================================//
type AdminRepository interface {
	GetAllUsers(ctx context.Context) ([]domain.User, error)
	//-----------------------------------------------------//
	PostProduct(ctx context.Context, product domain.Product) (domain.Product, error)
	DeleteProduct(ctx context.Context, product_id string) error
	UpdateProduct(ctx context.Context, product domain.Product) (domain.Product, error)
	GetProductById(ctx context.Context, productID string) (domain.Product, error)
	//-----------------------------------------------------//
	ToggleUserBan(ctx context.Context, user_id string) error
	ToggleUserAdmin(ctx context.Context, user_id string) error
	//-----------------------------------------------------//
}

//===================================================================================================//
type repository struct {
	db *sql.DB
}

//===================================================================================================//
func NewRepository(db *sql.DB) AdminRepository {
	return &repository{
		db: db,
	}
}

//===================================================================================================//
func (r *repository) GetAllUsers(ctx context.Context) ([]domain.User, error) { // ADD PAGINATION
	query := "SELECT _id, username,firstname, lastname, email, is_active, is_admin, is_banned FROM users;"
	rows, err := r.db.Query(query)
	if err != nil {
		return []domain.User{}, errors.New("An error has ocurred")
	}
	var users []domain.User
	for rows.Next() {
		u := domain.User{}
		_ = rows.Scan(&u.ID, &u.Username, &u.FirstName, &u.LastName, &u.Email, &u.IsActive, &u.IsAdmin, &u.IsBanned)
		users = append(users, u)
	}
	return users, nil
}

//===================================================================================================//
func (r *repository) PostProduct(ctx context.Context, product domain.Product) (domain.Product, error) {
	productQuery := "INSERT into products(\"_id\", \"product_name\", \"price\",\"product_description\", \"subgenre_id\", \"category_id\") VALUES ($1, $2, $3, $4, $5, $6);"
	stmt, err := r.db.Prepare(productQuery)
	if err != nil {
		return domain.Product{}, errors.New("Unable to create product")
	}

	product.ID = uuid.New().String()
	_, err = stmt.Exec(product.ID, product.Name, product.Price, product.Description, product.Subgenre, product.Category)
	if err != nil {
		return domain.Product{}, errors.New("Unable to create product")
	}

	return product, nil
}

//===================================================================================================//
func (r *repository) DeleteProduct(ctx context.Context, product_id string) error {
	productQuery := "DELETE FROM products WHERE _id=$1;"
	stmt, err := r.db.Prepare(productQuery)
	if err != nil {
		return errors.New("Unable to delete product")
	}

	res, err := stmt.Exec(product_id)
	if err != nil {
		return errors.New("Unable to delete product")
	}
	affected, err := res.RowsAffected()
	if affected == 0 {
		return errors.New("Product could not be found")
	} else if affected == 1 {
		return nil
	} else {
		return errors.New("Something weird happenned")
	}

}

//===================================================================================================//
func (r *repository) UpdateProduct(ctx context.Context, product domain.Product) (domain.Product, error) {
	productQuery := "UPDATE products SET product_name=$1, price=$2, product_description=$3, subgenre_id=$4, category_id=$5 WHERE _id=$6;"
	stmt, err := r.db.Prepare(productQuery)
	if err != nil {
		return domain.Product{}, errors.New("Product could not be updated")
	}

	res, err := stmt.Exec(product.Name, product.Price, product.Description, product.Subgenre, product.Category, product.ID)
	if err != nil {
		return domain.Product{}, errors.New("Product could not be updated")
	}
	affected, err := res.RowsAffected()
	if affected == 0 {
		return domain.Product{}, errors.New("Product could not be found")
	} else if affected == 1 {
		return product, nil
	} else {
		return domain.Product{}, errors.New("Something odd happenned")
	}
}

//===================================================================================================//
func (r *repository) ToggleUserBan(ctx context.Context, user_id string) error {
	var is_banned bool
	selectQuery := "SELECT is_banned FROM users WHERE _id=$1;"
	row := r.db.QueryRow(selectQuery, user_id)
	_ = row.Scan(&is_banned)

	query := "UPDATE users SET is_banned=$1 WHERE _id=$2;"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return errors.New("Unable to ban user")
	}

	res, err := stmt.Exec(!is_banned, user_id)
	if err != nil {
		return errors.New("Unable to ban user")
	}
	affected, err := res.RowsAffected()
	if affected == 0 {
		return errors.New("User could not be found")
	} else if affected == 1 {
		return nil
	} else {
		return errors.New("Something weird happenned")
	}
}

//===================================================================================================//
func (r *repository) ToggleUserAdmin(ctx context.Context, user_id string) error {
	var is_admin bool
	selectQuery := "SELECT is_admin FROM users WHERE _id=$1;"
	row := r.db.QueryRow(selectQuery, user_id)
	_ = row.Scan(&is_admin)

	updateQuery := "UPDATE users SET is_admin=$1 WHERE _id=$2;"
	stmt, err := r.db.Prepare(updateQuery)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(!is_admin, user_id)
	if err != nil {
		return err
	}

	return nil
}

//===================================================================================================//
func (r *repository) GetProductById(ctx context.Context, productID string) (domain.Product, error) {
	var product domain.Product
	productQuery := "SELECT _id, product_name, price, product_description, subgenre_id, category_id FROM products WHERE _id=$1;"

	row := r.db.QueryRow(productQuery, productID)
	err := row.Scan(&product.ID, &product.Name, &product.Price, &product.Description, &product.Subgenre, &product.Category)
	if err != nil {
		return domain.Product{}, errors.New("Product could not be retrieved")
	}
	return product, nil
}
