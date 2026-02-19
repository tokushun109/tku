package middleware

import "context"

type AuthenticatedUser struct {
	UserID       uint
	UUID         string
	Name         string
	Email        string
	IsAdmin      bool
	SessionToken string
}

type authenticatedUserContextKey struct{}

func ContextWithAuthenticatedUser(ctx context.Context, user AuthenticatedUser) context.Context {
	return context.WithValue(ctx, authenticatedUserContextKey{}, user)
}

func AuthenticatedUserFromContext(ctx context.Context) (AuthenticatedUser, bool) {
	user, ok := ctx.Value(authenticatedUserContextKey{}).(AuthenticatedUser)
	return user, ok
}
