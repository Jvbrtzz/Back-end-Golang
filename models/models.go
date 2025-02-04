package models

type User struct {
	Id    int    `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
	Senha string `json:"senha"`
}

var Users []User

type Card struct {
	Id          int    `json:"id"`
	User_id     string `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

var Cards []Card

type Comment struct {
	Id      int    `json:"id"`
	Card_id string `json:"card_id"`
	User_id string `json:"user_id"`
	Comment string `json:"comment"`
}

var Comments []Comment

func (User) TableName() string {
	return "user" // nome correto da tabela
}
