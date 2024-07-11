package backend

import (
	"context"
	"log"

	"github.com/google/uuid"
	"nhooyr.io/websocket"
)

type UserSession struct {
	ID     *uuid.UUID `json:"id"`
	socket *websocket.Conn
}

func NewUserSession(uuid uuid.UUID) *UserSession {
	return &UserSession{ID: &uuid}
}

func (u *UserSession) ConnectSocket(c *websocket.Conn) {
	u.socket = c
}

func (u *UserSession) SendMessage(ctx context.Context, msg Message) {
	log.Printf("To %s: %v", u.ID, msg.payload)
	if u.socket == nil {
		log.Printf("Socket of user %s not connected", u.ID)
		return
	}
	u.socket.Write(ctx, websocket.MessageBinary, msg.payload)
}

func (u *UserSession) readSocket(room *Room) {
	log.Printf("Starting reading for user %s\n", u.ID.String())
	if room == nil {
		log.Printf("Not in room, quitting goroutine. %s\n", u.ID)
		return
	}
	if u.socket == nil {
		log.Printf("Socket of user %s not connected.\n", u.ID)
		return
	}
	defer func() {
		userInRoom := room.GetUser(*u.ID)
		if userInRoom != nil {
			log.Printf("%s leaving room %s", u.ID, room.ID)
			room.removeUser(u)
		}
		u.socket.CloseNow()
	}()

	for {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		_, payload, err := u.socket.Read(ctx)
		if err != nil {
			log.Printf("Failed reading from socket: %v", err)
			break
		}

		msg := Message{sender: *u, payload: payload}
		if msg.payload != nil && room != nil {
			log.Printf("From %s: %v\n", msg.sender.ID.String(), msg.payload)
			log.Printf("Propagating to room %s\n", room.ID.String())
			for _, user := range room.Users {
				log.Printf("Propagating to %+v..\n", msg.sender.ID.String())
				user.SendMessage(ctx, msg)
			}
		}
	}
}
