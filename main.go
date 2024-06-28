package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/itrajkov/candywatch/backend"
)

var rm = backend.NewRoomManager()

func main() {
	r := chi.NewRouter()
	s := &http.Server{
		Addr:           ":8080",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
		Handler:        r,
	}

	r.HandleFunc("/rooms/new", handleNewRoom)
	r.HandleFunc("/rooms", handleGetRooms)

	fmt.Printf("Starting server on port %s", s.Addr)
	log.Fatal(s.ListenAndServe())

}

func handleNewRoom(w http.ResponseWriter, r *http.Request) {
	log.Println("Creating new room..")
	room := rm.NewRoom()
	log.Println("Room created..")
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(room)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func handleGetRooms(w http.ResponseWriter, r *http.Request) {
	log.Println("Getting rooms..")
	rooms := rm.GetRooms()
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(rooms)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
