package backend

import (
	"context"
	"log"

	"github.com/google/uuid"
	"nhooyr.io/websocket"
)

type UserSession struct {
	ID      *uuid.UUID `json:"id"`
	socket  *websocket.Conn
	room_ch chan *Room
}

func NewUserSession(uuid uuid.UUID) *UserSession {
	return &UserSession{ID: &uuid, room_ch: make(chan *Room)}
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

func (u *UserSession) readSocket() {
	log.Printf("Starting reading for user %s.\n", u.ID.String())
	if u.socket == nil {
		log.Printf("Socket of user %s not connected.\n", u.ID)
		return
	}

	defer func() {
		u.socket.CloseNow()
	}()

	var room *Room
readLoop:
	for {
		select {
		case r := <-u.room_ch:
			room = r
		default:
			{

				ctx, cancel := context.WithCancel(context.Background())
				defer cancel()
				_, payload, err := u.socket.Read(ctx)
				if err != nil {
					log.Printf("Failed reading from socket: %v", err)
					break readLoop
				}

				msg := Message{sender: *u, payload: payload}
				if msg.payload != nil && room != nil {
					log.Printf("From %s: %v\n", msg.sender.ID.String(), msg.payload)
					for _, user := range room.Users {
						if user.ID != u.ID {
							log.Printf("Propagating to %+v..\n", msg.sender.ID.String())
							user.SendMessage(ctx, msg)
							log.Println("Message sent!")
						}
					}
				}
			}
		}
	}
}
