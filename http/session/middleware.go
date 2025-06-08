package session

import (
	"context"
	"net/http"
)

// SessionKey is the key used to store the session in the request context
type SessionKey string

const (
	// ContextKeySession is the key used to store the session in the request context
	ContextKeySession SessionKey = "session"
)

// WithSession is a middleware that checks if a user is authenticated and injects the session into the request context
func (s *Sessions) WithSession(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the session from the request
		session, ok := s.GetSessionFromRequest(r)
		if !ok {
			// No session found, continue without a session
			next.ServeHTTP(w, r)
			return
		}

		// Add the session to the request context
		ctx := context.WithValue(r.Context(), ContextKeySession, session)
		r = r.WithContext(ctx)

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}

// RequireAuth is a middleware that requires authentication
func (s *Sessions) RequireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the session from the request
		session, ok := s.GetSessionFromRequest(r)
		if !ok {
			// No session found, redirect to login page
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		// Add the session to the request context
		ctx := context.WithValue(r.Context(), ContextKeySession, session)
		r = r.WithContext(ctx)

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}

// GetSessionFromContext gets the session from the request context
func GetSessionFromContext(ctx context.Context) (*Session, bool) {
	session, ok := ctx.Value(ContextKeySession).(*Session)
	return session, ok
}
