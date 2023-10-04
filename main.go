package main

import (
	"tomata-backend/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/login", routers.Login)
	router.POST("/register", routers.Register)

	router.Run()
}
