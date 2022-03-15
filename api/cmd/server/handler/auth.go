package handler

import (
	"context"
	_ "context"
	_ "strconv"

	"github.com/gin-gonic/gin"

	"github.com/Germanchrystan/GeekStore/api/internal/auth"
	"github.com/Germanchrystan/GeekStore/api/pkg/web"
)

type Auth struct {
	authService auth.AuthService
}

func NewAuthHandler(as auth.AuthService) *Auth {
	return &Auth{
		authService: as,
	}
}

func (a *Auth) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()
		session, err := a.authService.Login(ctx)
		if err != nil {
			c.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}

		c.JSON(200, web.NewResponse(200, session, ""))
		return
	}
}
