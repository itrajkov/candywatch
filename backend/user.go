package backend

import (
	"github.com/google/uuid"
	"log"
)

type UserSession struct {
	ID uuid.UUID `json:"id"`
}

func NewUser(name string) *UserSession {
	uuid, err := uuid.NewUUID()
	if err != nil {
		log.Fatal("failed to generate user uuid:", err)
	}
	return &UserSession{uuid}
}
