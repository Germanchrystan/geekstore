package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"

	"github.com/Germanchrystan/GeekStore/api/cmd/server/handler"
	"github.com/Germanchrystan/GeekStore/api/internal/admin"
	"github.com/Germanchrystan/GeekStore/api/internal/auth"
	"github.com/Germanchrystan/GeekStore/api/internal/middleware"
)

type Router interface {
	MapRoutes()
}

type router struct {
	r  *gin.Engine
	rg *gin.RouterGroup
	db *sql.DB
	m  middleware.Middleware
}

func NewRouter(r *gin.Engine, db *sql.DB) Router {
	return &router{
		r:  r,
		db: db,
		m:  middleware.NewMiddlewareRepository(db),
	}
}

func (r *router) MapRoutes() {
	r.setGroup()

	r.authRoutes()
}

func (r *router) setGroup() {
	r.rg = r.r.Group("/api/v1")
}

func (r *router) authRoutes() {
	authRepo := auth.NewRepository(r.db)
	authService := auth.NewService(authRepo)
	authHandler := handler.NewAuthHandler(authService)

	r.rg.POST("/auth/login", authHandler.Login())
	r.rg.POST("/auth/register", authHandler.Register())
	r.rg.PATCH("/auth/activate/:id", authHandler.ActivateUser())
}

func (r *router) adminRoutes() {
	adminRepo := admin.NewRepository(r.db)
	adminService := admin.NewService(adminRepo)
	adminHandler := handler.NewAdminHandler(adminService)

	r.rg.GET("/admin/users", r.m.IsAdminUserSession(), adminHandler.GetAllUsers())

	r.rg.POST("/admin/products", r.m.IsAdminUserSession(), adminHandler.PostProduct())
	r.rg.DELETE("/admin/products/:id", r.m.IsAdminUserSession(), adminHandler.DeleteProduct())
	r.rg.PUT("/admin/products/:id", r.m.IsAdminUserSession(), adminHandler.UpdateProduct())

	r.rg.PATCH("/admin/toggle/admin/:id", r.m.IsAdminUserSession(), adminHandler.ToggleUserAdmin())
	r.rg.PATCH("/admin/toggle/ban/:id", r.m.IsAdminUserSession(), adminHandler.ToggleUserBan())
}
