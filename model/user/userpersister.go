package user

// ValidateCredentials validates the user credentials
func ValidateCredentials(username, password string) (*User, bool) {
	// In a real application, you would check the credentials against a database
	// For this simple implementation, we'll just hardcode a user
	if username == "admin" && password == "password" {
		return &User{
			Name:     "Administrator",
			Login:    "admin",
			Password: "password", // In a real application, you would never store passwords in plain text
		}, true
	}
	return nil, false
}
