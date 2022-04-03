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

type service struct {
	repository AdminRepository
}

//===========================================================//
func NewService(repository AdminRepository) AdminService {
	return &service{
		repository: repository,
	}
}

//===========================================================//

func (s *service) GetAllUsers(ctx context.Context) ([]domain.User, error)
func (s *service) PostProduct(ctx context.Context, product domain.Product) (domain.Product, error)
func (s *service) DeleteProduct(ctx context.Context, product_id string) error
func (s *service) UpdateProduct(ctx context.Context, product domain.Product) (domain.Product, error)
func (s *service) BanUser(ctx context.Context, user_id string) error
func (s *service) ToggleUserAdmin(ctx context.Context, user_id string) error
