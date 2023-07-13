package main

import (
	"fmt"

	"github.com/manbomb/go-curso-api-rest/database"
	"github.com/manbomb/go-curso-api-rest/models"
	"github.com/manbomb/go-curso-api-rest/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Carmela Dutra", Historia: "Uma mulher muito famosa"},
		{Id: 2, Nome: "Maria da Penha", Historia: "Uma mulher muito importante"},
	}
	database.ConectaComBancoDeDados()
	fmt.Println("Iniciando servidor Rest com Go")
	routes.HandleRequest()
}
