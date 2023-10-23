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

func Login(ctx *gin.Context) {
	var loginData interfaces.LoginDataSchema
	err := ctx.BindJSON(&loginData)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
		})
		return
	}

	db := database.GetDB()
	user, err := db.GetUser(loginData.Email)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid email or password",
		})
		return
	}

	if !authentication.ValidatePassword(loginData.Password, user.Password) {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid email or password",
		})
		return
	}

	s, err := authentication.IssueToken(user)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Login successfuly",
		"token":   s,
	})
}

func LoginWithGoogle(ctx *gin.Context) {
	var tokenData interfaces.GoogleOpenIDParametersSchema
	err := ctx.BindJSON(&tokenData)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
		})
		return
	}
	redirectURI := os.Getenv("GOOGLE_LOGIN_REDIRECT_URI")
	userInfo, err := googleoauth.GetUserInfo(tokenData, redirectURI)

	db := database.GetDB()
	user, err := db.GetUser(userInfo.Email)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid email or password",
		})
		return
	}

	if userInfo.GoogleID != user.GoogleID {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid email or password",
		})
		return
	}

	s, err := authentication.IssueToken(user)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Login successfuly",
		"token":   s,
	})
}

func GetAntiForgeryToken(ctx *gin.Context) {
	state := authentication.GenerateState()
	authentication.StoreAntiForgeryToken(state)
	ctx.JSON(http.StatusOK, gin.H{
		"state": state,
	})
}
