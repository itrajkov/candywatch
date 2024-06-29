package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/itrajkov/candywatch/backend"
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
	var roomManager = backend.NewRoomManager()

	r.Use(backend.UserSessionMiddleware)
	r.HandleFunc("/rooms", roomManager.HandleGetRooms)
	r.HandleFunc("/rooms/new", roomManager.HandleNewRoom)
	r.HandleFunc("/rooms/{id}", roomManager.HandleGetRoom)
	r.HandleFunc("/", roomManager.HandleWebSocket)

	fmt.Printf("Starting server on port %s\n", s.Addr)
	log.Fatal(s.ListenAndServe())

}
