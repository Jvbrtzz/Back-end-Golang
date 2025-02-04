package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Jvbrtzz/Back-end-golang/database"
	"github.com/Jvbrtzz/Back-end-golang/models"
	"github.com/gorilla/mux"
)

func GetUserCard(w http.ResponseWriter, r *http.Request) {
	var cards []models.Card

	vars := mux.Vars(r)
	id := vars["id"]

	// Buscar todos os cards associados ao usuário
	if err := database.DB.Where("user_id = ?", id).Find(&cards).Error; err != nil {
		http.Error(w, "Cards não encontrados", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(cards)
}

func GetUserCardComments(w http.ResponseWriter, r *http.Request) {
	var comment []models.Comment

	vars := mux.Vars(r)
	id := vars["id"]

	// Buscar todos os cards associados ao usuário
	if err := database.DB.Where("card_id=?", id).Find(&comment).Error; err != nil {
		http.Error(w, "Cards não encontrados", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(comment)
}
