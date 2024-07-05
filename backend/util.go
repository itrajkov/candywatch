package backend

import "context"

type contextKey string

const UserSessionKey = contextKey("userSession")

func GetUserSession(ctx context.Context) *UserSession {
	session, ok := ctx.Value(UserSessionKey).(*UserSession)
	if !ok {
		return nil
	}
	return session
}
