package handler

import (
	"context"

	"github.com/Germanchrystan/GeekStore/api/internal/dto"
	"github.com/Germanchrystan/GeekStore/api/internal/user"
	"github.com/Germanchrystan/GeekStore/api/pkg/web"
	"github.com/gin-gonic/gin"
)

type User struct {
	userService user.UserInterface
}

//===================================================================================================//
func NewUserHandler(us user.UserInterface) *User {
	return &User{
		userService: us,
	}
}

//===================================================================================================//
func (u *User) AddAddress() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()

		user_id := c.Request.Header.Get("user_id")
		if user_id == "" {
			c.JSON(400, web.NewResponse(400, nil, "User id could not be found"))
		}

		var address_input dto.InputAddress_Dto
		bindingErr := c.ShouldBindJSON(&address_input)
		if bindingErr != nil {
			c.JSON(400, web.NewResponse(400, nil, "Request body could not be binded"))
			return
		}

	}
}

//===================================================================================================//
func (u *User) RemoveAddress() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()
		var user_id string
		var address_input dto.RemoveAddress_Dto
	}
}

//===================================================================================================//
func (u *User) AddCreditCard() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()
		var user_id string
		var credit_card_input dto.InputCreditCard_Dto
	}
}

//===================================================================================================//
func (u *User) RemoveCreditCard() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()
		var user_id string
		var credit_card_input dto.RemoveCreditCard_Dto
	}
}

//===================================================================================================//
func (u *User) ToggleProductWhishlist() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()
		var user_id string
		var product_id string
	}
}

//===================================================================================================//
func (u *User) AddProductToCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()
		var user_id string
	}
}

//===================================================================================================//
func (u *User) RemoveProductFromCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()
		var user_id string
		var order dto.Order_Dto
	}
}

//===================================================================================================//
func (u *User) IncreaseProductInCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()
		var user_id string
		var order_id string
	}
}

//===================================================================================================//
func (u *User) DecreaseProductInCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()
		var user_id string
		var order_id string
	}
}

//===================================================================================================//
