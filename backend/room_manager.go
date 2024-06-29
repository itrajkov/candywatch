package backend

import (
	"context"
	"fmt"
	"sync"
)

type RoomManager struct {
	rooms map[int64]Room
	sync.RWMutex
}

func NewRoomManager() *RoomManager {
	return &RoomManager{rooms: make(map[int64]Room, 5)}
}

// returns a room given it's id
// if such room doesn't exist, return nil
func (rm *RoomManager) GetRoomById(id int64) *Room {
	for _, room := range rm.rooms {
		if room.ID == id {
			return &room
		}
	}
	return nil
}

// returns a room Id available for usage
func (rm *RoomManager) newRoomId() int64 {
	num_rooms := len(rm.rooms)
	for i := 0; i <= num_rooms; i++ {
		if rm.GetRoomById(int64(i)) != nil {
			continue
		}
		return int64(i)
	}
	return int64(num_rooms + 1)
}

func (rm *RoomManager) GetRooms() map[int64]Room {
	return rm.rooms
}

func (rm *RoomManager) NewRoom() *Room {
	room := &Room{rm.newRoomId(), make([]UserSession, 5)}
	rm.rooms[room.ID] = *room
	return room
}

func (rm *RoomManager) JoinRoom(user *UserSession, roomId int64) {
	if room := rm.GetRoomById(roomId); room == nil {
		fmt.Printf("Room with id %d doesn't exist", roomId)
	} else {
		rm.Lock()
		defer rm.Unlock()
		room.addUser(user)
		fmt.Printf("User %d joined room %d.", user.ID, roomId)
	}
}

type contextKey string

const userSessionKey = contextKey("userSession")

func (rm *RoomManager) getUserSession(ctx context.Context) *UserSession {
	session, ok := ctx.Value(userSessionKey).(*UserSession)
	if !ok {
		return nil
	}
	return session
}
