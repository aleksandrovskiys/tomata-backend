package database

import (
	"errors"
	"tomata-backend/authentication"
	"tomata-backend/interfaces"
)

type InMemoryDatabase struct {
	users []interfaces.User
}

func (db *InMemoryDatabase) GetUser(email string) (interfaces.User, error) {
	for _, user := range db.users {
		if user.Email == email {
			return user, nil
		}
	}

	return interfaces.User{}, errors.New("User not found")

}

func (db *InMemoryDatabase) GetUsers() ([]interfaces.User, error) {
	return db.users, nil
}

func (db *InMemoryDatabase) AddUser(email string, password string) (interfaces.User, error) {
	user, err := db.GetUser(email)

	if err == nil {
		return user, errors.New("User already exists")
	}

	user = interfaces.User{
		Id:       len(db.users) + 1,
		Email:    email,
		Password: authentication.GeneratePasswordHash(password),
	}

	db.users = append(db.users, user)

	return user, nil

}

func (db *InMemoryDatabase) DeleteUser(user interfaces.User) error {
	for i, u := range db.users {
		if u.Id == user.Id {
			db.users = append(db.users[:i], db.users[i+1:]...)
			return nil
		}
	}

	return errors.New("User not found")
}
