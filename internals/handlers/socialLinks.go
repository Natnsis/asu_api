package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"UniCore/internals/db"
	"UniCore/internals/models"

	"github.com/gorilla/mux"
)

func CreateSocialLinks(w http.ResponseWriter, r *http.Request) {
	var socialLink models.SocialLink
	err := json.NewDecoder(r.Body).Decode(&socialLink)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	result := db.DbConnection.Create(&socialLink)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(socialLink)
}

func GetSocialLinks(w http.ResponseWriter, r *http.Request) {
	var socialLinks models.SocialLink
	result := db.DbConnection.Find(&socialLinks)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(socialLinks)
}

func GetSingleSocailLink(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var socialLink models.SocialLink
	result := db.DbConnection.First(&socialLink, id)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(socialLink)
}

func UpdateSocialLink(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var socialLink models.SocialLink
	result := db.DbConnection.First(&socialLink, id)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusBadRequest)
		return
	}
	socialLink.ID = uint(id)
	result = db.DbConnection.Save(&socialLink)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(socialLink)
}

func DeleteSocailLink(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var socialLink models.SocialLink
	result := db.DbConnection.Delete(&socialLink, id)
	if result != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	if result.RowsAffected == 0 {
		http.Error(w, "Link not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
