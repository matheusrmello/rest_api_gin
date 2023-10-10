package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/matheusrmello/api-go-gin/controller"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/alunos", controller.ExibeTodosAlunos)
	r.GET("/:nome", controller.Saudacao)
	r.Run(":8000")
}
