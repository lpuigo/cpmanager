package route

import (
	"github.com/lpuig/cpmanager/model/manager"
	"net/http"
)

// handleLogin handles the login form submission
func HandleLogin(m manager.Manager, w http.ResponseWriter, r *http.Request) {
	// Parse the form
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		m.Log.ErrorContextWithTime(r.Context(), "Error parsing form: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Get the login and password from the form
	login := r.Form.Get("login")
	password := r.Form.Get("password")

	// Validate the credentials
	usr, valid := m.Users.ValidateCredentials(login, password)
	if !valid {
		// Invalid credentials, redirect back to login page
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		m.Log.InfoContextWithTime(r.Context(), "invalid credentials", "login", login)
		return
	}

	// Create a new session
	session, err := m.Sessions.CreateSession(usr)
	if err != nil {
		http.Error(w, "Error creating session", http.StatusInternalServerError)
		m.Log.ErrorContextWithTime(r.Context(), "failed to create session", "login", login, "errmsg", err.Error())
		return
	}

	// Set the session cookie
	m.Sessions.SetSessionCookie(w, r, session)

	// Redirect to the main page
	http.Redirect(w, r, "/", http.StatusSeeOther)
	m.Log.InfoContextWithTime(r.Context(), "user logged in", "login", login)
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

	m.Log.InfoContextWithTime(r.Context(), "user logged out")

	// Redirect to the login page
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
