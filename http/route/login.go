package route

import (
	"github.com/lpuig/cpmanager/model/manager"
	"github.com/lpuig/cpmanager/model/user"
	"net/http"
)

// handleLogin handles the login form submission
func HandleLogin(m manager.Manager, w http.ResponseWriter, r *http.Request) {
	// Parse the form
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	// Get the username and password from the form
	username := r.Form.Get("username")
	password := r.Form.Get("password")

	// Validate the credentials
	user, valid := user.ValidateCredentials(username, password)
	if !valid {
		// Invalid credentials, redirect back to login page
		m.Log.InfoContextWithTime(r.Context(), "invalid credentials", "username", username)
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// Create a new session
	session, err := m.Sessions.CreateSession(user)
	if err != nil {
		m.Log.ErrorContextWithTime(r.Context(), "failed to create session", "username", username, "errmsg", err.Error())
		http.Error(w, "Error creating session", http.StatusInternalServerError)
		return
	}

	// Set the session cookie
	m.Sessions.SetSessionCookie(w, r, session)

	m.Log.InfoContextWithTime(r.Context(), "user logged in", "username", username)

	// Redirect to the main page
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// handleLogout handles the logout request
func HandleLogout(m manager.Manager, w http.ResponseWriter, r *http.Request) {
	// Get the session from the request
	var userId string
	session, ok := m.Sessions.GetSessionFromRequest(r)
	if ok {
		userId = session.UserID
		// Delete the session (this is now handled by ClearSessionCookie)
		m.Sessions.DeleteSession(userId)
	}

	// Clear the session cookie
	m.Sessions.ClearSessionCookie(w, r)

	m.Log.InfoContextWithTime(r.Context(), "user logged out", "username", userId)

	// Redirect to the login page
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
