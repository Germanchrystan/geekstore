package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"

	"github.com/Germanchrystan/GeekStore/api/cmd/server/handler"
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

	//r.rg.PATCH("/admin/toggle/:id", r.m.IsAdminUserSession(), authHandler.ToggleUserAdmin())
}
