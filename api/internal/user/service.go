package user

import (
	"context"

	"github.com/Germanchrystan/GeekStore/api/internal/domain"
	"github.com/Germanchrystan/GeekStore/api/internal/dto"
)

//===========================================================//
type service struct {
	repository UserInterface
}

//===========================================================//
func NewUserService(repository UserInterface) UserInterface {
	return &service{
		repository: repository,
	}
}

//===========================================================//
func (s *service) AddAddress(ctx context.Context, input dto.InputAddress_Dto, user_id string) (domain.Address, error) {
	return s.repository.AddAddress(ctx, input, user_id)
}

//===========================================================//
func (s *service) RemoveAddress(ctx context.Context, input dto.RemoveAddress_Dto, user_id string) (string, error) {
	return s.repository.RemoveAddress(ctx, input, user_id)
}

//===========================================================//
func (s *service) AddCreditCard(ctx context.Context, input dto.InputCreditCard_Dto, user_id string) (dto.DisplayCreditCard_Dto, error) {
	return s.repository.AddCreditCard(ctx, input, user_id)
}

//===========================================================//
func (s *service) RemoveCreditCard(ctx context.Context, input dto.RemoveCreditCard_Dto, user_id string) (string, error) {
	return s.repository.RemoveCreditCard(ctx, input, user_id)
}

//===========================================================//
func (s *service) ToggleProductWhishlist(ctx context.Context, user_id, product_id string) error {
	return s.repository.ToggleProductWhishlist(ctx, user_id, product_id)
}

//===========================================================//
func (s *service) AddProductToCart(ctx context.Context, user_id, stock_id string, quantity int, price float32) (string, error) {
	return s.repository.AddProductToCart(ctx, user_id, stock_id, quantity, price)
}

//===========================================================//
func (s *service) RemoveProductFromCart(ctx context.Context, order_id string) error {
	return s.repository.RemoveProductFromCart(ctx, order_id)
}

//===========================================================//
func (s *service) IncreaseProductInCart(ctx context.Context, order_id string) error {
	return s.repository.IncreaseProductInCart(ctx, order_id)
}

//===========================================================//
func (s *service) DecreaseProductInCart(ctx context.Context, order_id string) error {
	return s.repository.DecreaseProductInCart(ctx, order_id)
}

//===========================================================//
