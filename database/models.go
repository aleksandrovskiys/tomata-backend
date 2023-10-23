package database

import (
	"time"

	"gorm.io/gorm"
)

type Pomodoro struct {
	gorm.Model

	Task     *string
	Finished time.Time
	Duration int
	UserID   uint
	User     User
}

type User struct {
	gorm.Model

	Email     string
	Password  string
	GoogleID  string
	Pomodoros []Pomodoro
}
