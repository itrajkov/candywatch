package backend

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func UserSessionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		log.Println("Getting session..")
		cookie, err := r.Cookie("session_id")
		var session *UserSession

		if err == nil {
			sessionID, err := uuid.Parse(cookie.Value)
			if err == nil {
				log.Printf("Resuming session %s\n", sessionID)
				session = &UserSession{ID: &sessionID}
			}
		}

		if session == nil {
			log.Println("No session found, creating new session..")
			session = NewUser()
			session_id := session.ID.String()
			http.SetCookie(w, &http.Cookie{
				Name:    "session_id",
				Value:   session_id,
				Expires: time.Now().Add(24 * time.Hour),
			})
			log.Printf("New session created %s!\n", session_id)
		}

		ctx := context.WithValue(r.Context(), userSessionKey, session)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
