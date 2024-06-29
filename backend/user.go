package backend

import (
	"context"
	"log"

	"github.com/google/uuid"
)

type UserSession struct {
	ID uuid.UUID `json:"id"`
}

func NewUser() *UserSession {
	uuid, err := uuid.NewUUID()
	if err != nil {
		log.Fatal("failed to generate user uuid:", err)
	}
	return &UserSession{uuid}
}

type contextKey string

const userSessionKey = contextKey("userSession")

func getUserSession(ctx context.Context) *UserSession {
	session, ok := ctx.Value(userSessionKey).(*UserSession)
	if !ok {
		return nil
	}
	return session
}
