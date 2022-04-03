package admin

import (
	"context"
	"database/sql"

	"github.com/Germanchrystan/GeekStore/api/internal/domain"
)

//===================================================================================================//
type AdminRepository interface {
	GetAllUsers(ctx context.Context) ([]domain.User, error)
	//-----------------------------------------------------//
	PostProduct(ctx context.Context, product domain.Product) (domain.Product, error)
	DeleteProduct(ctx context.Context, product_id string) error
	UpdateProduct(ctx context.Context, product domain.Product) (domain.Product, error)
	//-----------------------------------------------------//
	BanUser(ctx context.Context, user_id string) error
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
func (r *repository) GetAllUsers(ctx context.Context) ([]domain.User, error)

func (r *repository) PostProduct(ctx context.Context, product domain.Product) (domain.Product, error)
func (r *repository) DeleteProduct(ctx context.Context, product_id string) error
func (r *repository) UpdateProduct(ctx context.Context, product domain.Product) (domain.Product, error)
func (r *repository) BanUser(ctx context.Context, user_id string) error
func (r *repository) ToggleUserAdmin(ctx context.Context, user_id string) error
