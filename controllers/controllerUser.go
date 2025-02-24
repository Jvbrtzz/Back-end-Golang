package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/Jvbrtzz/Back-end-golang/database"
	"github.com/Jvbrtzz/Back-end-golang/models"
	"github.com/golang-jwt/jwt/v4"
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
	json.NewEncoder(w).Encode(novaUser)
}

var jwtSecret = []byte("access-token")

func UserLogin(w http.ResponseWriter, r *http.Request) {

	var credenciais models.User
	if err := json.NewDecoder(r.Body).Decode(&credenciais); err != nil {
		log.Printf("Erro ao decodificar a requisição: %v", err)
		http.Error(w, "Dados inválidos", http.StatusBadRequest)
		return
	}

	var user models.User
	if err := database.DB.Where("nome = ?", credenciais.Nome).First(&user).Error; err != nil {
		log.Printf("Usuário não encontrado: %v", credenciais.Nome)
		http.Error(w, "Usuário ou senha inválidos", http.StatusUnauthorized)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Senha), []byte(credenciais.Senha)); err != nil {
		log.Printf("Falha na autenticação para o usuário: %v", credenciais.Nome)
		http.Error(w, "Usuário ou senha inválidos", http.StatusUnauthorized)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  user.Id,
		"username": user.Nome,
		"exp":      time.Now().Add(time.Hour * 360).Unix(),
	})
	tokenString, err := token.SignedString(jwtSecret)

	if err != nil {
		log.Println("Erro ao gerar token JWT:", err)
		http.Error(w, "Erro ao gerar token de autenticação", http.StatusInternalServerError)
		return
	}

	response := models.LoginResponse{
		Success:     true,
		Message:     "Login successful",
		User:        user,
		AccessToken: tokenString,
	}

	log.Printf("Usuário autenticado com sucesso: %v", user.Id)
	json.NewEncoder(w).Encode(response)
}
