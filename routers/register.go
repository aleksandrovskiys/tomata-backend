package routers

import (
	"fmt"
	"net/http"
	"os"
	"tomata-backend/authentication"
	googleoauth "tomata-backend/authentication/googleOAuth"
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

	user, err := db.AddUser(loginData.Email, authentication.GeneratePasswordHash(loginData.Password), "")

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "This email is already taken",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Successfully registered",
		"user":    user,
	})
}

func RegisterUsingGoogleOpenID(ctx *gin.Context) {
	var tokenData interfaces.GoogleOpenIDParametersSchema
	err := ctx.BindJSON(&tokenData)
	db := database.GetDB()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
		})
		return
	}

	redirectURI := os.Getenv("GOOGLE_SIGNUP_REDIRECT_URI")
	userInfo, err := googleoauth.GetUserInfo(tokenData, redirectURI)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	user, err := db.AddUser(userInfo.Email, "", userInfo.GoogleID)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "This email is already taken",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Successfully registered",
		"user":    user,
	})

}
