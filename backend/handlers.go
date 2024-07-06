package backend

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"nhooyr.io/websocket"
)

func errorHandler(w http.ResponseWriter, status int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(NewErrorResponse(msg))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

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

func (rm *RoomManager) HandleJoinRoom(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	roomIdStr := chi.URLParam(r, "id")
	roomId, err := strconv.ParseInt(roomIdStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid room ID", http.StatusBadRequest)
		return
	}
	user := GetUserSession(r.Context())


	log.Println("Trying to join room", roomIdStr)
	room, err := rm.JoinRoom(user, roomId)
	if err != nil {
		if errors.Is(ErrRoomNotFound, err) {
			log.Println(err)
			errorHandler(w, 404, fmt.Sprintf("Room not found."))
			return
		}
		errorHandler(w, 500, fmt.Sprintf("Unknown server error"))
		return
	}
	log.Printf("%s joined %d.\n", user.ID.String(), roomId)
	err = json.NewEncoder(w).Encode(room)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (rm *RoomManager) HandleLeaveRoom(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	roomIdStr := chi.URLParam(r, "id")
	roomId, err := strconv.ParseInt(roomIdStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid room ID", http.StatusBadRequest)
		return
	}
	// TODO: Find a way to make these two refer to the same object
	user := GetUserSession(r.Context())
	err = rm.LeaveRoom(*user.ID, roomId)
	if err != nil {
		if errors.Is(ErrRoomNotFound, err) {
			log.Println(err)
			errorHandler(w, 404, fmt.Sprintf("Room not found."))
			return
		}
		errorHandler(w, 500, fmt.Sprintf("Unknown server error"))
		return
	}
	err = json.NewEncoder(w).Encode(NewResponse("Left room successfully.", "ok"))
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

	user := GetUserSession(r.Context())
	user.ConnectSocket(c)
	room := rm.GetUserRoom(user)
	go user.readSocket(room)
}
