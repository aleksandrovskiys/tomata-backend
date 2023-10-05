package middlewares

import (
	"strconv"
	"strings"
	"tomata-backend/authentication"

	"github.com/gin-gonic/gin"
)

func AuthRequired() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		token = strings.Replace(token, "Bearer ", "", 1)

		claims, err := authentication.ValidateToken(token)

		if err != nil {
			ctx.AbortWithStatusJSON(401, gin.H{
				"message": "Unauthorized",
			})
			return
		}
		id, err := strconv.Atoi(claims.Id)
		ctx.Set("userId", id)
	}
}
