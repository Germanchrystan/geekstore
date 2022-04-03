package admin

import (
	"context"

	"github.com/Germanchrystan/GeekStore/api/internal/domain"
)

type AdminService interface {
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
