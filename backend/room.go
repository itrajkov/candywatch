package backend

import "github.com/google/uuid"

type Room struct {
	ID    int64         `json:"id"`
	Users []*UserSession `json:"users"`
}

func (r *Room) addUser(user *UserSession) {
	r.Users = append(r.Users, user)
}

func (r *Room) removeUser(user *UserSession) {
	user.socket.CloseNow()
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
