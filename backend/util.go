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

func RemoveIndex[T any](s []T, idx int)  []T {
	ret := make([]T, 0)
	ret = append(ret, s[:idx]...)
	ret = append(ret, s[idx+1:]...)
	return ret
}
