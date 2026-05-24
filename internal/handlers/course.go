package handlers

import (
	"encoding/json"
	"net/http"

	"unicore/internal/db"
	"unicore/internal/models/courses"
)

func CreateCourse(w http.ResponseWriter, r *http.Request) {
	var course courses.Courses
	err := json.NewDecoder(r.Body).Decode(&course)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	result := db.Db.Create(&course)
	if result.Error != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Context-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&course)
}
