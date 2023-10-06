package middlewares

import (
	"tomata-backend/database"

	"github.com/gin-gonic/gin"
)

func Database() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set("db", database.GetDB())
	}
}
