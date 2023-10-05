package routers

import (
	"net/http"
	"tomata-backend/database"

	"github.com/gin-gonic/gin"
)

func UserInfo(ctx *gin.Context) {
	userId, exists := ctx.Get("userId")

	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	user, err := database.GetDB().GetUserById(userId.(int))

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Hello world!",
		"user":    user,
	})
}
