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
	User_id     int    `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

var Cards []Card

type Comment struct {
	Id      int    `json:"id" gorm:"primaryKey"`
	Card_id string `json:"card_id"`
	User_id string `json:"user_id"`
	Comment string `json:"comment"`
}

var Comments []Comment

func (User) TableName() string {
	return "user" // nome correto da tabela
}

type LoginResponse struct {
	Success     bool   `json:"success"`
	Message     string `json:"message"`
	User        User   `json:"user"`
	AccessToken string `json:"accessToken"`
}

type CardUsers struct {
	UserID     uint   `json:"user_id"`
	CardID     uint   `json:"card_id"`
	Permission string `json:"user_permission"`

	// Adicionando o relacionamento com a struct User
	User User `json:"user" gorm:"foreignKey:UserID;references:Id"`
}
