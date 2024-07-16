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

func (r *Room) RemoveUser(sessionId uuid.UUID) {
	for i, u := range r.Users {
		if *u.ID == sessionId {
			r.Users = RemoveIndex(r.Users, i)
			return
		}
	}
}

func (r *Room) GetUser(sessionId uuid.UUID) *UserSession {
	for _, u := range r.Users {
		if u.ID != nil && *u.ID == sessionId {
			return u
		}
	}
	return nil
}
