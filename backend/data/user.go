package data

type User struct {
	Id           string
	Username     string
	PasswordHash string
	CreatedAt    int64
}

type UsersDataService interface {
	Add(username, passwordHash string) error
	FindByUsername(string) (*User, error)
	FindById(string) (*User, error)
}
