package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"

	"testing"

	"github.com/gin-gonic/gin"
	"github.com/matheusrmello/api-go-gin/controller"
	"github.com/matheusrmello/api-go-gin/database"
	"github.com/matheusrmello/api-go-gin/models"
	"github.com/stretchr/testify/assert"
)

var ID int

func SetupDasRotasDeTest() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	rotas := gin.Default()
	return rotas

}
func CriaAlunoMock()  {
	aluno := models.AlunosDados{Nome: "Nome do aluno teste", CPF: "12345678912", RG: "123456789"}
	database.DB.Create(&aluno)
	ID = int(aluno.ID)
}

func DeletaAlunoMock()  {
	var aluno models.AlunosDados
	database.DB.Delete(&aluno, ID)
}

func TestVerificaStatusCode(t *testing.T) {
	r := SetupDasRotasDeTest()
	r.GET("/:nome", controller.Saudacao)
	req, _ := http.NewRequest("GET", "/", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code, "Deveriam ser iguais")
	mockDaResposta := `{"API diz":"E ai matheus, tudo beleza"}`
	respostaBody, _ := io.ReadAll(resposta.Body)
	assert.Equal(t, mockDaResposta, string(respostaBody))
	fmt.Println(string(respostaBody))
	fmt.Println(mockDaResposta)
}

func TestVerificaTodosOsAlunos(t *testing.T) {
	database.ConnectionDB()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupDasRotasDeTest()
	r.GET("/alunos", controller.ExibeTodosAlunos)
	req, _ := http.NewRequest("GET", "/alunos", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
	fmt.Println(resposta.Body)

}

func TestBuscaCPFHandler(t *testing.T)  {
	database.ConnectionDB()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupDasRotasDeTest()
	r.GET("/alunos/cpf/:cpf", controller.BuscaPorCPF)
	req, _ := http.NewRequest("GET", "/alunos/cpf/12345678912", nil)
	resposta :=httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestBuscaPorIDHandler(t *testing.T)  {
	database.ConnectionDB()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupDasRotasDeTest()
	r.GET("/alunos/:id", controller.BuscaPorId)
	path := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("GET", path, nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	var nomeDoAlunoMock models.AlunosDados
	json.Unmarshal(resposta.Body.Bytes(), &nomeDoAlunoMock)
	fmt.Println(nomeDoAlunoMock.Nome)
	assert.Equal(t, "Nome do aluno teste", nomeDoAlunoMock.Nome, "Os nomes devem ser iguais")
	assert.Equal(t, "12345678912", nomeDoAlunoMock.CPF)
	assert.Equal(t, "123456789", nomeDoAlunoMock.RG)
	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestDeletaAlunoHandler(t *testing.T)  {
	database.ConnectionDB()
	CriaAlunoMock()
	r := SetupDasRotasDeTest()
	r.DELETE("/alunos/:id", controller.DeletaAluno)
	pathDeBusca := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("DELETE", pathDeBusca, nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestAtualizaAlunoHandler(t *testing.T)  {
	database.ConnectionDB()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupDasRotasDeTest()
	r.PATCH("/alunos/:id", controller.EditaAluno)
	aluno := models.AlunosDados{Nome: "Nome do aluno teste", CPF: "46345678914", RG: "453456780"}
	valorJson, _ := json.Marshal(aluno)
	pathParaEditar := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("PATCH", pathParaEditar, bytes.NewBuffer(valorJson))
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	var alunoMockAtualizado models.AlunosDados
	json.Unmarshal(resposta.Body.Bytes(), &alunoMockAtualizado)
	assert.Equal(t, "46345678914", alunoMockAtualizado.CPF)
	assert.Equal(t, "453456780", alunoMockAtualizado.RG)
	assert.Equal(t, "Nome do aluno teste", alunoMockAtualizado.Nome)

}

