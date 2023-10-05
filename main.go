package main

import (
	"os"
	"tomata-backend/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	var hostname string
	if len(os.Args) > 1 {
		hostname = os.Args[1]
	} else {
		hostname = "localhost:8080"
	}

	router.POST("/login", routers.Login)
	router.POST("/register", routers.Register)

	router.Run(hostname)
}
