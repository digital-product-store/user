package storage

type User struct {
	Uuid     string
	Email    string
	Username string
	Fullname string
	Password string
	Roles    []string
}

func (u User) ValidatePassword(password string) bool {
	return u.Password == password
}
