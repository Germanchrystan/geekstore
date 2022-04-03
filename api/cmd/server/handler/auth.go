package handler

import (
	"context"
	_ "strconv"

	"github.com/gin-gonic/gin"

	"github.com/Germanchrystan/GeekStore/api/internal/auth"
	"github.com/Germanchrystan/GeekStore/api/internal/dto"
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
		var loginReq dto.Login_Dto
		if bindingErr := c.ShouldBindJSON(&loginReq); bindingErr != nil {
			c.JSON(400, web.NewResponse(400, nil, "Something went wrong"))
			return
		} else {
			sessionDTO, err := a.authService.Login(ctx, loginReq)
			if err != nil {
				c.JSON(400, web.NewResponse(400, nil, err.Error()))
				return
			}
			c.JSON(200, web.NewResponse(200, sessionDTO, ""))
			return
		}
	}
}

//===================================================================//
func (a *Auth) Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()
		var registerReq dto.Register_Dto
		if bindingErr := c.ShouldBindJSON(&registerReq); bindingErr != nil {
			c.JSON(400, web.NewResponse(400, nil, "Something went wrong"))
			return
		} else {
			session, err := a.authService.Register(ctx, registerReq)
			if err != nil {
				c.JSON(400, web.NewResponse(400, nil, err.Error()))
				return
			}
			c.JSON(202, web.NewResponse(202, session, ""))
		}
	}
}

//===================================================================//
func (a *Auth) ActivateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()
		// Parsing param ID
		user_id, _ := c.Params.Get("id")
		err := a.authService.ActivateUser(ctx, user_id)
		if err != nil {
			c.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}
		c.JSON(200, web.NewResponse(200, "User activated", ""))
	}
}

//===================================================================//
func (a *Auth) BanUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()
		var req dto.AdminUserAction_Dto
		if bindingErr := c.ShouldBindJSON(&req); bindingErr != nil {
			c.JSON(400, web.NewResponse(400, nil, "Something went wrong"))
			return
		} else {
			err := a.authService.BanUser(ctx, req)
			if err != nil {
				c.JSON(400, web.NewResponse(400, nil, err.Error()))
				return
			}
			c.JSON(200, web.NewResponse(200, "User banned", ""))
		}
	}
}

//===================================================================//
func (a *Auth) ToggleUserAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()
		// Parsing param ID
		user_id, _ := c.Params.Get("id")
		err := a.authService.ToggleUserAdmin(ctx, user_id)
		if err != nil {
			c.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}
		c.JSON(200, web.NewResponse(200, "", ""))
	}
}
