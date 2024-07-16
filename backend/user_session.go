package backend

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/itrajkov/candywatch/backend/dtos"
	"nhooyr.io/websocket"
)

type UserSession struct {
	ID      *uuid.UUID `json:"id"`
	socket  *websocket.Conn `json:"-"`
	Room_ch chan *Room `json:"-"`
}

func NewUserSession(uuid uuid.UUID) *UserSession {
	return &UserSession{ID: &uuid, Room_ch: make(chan *Room)}
}

func (u *UserSession) ConnectSocket(c *websocket.Conn) {
	u.socket = c
}

func (u *UserSession) SendMessage(ctx context.Context, msg dtos.Message) {
	log.Printf("To %s: %v", u.ID, msg.Payload)
	if u.socket == nil {
		log.Printf("Socket of user %s not connected", u.ID)
		return
	}
	u.socket.Write(ctx, websocket.MessageBinary, msg.Payload)
}

func (u *UserSession) ReadSocket() {
	log.Printf("Starting reading for user %s.\n", u.ID.String())
	if u.socket == nil {
		log.Printf("Socket of user %s not connected.\n", u.ID)
		return
	}

	defer func() {
		u.socket.CloseNow()
	}()

	var room *Room

	log.Println("Starting read loop")
readLoop:
	for {
		log.Println("In loop")
		select {
		case r := <-u.Room_ch:
			log.Println("Got a new room")
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

				msg := dtos.Message{Sender: u.ID.String(), Payload: payload}
				if msg.Payload != nil && room != nil {
					log.Printf("From %s: %v\n", msg.Sender, msg.Payload)
					for _, user := range room.Users {
						if user.ID != u.ID {
							user.SendMessage(ctx, msg)
						}
					}
				}
			}
		}
	}
}
