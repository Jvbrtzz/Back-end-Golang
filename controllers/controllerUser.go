package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Jvbrtzz/Back-end-golang/database"
	"github.com/Jvbrtzz/Back-end-golang/models"
	"github.com/gorilla/mux"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	var p []models.User
	database.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
}

func GetUsersById(w http.ResponseWriter, r *http.Request) {
	var user models.User

	vars := mux.Vars(r)
	id := vars["id"]

	if err := database.DB.First(&user, id).Error; err != nil {
		http.Error(w, "Aluno não encontrado", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}
