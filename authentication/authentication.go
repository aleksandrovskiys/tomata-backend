package authentication

import (
	"crypto/sha256"
	"encoding/base64"
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
