package backend

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"strconv"
)

var rm = NewRoomManager()

func HandleNewRoom(w http.ResponseWriter, r *http.Request) {
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

func HandleGetRooms(w http.ResponseWriter, r *http.Request) {
	log.Println("Getting rooms..")
	rooms := rm.GetRooms()
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(rooms)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func HandleGetRoom(w http.ResponseWriter, r *http.Request) {
	roomIdStr := chi.URLParam(r, "id")
	log.Println(roomIdStr)
	roomId, err := strconv.ParseInt(roomIdStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid room ID", http.StatusBadRequest)
		return
	}
	log.Printf("Getting room with id %d..", roomId)
	room := rm.GetRoomById(roomId)

	w.Header().Set("Content-Type", "application/json")
	if room == nil {
		json.NewEncoder(w).Encode("{}")
		return
	}

	err = json.NewEncoder(w).Encode(room)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
