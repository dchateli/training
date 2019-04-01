package db

type User struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type DbContract interface {
	AddUser(u User) (User, error)
	ListUser() ([]User, error)
	DeleteUser(userId string) error
	GetUser(userId string) (User, error)
	UpdateUser(userId string, u User) (User, error)
}
