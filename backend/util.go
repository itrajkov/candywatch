package backend

import "context"

type contextKey string

const userSessionKey = contextKey("userSession")

func GetUserSession(ctx context.Context) *UserSession {
	session, ok := ctx.Value(userSessionKey).(*UserSession)
	if !ok {
		return nil
	}
	return session
}
