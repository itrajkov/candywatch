package backend

import (
	"github.com/google/uuid"
)

type Room struct {
	ID    int64         `json:"id"`
	Users []*UserSession `json:"users"`
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
	if user.socket != nil {
		user.socket.CloseNow()
	}
	user.ID = nil
}

func (r *Room) getUser(sessionID uuid.UUID) *UserSession {
	for _, user := range r.Users {
		if user.ID != nil && *user.ID == sessionID {
			return user
		}
	}
	return nil
}
