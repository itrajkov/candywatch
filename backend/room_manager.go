package backend

import (
	"fmt"
	"log"
	"sync"

	"github.com/google/uuid"
)

type RoomManager struct {
	rooms       map[uuid.UUID]*Room
	userRoomMap map[*UserSession]*Room
	sync.RWMutex
}

func NewRoomManager() *RoomManager {
	return &RoomManager{rooms: make(map[uuid.UUID]*Room, 5), userRoomMap: make(map[*UserSession]*Room)}
}

func (rm *RoomManager) GetRoomById(room_id uuid.UUID) *Room {
	rm.RLock()
	defer rm.RUnlock()
	for _, room := range rm.rooms {
		if room.ID == room_id {
			return room
		}
	}
	return nil
}

func (rm *RoomManager) AddRoom(room *Room) {
	r := rm.GetRoomById(room.ID)
	if r != nil {
		log.Printf("Room %s already exists.\n", room.ID.String())
		return
	}

	rm.Lock()
	rm.rooms[room.ID] = room
	rm.Unlock()
}

func (rm *RoomManager) GetRooms() []*Room {
	rm.RLock()
	defer rm.RUnlock()
	rooms := make([]*Room, 0)
	for _, v := range rm.rooms {
		rooms = append(rooms, v)
	}
	return rooms
}

var ErrRoomNotFound = fmt.Errorf("No such room")

func (rm *RoomManager) JoinRoom(user *UserSession, roomId uuid.UUID) (room *Room, err error) {
	current_room := rm.GetUserRoom(user)
	room = rm.GetRoomById(roomId)

	rm.Lock()
	if current_room != nil {
		log.Println("Leaving current room...")
		current_room.removeUser(user)
		rm.userRoomMap[user] = nil
		log.Println("Current room left!")
	}

	if room == nil {
		return nil, ErrRoomNotFound
	}

	if room.GetUser(*user.ID) == nil {
		room.addUser(user)
		rm.userRoomMap[user] = room
	}
	rm.Unlock()
	return room, nil
}

func (rm *RoomManager) LeaveRoom(userID uuid.UUID, roomId uuid.UUID) error {
	room := rm.GetRoomById(roomId)
	if room == nil {
		return ErrRoomNotFound
	}

	rm.Lock()
	user := room.GetUser(userID)
	if user != nil {
		room.removeUser(user)
	}
	rm.userRoomMap[user] = nil
	rm.Unlock()
	return nil
}

func (rm *RoomManager) GetUserRoom(user *UserSession) *Room {
	rm.RLock()
	defer rm.RUnlock()
	return rm.userRoomMap[user]
}
