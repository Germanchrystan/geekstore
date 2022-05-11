package middleware

import (
	"database/sql"

	"github.com/Germanchrystan/GeekStore/api/internal/domain"
	"github.com/gin-gonic/gin"
)

type Middleware interface {
	IsAdminUserSession() gin.HandlerFunc
	IsUserSession() gin.HandlerFunc
}

func respondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
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

		query := "SELECT user_id FROM sessions WHERE _id=$1"
		row := m.db.QueryRow(query, session)

		var userID string
		err := row.Scan(&userID)
		if err != nil {
			respondWithError(ctx, 401, "Something went wrong")
			return
		}

		userQuery := "SELECT is_admin FROM users WHERE _id=$1"
		var is_admin bool
		row = m.db.QueryRow(userQuery, userID)
		err = row.Scan(is_admin)
		if err != nil {
			respondWithError(ctx, 401, "Something went wrong")
			return
		}

		if !is_admin {
			respondWithError(ctx, 401, "This user cannot perform this action")
			return
		} else {
			ctx.Writer.Header().Set("user_id", userID)
			ctx.Next()
		}
	}
}

//===================================================================================================//
func (m *middlewareRepository) IsUserSession() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := ctx.Request.Header.Get("session_id")
		query := "SELECT user_id FROM sessions WHERE _id=$1"
		row := m.db.QueryRow(query, session)

		var userID string
		err := row.Scan(&userID)
		if err != nil {
			respondWithError(ctx, 302, "Wrong credentials")
			return
		}

		userQuery := "SELECT * FROM users WHERE _id=$1"
		user := domain.User{}
		row = m.db.QueryRow(userQuery, userID)
		err = row.Scan(&user.ID, &user.Username, &user.FirstName, &user.LastName, &user.Email, &user.IsActive, &user.HashedPassword, &user.IsAdmin, &user.IsBanned)
		if err != nil {
			respondWithError(ctx, 302, "Wrong credentials")
			return
		}

		if user.IsActive {
			ctx.Writer.Header().Set("user_id", user.ID)
			ctx.Next()
		} else {
			respondWithError(ctx, 401, "User is not active")
			return
		}
	}
}

//===================================================================================================//
