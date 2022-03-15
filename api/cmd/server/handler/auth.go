package handler

import (
	"context"
	_ "context"
	_ "strconv"

	"github.com/gin-gonic/gin"
)

type Auth struct {
	authService auth.Service
}

func NewAuthHandler(as auth.Service) *Auth {
	return &Auth{
		authService: as,
	}
}

func (a *Auth) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()

	}
}
