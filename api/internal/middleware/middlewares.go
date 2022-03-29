package middleware

import "context"

type Middleware interface {
	isAdminUserSession(ctx context.Context, session string) error
	isUserSession(ctx context.Context, session string) error
}

func isAdminUserSession(ctx context.Context, session string) error {
	return nil
}

func isUserSession(ctx context.Context, session string) error {
	return nil
}
