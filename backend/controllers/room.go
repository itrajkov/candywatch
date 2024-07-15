package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/itrajkov/candywatch/backend"
	"github.com/itrajkov/candywatch/backend/dtos"
	"github.com/itrajkov/candywatch/backend/interfaces"
	"nhooyr.io/websocket"
)

type RoomController struct {
	interfaces.RoomsManager
}

func (rc *RoomController) HandleNewRoom(w http.ResponseWriter, r *http.Request) {
	room := rc.RoomsManager.NewRoom()
	err := json.NewEncoder(w).Encode(room)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (rc *RoomController) HandleGetRooms(w http.ResponseWriter, r *http.Request) {
	log.Println("Getting rooms..")
	rooms := rc.RoomsManager.GetRooms()
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(rooms)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (rc *RoomController) HandleGetRoom(w http.ResponseWriter, r *http.Request) {
	roomIdStr := chi.URLParam(r, "id")
	roomId, err := uuid.Parse(roomIdStr)
	if err != nil {
		http.Error(w, "Invalid room ID", http.StatusBadRequest)
		return
	}
	log.Printf("Getting room with id %d..", roomId)
	room := rc.RoomsManager.GetRoomById(roomId)

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

func (rc *RoomController) HandleJoinRoom(w http.ResponseWriter, r *http.Request) {
	roomIdStr := chi.URLParam(r, "id")
	roomId, err := uuid.Parse(roomIdStr)
	if err != nil {
		http.Error(w, "Invalid room ID", http.StatusBadRequest)
		return
	}
	user := backend.GetUserSession(r.Context())

	room, err := rc.RoomsManager.JoinRoom(*user.ID, roomId)
	if err != nil {
		if errors.Is(backend.ErrRoomNotFound, err) {
			log.Println(err)
			errorHandler(w, 404, fmt.Sprintf("Room not found."))
			return
		}
		errorHandler(w, 500, fmt.Sprintf("Unknown server error"))
		return
	}

	go user.ReadSocket()
	user.Room_ch <- room
	log.Printf("%s joined %s.\n", user.ID.String(), roomId.String())

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(room)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (rc *RoomController) HandleLeaveRoom(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	roomIdStr := chi.URLParam(r, "id")
	roomId, err := uuid.Parse(roomIdStr)
	if err != nil {
		http.Error(w, "Invalid room ID", http.StatusBadRequest)
		return
	}
	user := backend.GetUserSession(r.Context())
	err = rc.RoomsManager.LeaveRoom(*user.ID, roomId)
	user.Room_ch <- nil
	if err != nil {
		if errors.Is(backend.ErrRoomNotFound, err) {
			log.Println(err)
			errorHandler(w, 404, fmt.Sprintf("Room not found."))
			return
		}
		errorHandler(w, 500, fmt.Sprintf("Unknown server error"))
		return
	}
	err = json.NewEncoder(w).Encode(dtos.NewResponse("Left room successfully.", "ok"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (rc *RoomController) HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	c, err := websocket.Accept(w, r, &websocket.AcceptOptions{
		OriginPatterns: []string{"localhost", "localhost:5173"}, // TODO: Pass these as env var
	})

	if err != nil {
		log.Printf("Failed connecting to websocket: %v", err)
	}

	user := backend.GetUserSession(r.Context())
	user.ConnectSocket(c)
	log.Println("websocket connected!")

	room := rc.RoomsManager.GetUserRoom(*user.ID)
	if room != nil {
		go user.ReadSocket()
		user.Room_ch <- room
	} else {
		log.Printf("Not in a room, goroutine not started. %s\n", user.ID)
		return
	}

}
