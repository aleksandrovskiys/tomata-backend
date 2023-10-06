package middlewares

import (
	"strconv"
	"strings"
	"tomata-backend/authentication"
	"tomata-backend/interfaces"

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
		db := ctx.MustGet("db").(interfaces.Database)

		user, err := db.GetUserById(id)

		if err != nil {
			ctx.AbortWithStatusJSON(401, gin.H{
				"message": "Unauthorized",
			})
		}

		ctx.Set("user", user)
	}
}
