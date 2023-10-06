package routers

import (
	"net/http"
	"tomata-backend/interfaces"

	"github.com/gin-gonic/gin"
)

func UserInfo(ctx *gin.Context) {
	user := ctx.MustGet("user").(interfaces.User)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"user":    user,
	})
}
