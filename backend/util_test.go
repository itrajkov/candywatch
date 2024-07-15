package backend

import (
	"context"
	"testing"

	"github.com/google/uuid"
)

func TestGetUserSession(t *testing.T) {
	session_id, err := uuid.NewRandom()
	if err != nil {
		t.Errorf("Failed generating a UUID.")
	}

	expected := NewUserSession(session_id)
	myCtx := context.WithValue(context.Background(), UserSessionKey, expected)

	session := GetUserSession(myCtx)
	if session == nil {
		t.Errorf("got: nil expected: %+v", expected)
	}

	got := session.ID.String()
	if got != expected.ID.String() {
		t.Errorf("got: %s expected: %+v", got, expected)
	}
}
