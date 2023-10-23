package main

import (
	"os"
	"tomata-backend/database"
	"tomata-backend/middlewares"
	"tomata-backend/routers"
	"tomata-backend/routers/pomodoros"

	"github.com/gin-gonic/gin"
)

func getHostname() string {

	var hostname string
	if len(os.Args) > 1 {
		hostname = os.Args[1]
	} else {
		hostname = "localhost:8080"
	}

	return hostname
}

func main() {
	router := gin.Default()
	router.Use(middlewares.AllowCORS())
	router.Use(middlewares.Database())

	router.POST("/login", routers.Login)
	router.POST("/register", routers.Register)
	router.POST("/google-auth/register", routers.RegisterUsingGoogleOpenID)
	router.POST("/google-auth/login", routers.LoginWithGoogle)
	router.GET("/anti-forgery-token", routers.GetAntiForgeryToken)

	users := router.Group("/users")
	users.Use(middlewares.AuthRequired())

	users.GET("/me", routers.UserInfo)
	users.POST("/pomodoros", pomodoros.AddPomodoro)
	users.GET("/pomodoros", pomodoros.GetUserPomodoros)
	users.GET("/tasks", pomodoros.GetUserTasks)
	db := database.GetDB()
	db.Init()

	router.Run(getHostname())
}
