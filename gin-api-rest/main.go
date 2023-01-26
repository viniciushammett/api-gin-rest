package main

import (
	"gin-api-rest/models"
	"gin-api-rest/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
