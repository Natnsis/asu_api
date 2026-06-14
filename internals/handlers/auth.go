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

type LoginRespose struct {
	Message string `json:"message"`
	Token   string `json:"token"`
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
