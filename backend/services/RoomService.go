package services

import (
	"log"
	"sync"

	"github.com/google/uuid"
	"github.com/itrajkov/candywatch/backend"
	"github.com/itrajkov/candywatch/backend/interfaces"
)

type RoomService struct {
	interfaces.RoomsManager
	rooms       map[uuid.UUID]*backend.Room
	userRoomMap map[uuid.UUID]*backend.Room
	sync.RWMutex
}

func NewRoomService() *RoomService {
	return &RoomService{rooms: make(map[uuid.UUID]*backend.Room, 5), userRoomMap: make(map[uuid.UUID]*backend.Room)}
}

func (rs *RoomService) GetRoomById(room_id uuid.UUID) *backend.Room {
	rs.RLock()
	defer rs.RUnlock()
	for _, room := range rs.rooms {
		if room.ID == room_id {
			return room
		}
	}
	return nil
}

func (rs *RoomService) NewRoom() *backend.Room {
	room := backend.NewRoom()
	log.Printf("New room: %s", room.ID.String())
	rs.addRoom(room)
	return room
}

func (rs *RoomService) addRoom(room *backend.Room) {
	r := rs.GetRoomById(room.ID)
	if r != nil {
		log.Println("Room ID collision, try again.")
		return
	}

	rs.Lock()
	rs.rooms[room.ID] = room
	rs.Unlock()
}

func (rs *RoomService) GetRooms() []*backend.Room {
	rs.RLock()
	defer rs.RUnlock()
	rooms := make([]*backend.Room, 0)
	for _, v := range rs.rooms {
		rooms = append(rooms, v)
	}
	return rooms
}

func (rs *RoomService) JoinRoom(user *backend.UserSession, roomId uuid.UUID) (room *backend.Room, err error) {

	room = rs.GetRoomById(roomId)
	if room == nil {
		return nil, backend.ErrRoomNotFound
	}

	userId := *user.ID

	current_room := rs.GetUserRoom(userId)

	rs.Lock()
	if current_room != nil && current_room.ID != roomId {
		log.Println("Leaving current room...")
		current_room.RemoveUser(userId)
		rs.userRoomMap[userId] = nil
		log.Println("Current room left!")
	}

	if room.GetUser(userId) == nil {
		room.AddUser(user)
		rs.userRoomMap[userId] = room
	}

	rs.Unlock()
	return room, nil
}

func (rs *RoomService) LeaveRoom(userId uuid.UUID, roomId uuid.UUID) error {
	room := rs.GetRoomById(roomId)
	if room == nil {
		return backend.ErrRoomNotFound
	}

	rs.Lock()
	user := room.GetUser(userId)
	if user != nil {
		room.RemoveUser(*user.ID)
	}
	rs.userRoomMap[userId] = nil
	rs.Unlock()
	return nil
}

func (rs *RoomService) GetUserRoom(userId uuid.UUID) *backend.Room {
	rs.RLock()
	defer rs.RUnlock()
	return rs.userRoomMap[userId]
}
