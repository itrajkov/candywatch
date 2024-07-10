package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
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

	sessionManager := backend.NewSessionManager()
	roomManager := backend.NewRoomManager()

	r.Use(cors.Handler(cors.Options{AllowedOrigins: []string{"https://*", "http://*"}}))
	r.Use(backend.UserSessionMiddleware(sessionManager))
	r.Use(middleware.Logger)

	r.Route("/rooms", func(r chi.Router) {
		r.Get("/", roomManager.HandleGetRooms)
		r.Post("/new", roomManager.HandleNewRoom)
		r.Get("/{id}", roomManager.HandleGetRoom)
		r.Post("/{id}/join", roomManager.HandleJoinRoom)
		r.Post("/{id}/leave", roomManager.HandleLeaveRoom)
	})
	r.HandleFunc("/", roomManager.HandleWebSocket)

	log.Printf("Starting server on port %s\n", s.Addr)
	log.Fatal(s.ListenAndServe())

}
