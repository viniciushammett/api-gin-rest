package main

import (
	"gin-api-rest/models"
	"gin-api-rest/routes"
)

func main() {
	models.Alunos = []models.Aluno{
		{Nome: "Vini Teixeira", CPF: "00000000000", RG: "1234657912"},
		{Nome: "Vini Santos", CPF: "00000000000", RG: "1234657914"},
	}
	routes.HandleRequests()
}
