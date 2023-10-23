package interfaces

type Database interface {
	GetUser(email string) (User, error)
	GetUserById(id int) (User, error)
	GetUsers() ([]User, error)
	AddUser(email string, password string) (User, error)
	DeleteUser(user User) error

	AddPomodoro(pomodoro AddPomodoroRequestSchema, user User) (Pomodoro, error)
	GetPomodoros(user User) []Pomodoro
	GetPomodoroById(id int) (Pomodoro, error)
	DeletePomodoro(pomodoro Pomodoro) error

	GetUserTasks(user User) []string

	Init()
	Migrate() error
	Initialized() bool
}
