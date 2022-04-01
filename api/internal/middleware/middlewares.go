package middleware

import (
	"database/sql"

	"github.com/Germanchrystan/GeekStore/api/internal/domain"
	"github.com/Germanchrystan/GeekStore/api/pkg/web"
	"github.com/gin-gonic/gin"
)

type Middleware interface {
	IsAdminUserSession() gin.HandlerFunc
	IsUserSession() gin.HandlerFunc
}

//===================================================================================================//
type middlewareRepository struct {
	db *sql.DB
}

//===================================================================================================//
func NewMiddlewareRepository(db *sql.DB) Middleware {
	return &middlewareRepository{
		db: db,
	}
}

//===================================================================================================//
func (m *middlewareRepository) IsAdminUserSession() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := ctx.Request.Header.Get("session_id")

		query := "SELECT * FROM sessions WHERE _id=$1"
		row := m.db.QueryRow(query, session)

		var userID string
		err := row.Scan(&userID)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, "Something went wrong"))
			return
		}

		userQuery := "SELECT * FROM users WHERE _id=$1"
		user := domain.User{}
		row = m.db.QueryRow(userQuery, userID)
		err = row.Scan(&user.ID, &user.Username, &user.FirstName, &user.LastName, &user.Email, &user.IsActive, &user.HashedPassword, &user.IsAdmin, &user.IsBanned)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, "Something went wrong"))
			return
		}

		if !user.IsAdmin {
			ctx.JSON(400, web.NewResponse(400, nil, "This user cannot perform this action"))
			return
		}
		ctx.Writer.Header().Set("user_id", user.ID)
		ctx.Next()
	}
}

//===================================================================================================//
func (m *middlewareRepository) IsUserSession() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := ctx.Request.Header.Get("session_id")
		query := "SELECT * FROM sessions WHERE _id=$1"
		row := m.db.QueryRow(query, session)

		var userID string
		err := row.Scan(&userID)
		if err != nil {

		}

		userQuery := "SELECT * FROM users WHERE _id=$1"
		user := domain.User{}
		row = m.db.QueryRow(userQuery, userID)
		err = row.Scan(&user.ID, &user.Username, &user.FirstName, &user.LastName, &user.Email, &user.IsActive, &user.HashedPassword, &user.IsAdmin, &user.IsBanned)
		if err != nil {

		}

		if user.IsActive {
			ctx.Writer.Header().Set("user_id", user.ID)
			ctx.Next()
		}
	}
}

//===================================================================================================//
