package auth

import (
	"context"

	"github.com/Germanchrystan/GeekStore/api/internal/domain"
)

//===========================================================//
type AuthService interface {
	Login(ctx context.Context) (domain.Session, error)
	Register(ctx context.Context) (domain.Session, error)
	ActivateUser(ctx context.Context, id string) error
}

type service struct {
	repository AuthRepository
}

//===========================================================//
func NewService(repository AuthRepository) AuthService {
	return &service{
		repository: repository,
	}
}

//===========================================================//

func (s *service) Login(ctx context.Context) (domain.Session, error) {
	return s.repository.Login(ctx)
}

func (s *service) Register(ctx context.Context) (domain.Session, error) {
	return s.repository.Register(ctx)
}

func (s *service) ActivateUser(ctx context.Context, id string) error {
	return s.repository.ActivateUser(ctx, id)
}
