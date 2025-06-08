package session

import (
	"github.com/lpuig/cpmanager/model/user"
	"net/http/httptest"
	"testing"
)

func TestSessionManagement(t *testing.T) {
	// Create a new Sessions manager
	sessions := New()

	// Create a test user
	testUser := &user.User{
		Name:     "Test User",
		Login:    "testuser",
		Password: "password",
	}

	// Create a new session
	session, err := sessions.CreateSession(testUser)
	if err != nil {
		t.Fatalf("Error creating session: %v", err)
	}

	// Verify the session was created
	if session.UserID != testUser.Login {
		t.Errorf("Expected session UserID to be %s, got %s", testUser.Login, session.UserID)
	}

	// Test session cookie
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	sessions.SetSessionCookie(w, r, session)
	cookies := w.Result().Cookies()
	if len(cookies) != 1 {
		t.Fatalf("Expected 1 cookie, got %d", len(cookies))
	}
	if cookies[0].Name != "session" {
		t.Errorf("Expected cookie name to be 'session', got '%s'", cookies[0].Name)
	}

	// Test getting session from request
	req := httptest.NewRequest("GET", "/", nil)
	req.AddCookie(cookies[0])
	sessionFromReq, ok := sessions.GetSessionFromRequest(req)
	if !ok {
		t.Fatalf("Session not found in request")
	}
	if sessionFromReq.UserID != session.UserID {
		t.Errorf("Expected session UserID to be '%s', got '%s'", session.UserID, sessionFromReq.UserID)
	}

	// Test clearing session cookie
	w = httptest.NewRecorder()
	r = httptest.NewRequest("GET", "/", nil)
	r.AddCookie(cookies[0])
	sessions.ClearSessionCookie(w, r)
	cookies = w.Result().Cookies()
	if len(cookies) != 1 {
		t.Fatalf("Expected 1 cookie, got %d", len(cookies))
	}
	if cookies[0].Name != "session" {
		t.Errorf("Expected cookie name to be 'session', got '%s'", cookies[0].Name)
	}
	if cookies[0].MaxAge != -1 {
		t.Errorf("Expected cookie MaxAge to be -1, got %d", cookies[0].MaxAge)
	}
}
