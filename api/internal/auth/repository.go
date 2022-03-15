package auth

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Germanchrystan/GeekStore/api/internal/domain"
)

type AuthRepository interface {
	Login(ctx context.Context) (domain.Session, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) AuthRepository {
	return &repository{
		db: db,
	}
}

func (r *repository) Login(ctx context.Context) (domain.Session, error) {
	return domain.Session{}, errors.New("Wrong Credentials")
}
