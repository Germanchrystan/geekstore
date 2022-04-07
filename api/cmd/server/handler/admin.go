package handler

import (
	"github.com/Germanchrystan/GeekStore/api/internal/admin"
	"github.com/gin-gonic/gin"
)

type Admin struct {
	adminService admin.AdminService
}

func NewAdminHandler(as admin.AdminService) *Admin {
	return &Admin{
		adminService: as,
	}
}

func (a *Admin) GetAllUsers() gin.HandlerFunc {

}

func (a *Admin) PostProduct() gin.HandlerFunc {}

func (a *Admin) DeleteProduct() gin.HandlerFunc {}

func (a *Admin) UpdateProduct() gin.HandlerFunc {}

func (a *Admin) ToggleUserBan() gin.HandlerFunc {}

func (a *Admin) ToggleUserAdmin() gin.HandlerFunc {}
