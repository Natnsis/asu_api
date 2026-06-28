package handlers

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

var secret = []byte(os.Getenv("JWT_SECRET"))

func GeneratedAccessToken(userID uint, email string) (string, error) {
}
func GeneratedRefreshToken(userID uint, email string) (string, error) {}
