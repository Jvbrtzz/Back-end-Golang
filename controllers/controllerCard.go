package controllers

import (
	"encoding/json"
	"log"
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

func RegisterNewCard(w http.ResponseWriter, r *http.Request) {
	var novoCard models.Card
	if err := json.NewDecoder(r.Body).Decode(&novoCard); err != nil {
		log.Printf("Erro ao decodificar o corpo da requisição: %v", err)
		http.Error(w, "Dados inválidos", http.StatusBadRequest)
		return
	}

	if err := database.DB.Create(&novoCard).Error; err != nil {
		log.Printf("Erro ao criar o novo usuário no banco de dados: %v", err)
		http.Error(w, "Erro ao registrar o usuário", http.StatusInternalServerError)
		return
	}

	log.Printf("Usuário registrado com sucesso: %v", novoCard.Id)
	json.NewEncoder(w).Encode(novoCard)
}
