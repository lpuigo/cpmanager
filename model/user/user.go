package user

type User struct {
	FullName string
	Login    string
	Password string
}

func New() *User {
	return &User{
		FullName: "",
		Login:    "",
		Password: "",
	}
}
