package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"UniCore/internals/db"
	"UniCore/internals/models"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Email    string
	Password string
}

type RefreshToken struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint      `gorm:"not null"`
	Token     string    `gorm:"uniqueIndex;not null"`
	ExpiresAt time.Time `gorm:"not null"`
}

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type LoginRespose struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

func GenerateAccessToken(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(time.Minute * 15).Unix(), // 15 Minute short lifespan
		"iat": time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	// get request data
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var req models.User
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Unable to fetch user data", http.StatusBadRequest)
		return
	}

	if len(req.Password) < 8 {
		http.Error(w, "Password must be atleast 8 characters", http.StatusBadRequest)
		return
	}

	// hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Unable to hash passwords", http.StatusBadRequest)
		return
	}

	// contain the data together
	newUser := models.User{
		Email:    req.Email,
		Password: string(hashedPassword),
		FullName: req.FullName,
	}

	// save user
	if result := db.Db.Create(&newUser); result.Error != nil {
		http.Error(w, "user saved successfully", http.StatusCreated)
	}

	// w data
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	// method check
	if r.Method != http.MethodPost {
		http.Error(w, "only post method works here", http.StatusBadGateway)
		return
	}

	// get request data
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "unable to fetch user request", http.StatusBadGateway)
		return
	}

	// check email
	var user models.User
	if err := db.Db.Where("email=?", req.Email).First(&user).Error; err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid email"})
		return
	}

	// verify passwords
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "Incorrect password"})
		return
	}

	// generate the jwt claims payload
	claims := jwt.MapClaims{
		"sub": user.ID,                               //subject (userID)
		"exp": time.Now().Add(time.Hour * 24).Unix(), // expires at
		"iat": time.Now().Unix(),                     // issued at
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "jwt secret key is missing"})
		return
	}

	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "failed to generate token"})
		return
	}

	response := LoginRespose{
		Message: "Login successful",
		Token:   tokenString,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// LogoutHandler - POST /logout
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	// Parse the body to find out which token to destroy
	var payload RefreshRequest // Reusing the struct from the refresh step
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid payload"})
		return
	}

	// Delete the token tracking record using GORM
	// Unscoped() completely purges the record instead of soft-deleting it
	result := db.Db.Unscoped().Where("token = ?", payload.RefreshToken).Delete(RefreshToken{})

	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Could not complete logout action"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Logged out successfully. Token revoked."})
}

// RefreshHandler - POST /refresh
func RefreshHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var payload RefreshRequest
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid payload"})
		return
	}

	// 1. Look up the refresh token in the DB
	var storedToken RefreshToken
	if err := db.Db.Where("token = ?", payload.RefreshToken).First(&storedToken).Error; err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid refresh token"})
		return
	}

	// 2. Check if the refresh token has expired
	if time.Now().After(storedToken.ExpiresAt) {
		db.Db.Delete(&storedToken) // Clean up expired token from DB
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "Refresh token expired. Please login again."})
		return
	}

	// 3. Token is valid! Issue a fresh brand new 15-minute access token
	newAccessToken, err := GenerateAccessToken(storedToken.UserID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Could not generate access token"})
		return
	}

	// 4. Send the new access token back to the frontend
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"access_token": newAccessToken,
	})
}
