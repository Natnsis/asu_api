package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"UniCore/internals/db"
	"UniCore/internals/models"

	"github.com/gorilla/mux"
)

func CreateCurriculum(w http.ResponseWriter, r *http.Request) {
	var curriculum models.Curriculum
	err := json.NewDecoder(r.Body).Decode(&curriculum)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	result := db.DbConnection.Create(&curriculum)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(curriculum)
}

func GetCurriculum(w http.ResponseWriter, r *http.Request) {
	var curriculum models.Curriculum
	result := db.DbConnection.Find(&curriculum)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(curriculum)
}

func GetSingleCurriculum(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var curriculum models.Curriculum
	result := db.DbConnection.First(&curriculum, id)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(curriculum)
}

func UpdateCurriculum(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var curriculum models.Curriculum
	result := db.DbConnection.First(&curriculum, id)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusBadRequest)
		return
	}
	curriculum.ID = uint(id)
	result = db.DbConnection.Save(&curriculum)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(curriculum)
}

func DeleteCurriculum(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var curriculum models.Curriculum
	result := db.DbConnection.Delete(&curriculum, id)
	if result != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	if result.RowsAffected == 0 {
		http.Error(w, "Curriculum not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
