package database

import "tomata-backend/interfaces"

type Database interface {
	GetUser(email string) (interfaces.User, error)
	GetUsers() ([]interfaces.User, error)
	AddUser(email string, password string) (interfaces.User, error)
	DeleteUser(user interfaces.User) error
}

var database = &InMemoryDatabase{}

func GetDB() Database {
	return database
}
