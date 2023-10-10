package main

import (
	"github.com/matheusrmello/api-go-gin/database"
	"github.com/matheusrmello/api-go-gin/models"
	"github.com/matheusrmello/api-go-gin/routes"
)

func main() {
	database.ConnectionDB()
	models.Alunos = []models.AlunosDados{
		{Nome: "Matheus Mello", CPF: "12345678903", RG: "318156635"},
		{Nome: "Vitoria Sirico", CPF: "78095070017", RG: "229316384"},
		{Nome: "Maria Luiza", CPF: "05448214088", RG: "182618080"},
		{Nome: "Joao Qualquer Coisa", CPF: "62413773029", RG: "425689219"},
		{Nome: "Joana Das Neves", CPF: "99287262055", RG: "484607765"},
	}
	routes.HandleRequests()
}
