package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	USER := os.Getenv("USER")
	SENHA := os.Getenv("SENHA")
	BD := os.Getenv("DB")
	HOST := os.Getenv("HOST")

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", USER, SENHA, HOST, BD)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic("Erro ao conectar com o banco de dados:", err)
	}
}
