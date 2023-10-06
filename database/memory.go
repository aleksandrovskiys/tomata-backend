package database

import (
	"errors"
	"tomata-backend/authentication"
	"tomata-backend/interfaces"
)

type InMemoryDatabase struct {
	users     []interfaces.User
	pomodoros []interfaces.Pomodoro
}

func (db *InMemoryDatabase) GetUser(email string) (interfaces.User, error) {
	for _, user := range db.users {
		if user.Email == email {
			return user, nil
		}
	}

	return interfaces.User{}, errors.New("User not found")

}

func (db *InMemoryDatabase) GetUserById(id int) (interfaces.User, error) {
	for _, user := range db.users {
		if user.Id == id {
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

func (db *InMemoryDatabase) AddPomodoro(pomodoro interfaces.AddPomodoroRequestSchema, user interfaces.User) (interfaces.Pomodoro, error) {
	newPomodoro := interfaces.Pomodoro{
		Id:        len(db.pomodoros) + 1,
		Task:      pomodoro.Task,
		StartTime: pomodoro.StartTime,
		Duration:  pomodoro.Duration,
		User:      user,
	}
	db.pomodoros = append(db.pomodoros, newPomodoro)
	return newPomodoro, nil
}

func (db *InMemoryDatabase) GetPomodoros(user interfaces.User) []interfaces.Pomodoro {
	var pomodoros []interfaces.Pomodoro

	for _, pomodoro := range db.pomodoros {
		if pomodoro.User.Id == user.Id {
			pomodoros = append(pomodoros, pomodoro)
		}
	}

	return pomodoros
}

func (db *InMemoryDatabase) GetPomodoroById(id int) (interfaces.Pomodoro, error) {
	for _, pomodoro := range db.pomodoros {
		if pomodoro.Id == id {
			if pomodoro.User.Id == id {
				return pomodoro, nil
			} else {
				return interfaces.Pomodoro{}, errors.New("Pomodoro not found")
			}
		}
	}

	return interfaces.Pomodoro{}, errors.New("Pomodoro not found")
}

func (db *InMemoryDatabase) DeletePomodoro(pomodoro interfaces.Pomodoro) error {
	for i, p := range db.pomodoros {
		if p.Id == pomodoro.Id && p.User.Id == pomodoro.User.Id {
			db.pomodoros = append(db.pomodoros[:i], db.pomodoros[i+1:]...)
			return nil
		}
	}

	return errors.New("Pomodoro not found")
}

func (db *InMemoryDatabase) GetUserTasks(user interfaces.User) []string {
	var tasks []string

	for _, pomodoro := range db.pomodoros {
		if pomodoro.User.Id == user.Id {
			tasks = append(tasks, pomodoro.Task)
		}
	}

	return tasks
}
