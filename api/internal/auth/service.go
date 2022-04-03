package auth

import (
	"context"
	"errors"
	"net/mail"

	"github.com/Germanchrystan/GeekStore/api/internal/dto"
	"golang.org/x/crypto/bcrypt"
)

//===========================================================//
type AuthService interface {
	Login(ctx context.Context, loginReq dto.Login_Dto) (dto.Session_Dto, error)
	Register(ctx context.Context, registerReq dto.Register_Dto) (string, error)
	//-----------------------------------------------------//
	ActivateUser(ctx context.Context, user_id string) error
	BanUser(ctx context.Context, req dto.AdminUserAction_Dto) error
	ToggleUserAdmin(ctx context.Context, user_id string) error
	//-----------------------------------------------------//
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
//===================================================================================================//
func (s *service) Login(ctx context.Context, loginReq dto.Login_Dto) (dto.Session_Dto, error) {

	return s.repository.Login(ctx, loginReq, isEmail(loginReq.EmailOrUsername))
}

//===================================================================================================//
func (s *service) Register(ctx context.Context, registerDto dto.Register_Dto) (string, error) {
	// Checking email is correct
	isEmailCorrect := isEmail(registerDto.Email)
	if !isEmailCorrect {
		return "", errors.New("E-mail is incorrect")
	}

	// Checking user is unique
	isUserUnique := s.repository.IsUserUnique(ctx, registerDto.Email, registerDto.Username)
	if isUserUnique != nil {
		return "", isUserUnique
	}

	// Hashing password
	hashed_password, _ := bcrypt.GenerateFromPassword([]byte(registerDto.Password), 10)
	registerDto.Password = string(hashed_password)

	return s.repository.Register(ctx, registerDto)
}

//===================================================================================================//
func (s *service) ActivateUser(ctx context.Context, user_id string) error {
	return s.repository.ActivateUser(ctx, user_id)
}

//===================================================================================================//
func (s *service) BanUser(ctx context.Context, req dto.AdminUserAction_Dto) error {
	return s.repository.BanUser(ctx, req)
}

//===================================================================================================//

func (s *service) ToggleUserAdmin(ctx context.Context, user_id string) error {
	return s.repository.ToggleUserAdmin(ctx, user_id)
}

//===========================================================//

func isEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
