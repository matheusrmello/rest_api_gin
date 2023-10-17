package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/matheusrmello/api-go-gin/controller"
)

func HandleRequests() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")
	r.GET("/alunos", controller.ExibeTodosAlunos)
	r.GET("/alunos/:id", controller.BuscaPorId)
	r.GET("/:nome", controller.Saudacao)
	r.POST("/alunos", controller.CriaNovoAluno)
	r.DELETE("/alunos/:id", controller.DeletaAluno)
	r.PATCH("/alunos/:id", controller.EditaAluno)
	r.GET("/alunos/cpf/:cpf", controller.BuscaPorCPF)
	r.GET("/index", controller.ExibePaginaIndex)
	r.NoRoute(controller.RotaNaoEncontrada)
	r.Run(":8000")
}
