package auth

import (
	"context"

	"github.com/Germanchrystan/GeekStore/api/internal/domain"
)

type AuthRepository interface {
	Login(ctx context.Context) domain.User
}
