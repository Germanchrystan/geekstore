package auth

import (
	"context"
	"crypto/sha1"
	"encoding/base64"
	"net/mail"
	_ "regexp"

	"github.com/Germanchrystan/GeekStore/api/internal/dto"
)

//===========================================================//
type AuthService interface {
	Login(ctx context.Context, loginReq dto.Login_Dto) (dto.Session_Dto, error)
	Register(ctx context.Context, registerReq dto.Register_Dto) (string, error)
	//-----------------------------------------------------//
	ActivateUser(ctx context.Context, req dto.AdminUserAction_Dto) error
	BanUser(ctx context.Context, req dto.AdminUserAction_Dto) error
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

func (s *service) Login(ctx context.Context, loginReq dto.Login_Dto) (dto.Session_Dto, error) {

	return s.repository.Login(ctx, loginReq, isEmail(loginReq.EmailOrUsername))
}

func (s *service) Register(ctx context.Context, registerDto dto.Register_Dto) (string, error) {
	// Checking user is unique
	isUserUnique := s.repository.IsUserUnique(ctx, registerDto.Email, registerDto.Username)
	if isUserUnique != nil {
		return "", isUserUnique
	}

	// Hashing password
	hasher := sha1.New()
	hasher.Write([]byte(registerDto.Password))
	registerDto.Password = base64.URLEncoding.EncodeToString(hasher.Sum(nil))

	return s.repository.Register(ctx, registerDto)
}

func (s *service) ActivateUser(ctx context.Context, req dto.AdminUserAction_Dto) error {
	return s.repository.ActivateUser(ctx, req)
}

func (s *service) BanUser(ctx context.Context, req dto.AdminUserAction_Dto) error {
	return s.repository.BanUser(ctx, req)
}

func (s *service) MakeUserAdmin(ctx context.Context, req dto.AdminUserAction_Dto) error {
	return s.repository.MakeUserAdmin(ctx, req)
}

//===========================================================//

func isEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
