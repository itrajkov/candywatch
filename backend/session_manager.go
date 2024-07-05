package backend

import (
	"fmt"

	"github.com/google/uuid"
)


// TODO: This should be a database repository instead with small in-memory cache
// postponed for later in development
type SessionManager struct {
	userSessions []UserSession
}

func NewSessionManager() *SessionManager {
	return &SessionManager {
		userSessions: make([]UserSession, 2),
	}
}

func (rm *SessionManager) GetUserSession(sessionID uuid.UUID) *UserSession {
	// TODO: Implement fetching the UserSession
	fmt.Printf("GetUserSession: %s\n", sessionID)
	return nil
}

func (sm *SessionManager)AddSession(user UserSession) {
	sm.userSessions = append(sm.userSessions, user)
}
