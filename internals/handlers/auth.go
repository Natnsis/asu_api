package handlers

import (
	"encoding/json"
	"net/http"

	"UniCore/internals/db"
	"UniCore/internals/models"

	"golang.org/x/crypto/bcrypt"
)

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
}
