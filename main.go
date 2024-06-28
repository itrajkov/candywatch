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

	r.HandleFunc("/rooms/new", backend.HandleNewRoom)
	r.HandleFunc("/rooms", backend.HandleGetRooms)
	r.HandleFunc("/rooms/{id}", backend.HandleGetRoom)

	fmt.Printf("Starting server on port %s", s.Addr)
	log.Fatal(s.ListenAndServe())

}

