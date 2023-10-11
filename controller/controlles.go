package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/matheusrmello/api-go-gin/database"
	"github.com/matheusrmello/api-go-gin/models"
)

func ExibeTodosAlunos(c *gin.Context) {
	var alunos []models.AlunosDados
	database.DB.Find(&alunos)
	c.JSON(200, alunos)
}

func Saudacao(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"API diz:": "E ai " + nome + ", tudo beleza",
	})
}

func CriaNovoAluno(c *gin.Context) {
	var aluno models.AlunosDados
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return

	}
	database.DB.Create(&aluno)
	c.JSON(http.StatusOK, aluno)
}

func BuscaPorId(c *gin.Context) {
	var aluno models.AlunosDados
	id := c.Params.ByName("id")
	database.DB.First(&aluno, id)
	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Aluno nao encontrado"})
		return
	}

	c.JSON(http.StatusOK, aluno)
}

func DeletaAluno(c *gin.Context) {
	var aluno models.AlunosDados
	id := c.Params.ByName("id")
	database.DB.Delete(&aluno, id)
	c.JSON(http.StatusOK, gin.H{"data": "Aluno deletado com sucesso"})
}

func EditaAluno(c *gin.Context) {
	var aluno models.AlunosDados
	id := c.Params.ByName("id")
	database.DB.First(&aluno, id)
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return

	}
	database.DB.Model(&aluno).UpdateColumns(aluno)
	c.JSON(http.StatusOK, aluno)

}

func BuscaPorCPF(c *gin.Context) {
	var aluno models.AlunosDados
	cpf := c.Param("cpf")
	database.DB.Where(&models.AlunosDados{CPF: cpf}).First(&aluno)
	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Aluno nao encontrado"})
		return
	}
	c.JSON(http.StatusOK, aluno)

}
