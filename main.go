package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello, world!")
	router := gin.Default()

	router.GET("/login", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "login",
		})
	})

	router.POST("/registter", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "registration",
		})
	})

	router.Run()
}
