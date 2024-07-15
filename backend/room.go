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
	roomId, err := uuid.NewRandom()
	if err != nil {
		log.Fatal("Failed generating a UUID.")
	}
	room := &Room{roomId, make([]*UserSession, 0)}
	return room
}

func (r *Room) AddUser(user *UserSession) {
	// Look for an empty object to re-use
	for _, u := range r.Users {
		if u.ID == nil {
			*u = *user
			return
		}
	}
	r.Users = append(r.Users, user)
}

func (r *Room) RemoveUser(user *UserSession) {
	for i, u := range r.Users {
		if u.ID == user.ID {
			r.Users = RemoveIndex(r.Users, i)
			return
		}
	}
}

func (r *Room) GetUser(sessionID uuid.UUID) *UserSession {
	for _, user := range r.Users {
		if user.ID != nil && *user.ID == sessionID {
			return user
		}
	}
	return nil
}
