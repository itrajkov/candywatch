package backend

type Room struct {
	ID    int64         `json:"id"`
	Users []UserSession `json:"users"`
}

func (r *Room) addUser(user *UserSession) {
	r.Users = append(r.Users, *user)
}

func (r *Room) removeUser(user *UserSession) {
	user.socket.CloseNow()
	user.ID = nil
}
