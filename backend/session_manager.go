package backend

import (
	"github.com/google/uuid"
)

// TODO: This should be a database repository instead with small in-memory cache
// postponed for later in development
type SessionManager struct {
	userSessions []UserSession
}

func NewSessionManager() *SessionManager {
	return &SessionManager {
		userSessions: make([]UserSession, 0),
	}
}

func (sm *SessionManager) GetUserSession(sessionID uuid.UUID) *UserSession {
	for _, session := range sm.userSessions {
		if sessionID == *session.ID {
			return &session
		}
	}
	return nil
}

func (sm *SessionManager)AddSession(user UserSession) {
	// TODO: Check if UserSession with that sessionID already exists
	sm.userSessions = append(sm.userSessions, user)
}
