package auth

import (
	"context"

	"github.com/Germanchrystan/GeekStore/api/internal/domain"
)

//===========================================================//
type AuthService interface {
	Login(ctx context.Context) (domain.Session, error)
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
