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
		address, err := u.userService.AddAddress(ctx, address_input, user_id)

	}
}

//===================================================================================================//
func (u *User) RemoveAddress() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()

		user_id := c.Request.Header.Get("user_id")
		if user_id == "" {
			c.JSON(400, web.NewResponse(400, nil, "User id could not be found"))
		}

		var address_input dto.RemoveAddress_Dto

		addressId, err := u.userService.RemoveAddress(ctx, address_input, user_id)
	}
}

//===================================================================================================//
func (u *User) AddCreditCard() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()

		user_id := c.Request.Header.Get("user_id")
		if user_id == "" {
			c.JSON(400, web.NewResponse(400, nil, "User id could not be found"))
		}

		var credit_card_input dto.InputCreditCard_Dto

		displayCreditCard, err := u.userService.AddCreditCard(ctx, credit_card_input, user_id)

	}
}

//===================================================================================================//
func (u *User) RemoveCreditCard() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()

		user_id := c.Request.Header.Get("user_id")
		if user_id == "" {
			c.JSON(400, web.NewResponse(400, nil, "User id could not be found"))
		}

		var credit_card_input dto.RemoveCreditCard_Dto

		creditCardId, err := u.userService.RemoveCreditCard(ctx, credit_card_input, user_id)
	}
}

//===================================================================================================//
func (u *User) ToggleProductWhishlist() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()

		user_id := c.Request.Header.Get("user_id")
		if user_id == "" {
			c.JSON(400, web.NewResponse(400, nil, "User id could not be found"))
		}

		var product_id string

		err := u.userService.ToggleProductWhishlist(ctx, user_id, product_id)
	}
}

//===================================================================================================//
func (u *User) AddProductToCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()

		user_id := c.Request.Header.Get("user_id")
		if user_id == "" {
			c.JSON(400, web.NewResponse(400, nil, "User id could not be found"))
		}

		var order dto.Order_Dto

		orderId, err := u.userService.AddProductToCart(ctx, user_id, order.StockId, order.Quantity, order.Price)
		if err != nil {
			c.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}
		c.JSON(200, web.NewResponse(201, orderId, ""))
		return
	}
}

//===================================================================================================//
func (u *User) RemoveProductFromCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()
		// var user_id string
		var order_id string

		err := u.userService.RemoveProductFromCart(ctx, order_id)
		if err != nil {
			c.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}
		c.JSON(200, web.NewResponse(200, "", ""))
		return
	}
}

//===================================================================================================//
func (u *User) IncreaseProductInCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()
		// var user_id string
		var order_id string

		err := u.userService.IncreaseProductInCart(ctx, order_id)
		if err != nil {
			c.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}
		c.JSON(200, web.NewResponse(200, "", ""))
		return
	}
}

//===================================================================================================//
func (u *User) DecreaseProductInCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()
		// var user_id string
		var order_id string

		err := u.userService.DecreaseProductInCart(ctx, order_id)
		if err != nil {
			c.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}
		c.JSON(200, web.NewResponse(200, "", ""))
		return
	}
}

//===================================================================================================//
