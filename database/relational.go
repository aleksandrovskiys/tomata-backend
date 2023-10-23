package database

import (
	"fmt"
	"tomata-backend/interfaces"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type RelationalDB struct {
	instance    *gorm.DB
	initialized bool
}

func (db *RelationalDB) Init() {
	if db.Initialized() {
		return
	}
	instance, err := gorm.Open(sqlite.Open("tomata.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to create database connection")
	}

	db.instance = instance
	err = db.Migrate()
	if err != nil {
		fmt.Println(err)
		panic("Failed to migrate database")
	}
	fmt.Println("Database migrated successfully")
	db.initialized = true
}

func (db *RelationalDB) Migrate() error {
	err := db.instance.AutoMigrate(&User{})
	if err != nil {
		return err
	}
	err = db.instance.AutoMigrate(&Pomodoro{})
	if err != nil {
		return err
	}
	return nil
}

func (db *RelationalDB) GetUser(email string) (interfaces.User, error) {
	var user User
	result := db.instance.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return interfaces.User{}, result.Error
	}

	interfaceUser := db.convertUserToInterface(user)
	return interfaceUser, nil
}

func (db *RelationalDB) GetUserById(id int) (interfaces.User, error) {
	var user User
	result := db.instance.Where("id = ?", id).First(&user)
	if result.Error != nil {
		return interfaces.User{}, result.Error
	}

	return db.convertUserToInterface(user), nil
}

func (db *RelationalDB) GetUsers() ([]interfaces.User, error) {
	var users []User
	result := db.instance.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	var interfaceUsers []interfaces.User
	for _, user := range users {
		interfaceUsers = append(interfaceUsers, db.convertUserToInterface(user))
	}

	return interfaceUsers, nil
}

func (db *RelationalDB) AddUser(email string, password string) (interfaces.User, error) {
	user := User{
		Email:    email,
		Password: password,
	}

	result := db.instance.Create(&user)
	if result.Error != nil {
		return interfaces.User{}, result.Error
	}

	return db.convertUserToInterface(user), nil
}

func (db *RelationalDB) DeleteUser(user interfaces.User) error {
	var dbUser User

	result := db.instance.Where("id = ?", user.Id).First(&dbUser)
	if result.Error != nil {
		return result.Error
	}

	result = db.instance.Delete(&dbUser)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (db *RelationalDB) AddPomodoro(pomodoro interfaces.AddPomodoroRequestSchema, user interfaces.User) (interfaces.Pomodoro, error) {
	pom := Pomodoro{
		Task:     &pomodoro.Task,
		Duration: pomodoro.Duration,
		UserID:   uint(user.Id),
		Finished: pomodoro.Finished,
	}

	result := db.instance.Create(&pom)
	if result.Error != nil {
		return interfaces.Pomodoro{}, result.Error
	}

	return db.convertPomodoroToInterface(pom), nil
}

func (db *RelationalDB) GetPomodoros(user interfaces.User) []interfaces.Pomodoro {
	var pomodoros []Pomodoro
	db.instance.Where("user_id = ?", user.Id).Find(&pomodoros)

	interfacePomodoros := []interfaces.Pomodoro{}
	for _, pomodoro := range pomodoros {
		interfacePomodoros = append(interfacePomodoros, db.convertPomodoroToInterface(pomodoro))
	}
	return interfacePomodoros
}

func (db *RelationalDB) GetPomodoroById(id int) (interfaces.Pomodoro, error) {
	var pomodoro Pomodoro
	result := db.instance.Where("id = ?", id).First(&pomodoro)
	if result.Error != nil {
		return interfaces.Pomodoro{}, result.Error
	}

	return db.convertPomodoroToInterface(pomodoro), nil
}

func (db *RelationalDB) DeletePomodoro(pomodoro interfaces.Pomodoro) error {
	var dbPomodoro Pomodoro

	result := db.instance.Where("id = ?", pomodoro.Id).First(&dbPomodoro)
	if result.Error != nil {
		return result.Error
	}

	result = db.instance.Delete(&dbPomodoro)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (db *RelationalDB) GetUserTasks(user interfaces.User) []string {
	var pomodoros []Pomodoro

	db.instance.Where("user_id = ?", user.Id).Find(&pomodoros)

	var tasks []string
	for _, pomodoro := range pomodoros {
		tasks = append(tasks, *pomodoro.Task)
	}

	return tasks
}

func (db *RelationalDB) convertUserToInterface(user User) interfaces.User {
	return interfaces.User{
		Id:       int(user.ID),
		Email:    user.Email,
		Password: user.Password,
	}
}

func (db *RelationalDB) convertPomodoroToInterface(pomodoro Pomodoro) interfaces.Pomodoro {
	return interfaces.Pomodoro{
		Id:       int(pomodoro.ID),
		Task:     *pomodoro.Task,
		Finished: pomodoro.Finished,
		Duration: pomodoro.Duration,
	}
}

func (db *RelationalDB) Initialized() bool {
	return db.initialized
}
