package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"UniCore/internals/db"
	"UniCore/internals/models"

	"github.com/gorilla/mux"
)

func CreateUniversity(w http.ResponseWriter, r *http.Request) {
	var university models.University
	err := json.NewDecoder(r.Body).Decode(&university)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	result := db.DbConnection.Create(&university)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(university)
}

func GetUniversity(w http.ResponseWriter, r *http.Request) {
	var university models.University
	result := db.DbConnection.Find(&university)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(university)
}

func GetSingleUniversity(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var university models.University
	result := db.DbConnection.First(&university, id)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(university)
}

func UpdateUniversity(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var university models.University
	result := db.DbConnection.First(&university, id)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusBadRequest)
		return
	}
	university.ID = uint(id)
	result = db.DbConnection.Save(&university)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(university)
}

func DeleteUniversity(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var university models.University
	result := db.DbConnection.Delete(&university, id)
	if result != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	if result.RowsAffected == 0 {
		http.Error(w, "University not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
