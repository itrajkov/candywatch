package backend

import (
	"log"

	"github.com/google/uuid"
)

type Room struct {
	ID    uuid.UUID      `json:"id"`
	Users []*UserSession `json:"users"`
}

func NewRoom() *Room {
	roomId, err := uuid.NewUUID()
	if err != nil {
		log.Fatal("Failed generating a UUID.")
	}
	room := &Room{roomId, make([]*UserSession, 0)}
	return room
}

func (r *Room) addUser(user *UserSession) {
	// Look for an empty object to re-use
	for _, u := range r.Users {
		if u.ID == nil {
			*u = *user
			return
		}
	}

	// If we don't find an empty one, append.
	r.Users = append(r.Users, user)
}

func (r *Room) removeUser(user *UserSession) {
	for i, u := range r.Users {
		if u.ID == user.ID {
			r.Users = RemoveIndex(r.Users, i)
			return
		}
	}
}

func (r *Room) getUser(sessionID uuid.UUID) *UserSession {
	log.Printf("Users state: %v", r.Users)
	for _, user := range r.Users {
		if user.ID != nil && *user.ID == sessionID {
			return user
		}
	}
	return nil
}
