package middleware

import (
	"net/http"
	"strings"
	"time"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("your-neon-app-secret")

func GenerateToken ()
