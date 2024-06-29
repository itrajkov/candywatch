package backend

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"nhooyr.io/websocket"
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

func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	// _ := getUserSession(r.Context())
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
	msg[0] = 10
	c.Write(ctx, websocket.MessageBinary, msg)
}
