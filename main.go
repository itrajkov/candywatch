package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"github.com/go-chi/chi/v5"
	// "github.com/gorilla/websocket"
)

func main() {
	r := chi.NewRouter()
	s := &http.Server{
		Addr:          ":8080",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
		Handler: r,
	}

	r.HandleFunc("/ping", pong)

	log.Fatal(s.ListenAndServe())

}

func pong(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong\n")
}
