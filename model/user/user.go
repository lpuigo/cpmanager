package user

type User struct {
	Name     string
	Login    string
	Password string
}

func New() *User {
	return &User{
		Name:     "",
		Login:    "",
		Password: "",
	}
}
