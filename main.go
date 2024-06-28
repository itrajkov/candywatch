package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/itrajkov/candywatch/backend"
	"nhooyr.io/websocket"
)

func main() {
	r := chi.NewRouter()
	s := &http.Server{
		Addr:           ":8080",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
		Handler:        r,
	}

	r.HandleFunc("/rooms/new", backend.HandleNewRoom)
	r.HandleFunc("/rooms", backend.HandleGetRooms)
	r.HandleFunc("/rooms/{id}", backend.HandleGetRoom)
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		c, err := websocket.Accept(w, r, &websocket.AcceptOptions{
			OriginPatterns: []string{"localhost:5173"},
		})
		if err != nil {
			log.Printf("Failed connecting to websocket: %v", err)
		}
		defer c.CloseNow()

		ctx, cancel := context.WithTimeout(r.Context(), time.Second*10)
		defer cancel()

		_, msg, err := c.Read(ctx)
		if err != nil {
			log.Printf("Failed reading from socket: %v", err)
		}

		log.Printf("Message: %v", msg)
		msg = make([]byte, 5, 5)
		c.Write(ctx, websocket.MessageBinary, msg)
	})

	fmt.Printf("Starting server on port %s\n", s.Addr)
	log.Fatal(s.ListenAndServe())

}
