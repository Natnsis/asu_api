package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"UniCore/internals/db"
	"UniCore/internals/models"

	"github.com/gorilla/mux"
)

func CreateRecent(w http.ResponseWriter, r *http.Request) {
	var recent models.RecentActivity
	err := json.NewDecoder(r.Body).Decode(&recent)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	result := db.DbConnection.Create(&recent)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(recent)
}

func GetRecentActivities(w http.ResponseWriter, r *http.Request) {
	var recent models.RecentActivity
	result := db.DbConnection.Find(&recent)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(recent)
}

func GetSingleRecentActivity(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var recent models.RecentActivity
	result := db.DbConnection.First(&recent, id)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(recent)
}
