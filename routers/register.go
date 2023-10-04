package routers

import (
	"fmt"
	"net/http"
	"os"
	"tomata-backend/database"
	"tomata-backend/interfaces"

	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {
	loginData := interfaces.LoginDataSchema{}
	err := ctx.BindJSON(&loginData)

	db := database.GetDB()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
		})

		return
	}

	user, err := db.AddUser(loginData.Email, loginData.Password)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "This email is already registered",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Successfully registered",
		"user":    user,
	})
}
