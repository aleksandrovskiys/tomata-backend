package main

import (
	"os"
	"tomata-backend/middlewares"
	"tomata-backend/routers"

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

	router.POST("/login", routers.Login)
	router.POST("/register", routers.Register)

	userInfo := router.Group("/users")
	userInfo.Use(middlewares.AuthRequired())

	userInfo.GET("/me", routers.UserInfo)

	router.Run(getHostname())
}
