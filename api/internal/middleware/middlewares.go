package middleware

import (
	"context"
	"database/sql"
)

type Middleware interface {
	isAdminUserSession(ctx context.Context, session string) error
	isUserSession(ctx context.Context, session string) error
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
func (m *middlewareRepository) isAdminUserSession(ctx context.Context, session string) error {
	return nil
}

//===================================================================================================//
func (m *middlewareRepository) isUserSession(ctx context.Context, session string) error {
	return nil
}

//===================================================================================================//
