package auth

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Germanchrystan/GeekStore/api/internal/domain"
)

//===================================================================//
type AuthRepository interface {
	Login(ctx context.Context) (domain.Session, error)
	Register(ctx context.Context) (domain.Session, error)
	ActivateUser(ctx context.Context, id string) error
	//-----------------------------------------------------//
	BanUser(ctx context.Context, id string) error
}

//===================================================================//
type repository struct {
	db *sql.DB
}

//===================================================================//
func NewRepository(db *sql.DB) AuthRepository {
	return &repository{
		db: db,
	}
}

//===================================================================//
func (r *repository) Login(ctx context.Context) (domain.Session, error) {
	return domain.Session{}, errors.New("Wrong Credentials")
}

func (r *repository) Register(ctx context.Context) (domain.Session, error) {
	return domain.Session{}, errors.New("Couldn't Register")
}

func (r *repository) ActivateUser(ctx context.Context, id string) error {
	return nil

}

func (r *repository) BanUser(ctx context.Context, id string) error {
	return nil
}

//===================================================================//
