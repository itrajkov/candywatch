package backend

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func UserSessionMiddleware(sessionManager *SessionManager) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var session *UserSession

			log.Println("Getting session cookie..")
			cookie, err := r.Cookie("session_id")

			if err != nil {
				log.Println("No session cookie found, creating new session..")
				sessionID, err := uuid.NewRandom()
				if err != nil {
					log.Fatal("failed to generate user sessionID:", err)
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				http.SetCookie(w, &http.Cookie{
					Name:    "session_id",
					Path: "/",
					Value:   sessionID.String(),
					Expires: time.Now().Add(24 * time.Hour),
				})

				session = NewUserSession(sessionID)
				sessionManager.AddSession(session)
				log.Printf("New session created %s!\n", sessionID)
				log.Printf("SessionManager state: %+v", sessionManager)
			} else {
				log.Println("Parsing cookie as UUID")
				sessionID, err := uuid.Parse(cookie.Value)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				session = sessionManager.GetUserSession(sessionID)
				if session == nil {
					log.Println("Creating new user session..")
					session = NewUserSession(sessionID)
					sessionManager.AddSession(session)
				}
				log.Printf("SessionManager state: %+v", sessionManager)
			}

			ctx := context.WithValue(r.Context(), UserSessionKey, session)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}

}
