package interfaces

import (
	"github.com/google/uuid"
	"github.com/itrajkov/candywatch/backend"
)

type RoomsManager interface {
	GetRooms() []*backend.Room
	GetRoomById(uuid.UUID) *backend.Room
	JoinRoom(*backend.UserSession, uuid.UUID) (*backend.Room, error)
	LeaveRoom(uuid.UUID, uuid.UUID) (error)
	GetUserRoom(uuid.UUID) *backend.Room
	NewRoom() *backend.Room
}
