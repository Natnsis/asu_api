package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"UniCore/internals/db"
	"UniCore/internals/models"

	"github.com/gorilla/mux"
)

func CreateCafeSchedule(w http.ResponseWriter, r *http.Request) {
	var schedule models.Cafeteria
	err := json.NewDecoder(r.Body).Decode(&schedule)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	result := db.DbConnection.Create(&schedule)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusConflict)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(schedule)
}

func GetCafeSchedule(w http.ResponseWriter, r *http.Request) {
	var schedules []models.Cafeteria
	result := db.DbConnection.Find(&schedules)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(schedules)
}

func GetSingleSchedule(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var Schedule models.Cafeteria
	result := db.DbConnection.First(&Schedule, id)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Schedule)
}

func UpdateCafeSchedule(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var Schedule models.Cafeteria
	result := db.DbConnection.First(&Schedule, id)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusBadRequest)
	}
	err = json.NewDecoder(r.Body).Decode(&Schedule)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	Schedule.ID = uint(id)
	result = db.DbConnection.Save(&Schedule)
	if result.Error != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Schedule)
}

func DeleteCafeSchedule(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var Schedule models.Cafeteria
	result := db.DbConnection.Delete(&Schedule, id)
	if result != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	if result.RowsAffected == 0 {
		http.Error(w, "Schedule not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
