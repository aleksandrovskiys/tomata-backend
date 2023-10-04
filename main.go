package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/login", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "login",
		})
	})

	router.POST("/register", func(ctx *gin.Context) {
		ctx.JSON(http.StatusCreated, gin.H{
			"message": "registration",
		})
	})

	router.Run()
}
