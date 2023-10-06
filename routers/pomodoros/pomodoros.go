package pomodoros

import (
	"net/http"
	"tomata-backend/interfaces"

	"github.com/gin-gonic/gin"
)

func AddPomodoro(ctx *gin.Context) {

	user := ctx.MustGet("user").(interfaces.User)
	pomodoro := interfaces.AddPomodoroRequestSchema{}

	ctx.BindJSON(&pomodoro)

	db := ctx.MustGet("db").(interfaces.Database)

	newPomodoro, err := db.AddPomodoro(pomodoro, user)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	ctx.JSON(http.StatusCreated, newPomodoro)

}

func GetUserPomodoros(ctx *gin.Context) {
	user := ctx.MustGet("user").(interfaces.User)
	db := ctx.MustGet("db").(interfaces.Database)

	pomodoros := db.GetPomodoros(user)

	ctx.JSON(http.StatusOK, pomodoros)
}

func GetUserTasks(ctx *gin.Context) {
	user := ctx.MustGet("user").(interfaces.User)
	db := ctx.MustGet("db").(interfaces.Database)

	tasks := db.GetUserTasks(user)

	ctx.JSON(http.StatusOK, tasks)
}
