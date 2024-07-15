package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/itrajkov/candywatch/backend"
	"github.com/itrajkov/candywatch/backend/controllers"
	"github.com/itrajkov/candywatch/backend/middlewares"
	"github.com/itrajkov/candywatch/backend/services"
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
	roomService := services.NewRoomService()
	roomController := controllers.RoomController{RoomsManager: roomService}

	r.Use(cors.Handler(cors.Options{AllowedOrigins: []string{"https://*", "http://*"}, AllowCredentials: true}))
	r.Use(middlewares.UserSessionMiddleware(sessionManager))
	r.Use(middleware.Logger)

	r.Route("/rooms", func(r chi.Router) {
		r.Get("/", roomController.HandleGetRooms)
		r.Post("/new", roomController.HandleNewRoom)
		r.Get("/{id}", roomController.HandleGetRoom)
		r.Post("/{id}/join", roomController.HandleJoinRoom)
		r.Post("/{id}/leave", roomController.HandleLeaveRoom)
	})
	r.HandleFunc("/", roomController.HandleWebSocket)

	log.Printf("Starting server on port %s\n", s.Addr)
	log.Fatal(s.ListenAndServe())

}
