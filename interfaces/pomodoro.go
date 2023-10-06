package interfaces

import "time"

type AddPomodoroRequestSchema struct {
	Task      string    `json:"task"`
	StartTime time.Time `json:"startTime"`
	Duration  int       `json:"duration"`
}
type Pomodoro struct {
	Id        int       `json:"id"`
	Task      string    `json:"task"`
	StartTime time.Time `json:"startTime"`
	Duration  int       `json:"duration"`
	User      User      `json:"-"`
}
