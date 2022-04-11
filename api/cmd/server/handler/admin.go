package handler

import (
	"context"

	"github.com/Germanchrystan/GeekStore/api/internal/admin"
	"github.com/Germanchrystan/GeekStore/api/internal/domain"
	"github.com/Germanchrystan/GeekStore/api/pkg/web"
	"github.com/gin-gonic/gin"
)

type Admin struct {
	adminService admin.AdminService
}

//===================================================================================================//

func NewAdminHandler(as admin.AdminService) *Admin {
	return &Admin{
		adminService: as,
	}
}

//===================================================================================================//

func (a *Admin) GetAllUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()
		userList, err := a.adminService.GetAllUsers(ctx)
		if err != nil {
			c.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}
		c.JSON(200, web.NewResponse(200, userList, ""))
		return
	}
}

//===================================================================================================//

func (a *Admin) PostProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()
		var bodyProduct domain.Product
		bindingErr := c.ShouldBindJSON(&bodyProduct)
		if bindingErr != nil {
			c.JSON(400, web.NewResponse(400, nil, "Request body could not be binded"))
			return
		}
		newProduct, err := a.adminService.PostProduct(ctx, bodyProduct)
		if err != nil {
			c.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}
		c.JSON(201, web.NewResponse(201, newProduct, ""))
		return
	}
}

//===================================================================================================//

func (a *Admin) DeleteProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()
		productId := c.Param("id")

		err := a.adminService.DeleteProduct(ctx, productId)
		if err != nil {
			c.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}

		c.JSON(204, web.NewResponse(204, productId, ""))
	}
}

//===================================================================================================//

func (a *Admin) UpdateProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()
		var bodyProduct domain.Product
		bindingErr := c.ShouldBindJSON(&bodyProduct)
		if bindingErr != nil {
			c.JSON(400, web.NewResponse(400, nil, "Request body could not be binded"))
			return
		}

		productUpdated, err := a.adminService.UpdateProduct(ctx, bodyProduct)
		if err != nil {
			c.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}
		c.JSON(201, web.NewResponse(201, productUpdated, err.Error()))
		return
	}
}

//===================================================================================================//

func (a *Admin) ToggleUserBan() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()
		// Parsing param ID
		user_id, _ := c.Params.Get("id")
		if user_id == "" {
			c.JSON(400, web.NewResponse(400, nil, "User Id could not be retrieved"))
			return
		}
		err := a.adminService.ToggleUserBan(ctx, user_id)
		if err != nil {
			c.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}
		c.JSON(200, web.NewResponse(200, user_id, ""))
		return
	}
}

//===================================================================================================//

func (a *Admin) ToggleUserAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()
		// Parsing param ID
		user_id, _ := c.Params.Get("id")
		if user_id == "" {
			c.JSON(400, web.NewResponse(400, nil, "User Id could not be retrieved"))
			return
		}
		err := a.adminService.ToggleUserAdmin(ctx, user_id)
		if err != nil {
			c.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}
		c.JSON(200, web.NewResponse(200, user_id, ""))
		return
	}
}

//===================================================================================================//
