package authentication

import (
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"os"
	"time"

	"tomata-backend/interfaces"

	"github.com/golang-jwt/jwt"
)

const expirationTime = 24 * time.Hour

func GeneratePasswordHash(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))
	passwordHash := base64.URLEncoding.EncodeToString(hash.Sum(nil))
	return passwordHash
}

func ValidatePassword(password string, passwordHash string) bool {
	return GeneratePasswordHash(password) == passwordHash
}

func IssueToken(user interfaces.User) (string, error) {
	jwtKey := os.Getenv("JWT_KEY")

	nowTime := time.Now()

	expirationTime := nowTime.Add(expirationTime)

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Subject:   user.Email,
		Id:        fmt.Sprint(user.Id),
		Issuer:    "tomata-backend",
		Audience:  "tomata-frontend",
		ExpiresAt: expirationTime.Unix(),
		IssuedAt:  nowTime.Unix(),
		NotBefore: nowTime.Unix(),
	})
	s, err := t.SignedString([]byte(jwtKey))

	return s, err
}

func ValidateToken(tokenString string) (jwt.StandardClaims, error) {
	jwtKey := os.Getenv("JWT_KEY")
	return extractClaims(tokenString, jwtKey)
}

func extractClaims(tokenString string, jwtKey string) (jwt.StandardClaims, error) {
	claims := jwt.StandardClaims{}

	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	})

	if err != nil {
		return jwt.StandardClaims{}, err
	}

	if !token.Valid {
		return jwt.StandardClaims{}, errors.New("Invalid token")
	}

	return claims, nil
}

func GenerateState() string {
	hash := sha256.New()
	_, err := hash.Write([]byte(time.Now().String()))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return ""
	}
	return base64.URLEncoding.EncodeToString(hash.Sum(nil))
}

func GetUserInfoFromToken(token string) (interfaces.GoogleLoginDataSchema, error) {
	claims := jwt.MapClaims{}
	_, _, err := new(jwt.Parser).ParseUnverified(token, &claims)

	if err != nil {
		return interfaces.GoogleLoginDataSchema{}, err
	}
	return interfaces.GoogleLoginDataSchema{
		Email:    claims["email"].(string),
		Name:     claims["name"].(string),
		GoogleID: claims["sub"].(string),
	}, nil
}
