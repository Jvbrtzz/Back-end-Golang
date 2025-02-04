package main

import (
	"fmt"

	"github.com/Jvbrtzz/Back-end-golang/database"
	"github.com/Jvbrtzz/Back-end-golang/routes"
)

func main() {
	// Conectar ao banco de dados
	database.ConectaComBancoDeDados()

	// Iniciar o servidor
	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleRequest()

}
