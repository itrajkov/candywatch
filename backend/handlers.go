package backend

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"nhooyr.io/websocket"
)

func (rm *RoomManager) HandleNewRoom(w http.ResponseWriter, r *http.Request) {
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

func (rm *RoomManager) HandleGetRooms(w http.ResponseWriter, r *http.Request) {
	log.Println("Getting rooms..")
	rooms := rm.GetRooms()
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(rooms)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (rm *RoomManager) HandleGetRoom(w http.ResponseWriter, r *http.Request) {
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

func (rm *RoomManager) HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	c, err := websocket.Accept(w, r, &websocket.AcceptOptions{
		OriginPatterns: []string{"localhost:5173"},
	})

	if err != nil {
		log.Printf("Failed connecting to websocket: %v", err)
	}

	user := rm.getUserSession(r.Context())
	user.ConnectSocket(c)
	go user.readSocket()
}
