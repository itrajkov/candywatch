package backend

type Room struct {
	ID    int64         `json:"id"`
	users []UserSession `json:"users"`
}

func (r *Room) addUser(user *UserSession) {
	if len(r.users) == cap(r.users) {
		r.users = make([]UserSession, 2*len(r.users))
	}
	r.users = append(r.users, *user)
}

func (r *Room) removeUser(user *UserSession) {
	user.socket.CloseNow()
	user.ID = nil
}
