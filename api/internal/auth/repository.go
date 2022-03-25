package auth

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Germanchrystan/GeekStore/api/internal/dto"
)

//===================================================================//
type AuthRepository interface {
	Login(ctx context.Context, loginReq dto.Login_Dto) (dto.Session_Dto, error)
	Register(ctx context.Context, registerReq dto.Register_Dto) (string, error)
	//-----------------------------------------------------//
	ActivateUser(ctx context.Context, id string) error
	BanUser(ctx context.Context, id string) error
	//-----------------------------------------------------//
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
func (r *repository) Login(ctx context.Context, loginReq dto.Login_Dto) (dto.Session_Dto, error) {
	return dto.Session_Dto{}, errors.New("Wrong Credentials")
}

func (r *repository) Register(ctx context.Context, registerDto dto.Register_Dto) (string, error) {
	return "", errors.New("Couldn't Register")
}

func (r *repository) ActivateUser(ctx context.Context, id string) error {
	return nil

}

func (r *repository) BanUser(ctx context.Context, id string) error {
	return nil
}

//===================================================================//
