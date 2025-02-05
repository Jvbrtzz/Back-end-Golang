package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Jvbrtzz/Back-end-golang/database"
	"github.com/Jvbrtzz/Back-end-golang/models"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
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

func RegisterNewUser(w http.ResponseWriter, r *http.Request) {
	var novaUser models.User
	if err := json.NewDecoder(r.Body).Decode(&novaUser); err != nil {
		log.Printf("Erro ao decodificar o corpo da requisição: %v", err)
		http.Error(w, "Dados inválidos", http.StatusBadRequest)
		return
	}

	// Criptografar a senha do usuário
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(novaUser.Senha), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Erro ao criptografar a senha: %v", err)
		http.Error(w, "Erro interno ao processar a senha", http.StatusInternalServerError)
		return
	}
	novaUser.Senha = string(hashedPassword)

	if err := database.DB.Create(&novaUser).Error; err != nil {
		log.Printf("Erro ao criar o novo usuário no banco de dados: %v", err)
		http.Error(w, "Erro ao registrar o usuário", http.StatusInternalServerError)
		return
	}

	log.Printf("Usuário registrado com sucesso: %v", novaUser.Id)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(novaUser)
}
