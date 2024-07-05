package backend

import (
	"fmt"

	"github.com/google/uuid"
)

type Room struct {
	ID    int64         `json:"id"`
	Users []*UserSession `json:"users"`
}

func (r *Room) addUser(user *UserSession) {
	fmt.Printf("Looking for a nulled out struct\n")
	for _, u := range r.Users {
		if u.ID == nil {
			*u = *user
			return
		}
	}

	fmt.Printf("Couldn't find an empty slot, appending..\n")
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
		if *user.ID == sessionID {
			return user
		}
	}
	return nil
}
