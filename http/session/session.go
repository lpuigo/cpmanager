package session

import (
	"github.com/gorilla/sessions"
	"github.com/lpuig/cpmanager/config"
	"github.com/lpuig/cpmanager/model/user"
	"net/http"
	"time"
)

// Session represents a user session
type Session struct {
	UserID    string
	User      *user.User
	CreatedAt time.Time
	ExpiresAt time.Time
}

// Sessions manages all active sessions using gorilla/sessions
type Sessions struct {
	store  sessions.Store
	name   string
	maxAge int
}

// New creates a new Sessions manager with a cookie store using default settings
func New() *Sessions {
	// Create a key for cookie store
	key := []byte("cpmanager-session-key-replace-in-production")

	// Create a cookie store
	store := sessions.NewCookieStore(key)

	// Configure the store
	store.Options = &sessions.Options{
		Path:     "/",
		HttpOnly: true,
		Secure:   false, // Set to true in production with HTTPS
		SameSite: http.SameSiteLaxMode,
		MaxAge:   86400, // 24 hours in seconds
	}

	return &Sessions{
		store:  store,
		name:   "session",
		maxAge: 86400, // 24 hours in seconds
	}
}

// NewWithConfig creates a new Sessions manager with a cookie store using the provided config
func NewWithConfig(cfg config.Config) *Sessions {
	// Use the session key from config
	key := []byte(cfg.SessionKey)

	// Create a cookie store
	store := sessions.NewCookieStore(key)

	// Configure the store
	store.Options = &sessions.Options{
		Path:     "/",
		HttpOnly: true,
		Secure:   false, // Set to true in production with HTTPS
		SameSite: http.SameSiteLaxMode,
		MaxAge:   86400, // 24 hours in seconds
	}

	return &Sessions{
		store:  store,
		name:   "session",
		maxAge: 86400, // 24 hours in seconds
	}
}

// CreateSession creates a new session for the given user
func (s *Sessions) CreateSession(u *user.User) (*Session, error) {
	// Create a new session object
	session := &Session{
		UserID:    u.Login,
		User:      u,
		CreatedAt: time.Now(),
		ExpiresAt: time.Now().Add(time.Duration(s.maxAge) * time.Second),
	}

	return session, nil
}

// GetSession retrieves a session by ID (not used with gorilla/sessions)
func (s *Sessions) GetSession(sessionID string) (*Session, bool) {
	// This method is kept for compatibility but not used with gorilla/sessions
	return nil, false
}

// DeleteSession removes a session
func (s *Sessions) DeleteSession(sessionID string) {
	// This method is kept for compatibility but not used directly with gorilla/sessions
	// The actual session deletion happens in ClearSessionCookie
}

// SetSessionCookie sets a session cookie in the HTTP response
func (s *Sessions) SetSessionCookie(w http.ResponseWriter, r *http.Request, session *Session) {
	// Get a new session from the store
	gSession, err := s.store.New(r, s.name)
	if err != nil {
		// Handle error
		return
	}

	// Store user data in the session
	gSession.Values["userID"] = session.UserID
	gSession.Values["userName"] = session.User.Name
	gSession.Values["createdAt"] = session.CreatedAt.Unix()

	// Save the session
	err = gSession.Save(r, w)
	if err != nil {
		// Handle error
		return
	}
}

// ClearSessionCookie clears the session cookie
func (s *Sessions) ClearSessionCookie(w http.ResponseWriter, r *http.Request) {
	// Get the existing session
	gSession, err := s.store.Get(r, s.name)
	if err != nil {
		// If there's an error, just set a new cookie with negative max age
		gSession, _ = s.store.New(r, s.name)
	}

	// Set the max age to -1 to expire the cookie
	gSession.Options.MaxAge = -1

	// Save the session (which will delete it due to negative MaxAge)
	gSession.Save(r, w)
}

// GetSessionFromRequest gets the session from the request cookies
func (s *Sessions) GetSessionFromRequest(r *http.Request) (*Session, bool) {
	// Get the session from the store
	gSession, err := s.store.Get(r, s.name)
	if err != nil {
		return nil, false
	}

	// Check if the session is new (not previously stored)
	if gSession.IsNew {
		return nil, false
	}

	// Get user data from the session
	userID, ok := gSession.Values["userID"].(string)
	if !ok {
		return nil, false
	}

	userName, ok := gSession.Values["userName"].(string)
	if !ok {
		return nil, false
	}

	createdAtUnix, ok := gSession.Values["createdAt"].(int64)
	if !ok {
		return nil, false
	}

	createdAt := time.Unix(createdAtUnix, 0)
	expiresAt := createdAt.Add(time.Duration(s.maxAge) * time.Second)

	// Create a user object
	u := &user.User{
		Login: userID,
		Name:  userName,
	}

	// Create a session object
	session := &Session{
		UserID:    userID,
		User:      u,
		CreatedAt: createdAt,
		ExpiresAt: expiresAt,
	}

	return session, true
}
