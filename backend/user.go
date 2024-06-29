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

func NewUser() *UserSession {
	uuid, err := uuid.NewUUID()
	if err != nil {
		log.Fatal("failed to generate user uuid:", err)
	}
	return &UserSession{ID: &uuid}
}

func (u *UserSession) ConnectSocket(c *websocket.Conn) {
	u.socket = c
}

func (u *UserSession) SendMessage(ctx context.Context, msg Message) {
	log.Printf("To %s: %v", u.ID, msg.payload)
	u.socket.Write(ctx, websocket.MessageBinary, msg.payload)
}

func (u *UserSession) readSocket() {
	log.Printf("Starting reading for user %s", u.ID.String())
	defer func() { u.socket.CloseNow() }()
	for {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		_, payload, err := u.socket.Read(ctx)
		if err != nil {
			log.Printf("Failed reading from socket: %v", err)
			break
		}

		msg := Message{sender: *u, payload: payload}

		if msg.payload != nil {
			log.Printf("From %s: %v\n", msg.sender.ID.String(), msg.payload)
			msg.sender.SendMessage(ctx, msg)
		} else {
			log.Printf("From %s: connected!\n", msg.sender.ID.String())
		}
	}
}
