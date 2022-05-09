package main

import (
	"bytes"
	"encoding/json"
	"gin-api/controllers"
	"gin-api/database"
	"gin-api/models"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var ID int

func ConfigRotasDeTeste() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	rotas := gin.Default()

	return rotas
}

func CriarAlunoMock() {
	aluno := models.Aluno{Nome: "Aquaman", CPF: "12345678901", RG: "123456789"}
	database.DB.Create(&aluno)
	ID = int(aluno.ID)
}

func DeletaAlunoMock() {
	var aluno models.Aluno
	database.DB.Delete(&aluno, ID)
}

func TestStatusCodeSaudacao(t *testing.T) {
	r := ConfigRotasDeTeste()
	r.GET("/:nome", controllers.Saudacao)
	req, _ := http.NewRequest("GET", "/Ícaro", nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code, "Igualdade aguardada")
	mockDaResposta := `{"API diz:":"Bem vindo Ícaro"}`
	respostBody, _ := ioutil.ReadAll(res.Body)
	assert.Equal(t, mockDaResposta, string(respostBody))

}

func TestGetTodosALunos(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriarAlunoMock()
	defer DeletaAlunoMock()
	r := ConfigRotasDeTeste()
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	req, _ := http.NewRequest("GET", "/alunos", nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)

}

func TestBucaAlunoPorCPF(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriarAlunoMock()
	defer DeletaAlunoMock()
	r := ConfigRotasDeTeste()
	r.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoPorCPF)
	req, _ := http.NewRequest("GET", "/alunos/cpf/12345678901", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestBuscaAlunoPorId(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriarAlunoMock()
	//defer DeletaAlunoMock()
	r := ConfigRotasDeTeste()
	r.GET("/alunos/:id", controllers.BuscaAlunoPorId)
	pathDaBusca := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("GET", pathDaBusca, nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)
	var alunoMock models.Aluno
	json.Unmarshal(res.Body.Bytes(), &alunoMock)
	assert.Equal(t, "Aquaman", alunoMock.Nome, "Os nomes devem ser iguais")
	assert.Equal(t, "12345678901", alunoMock.CPF)
	assert.Equal(t, "123456789", alunoMock.RG)
	assert.Equal(t, http.StatusOK, res.Code)
}

func TestDeletaAluno(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriarAlunoMock()
	r := ConfigRotasDeTeste()
	r.DELETE("/alunos/:id", controllers.DeletaAluno)
	pathDeBusca := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("DELETE", pathDeBusca, nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestEditaAluno(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriarAlunoMock()
	defer DeletaAlunoMock()
	r := ConfigRotasDeTeste()
	r.PUT("/alunos/:id", controllers.EditaAluno)
	aluno := models.Aluno{Nome: "Nome do Aluno Teste", CPF: "47123456789", RG: "123456700"}
	valorJson, _ := json.Marshal(aluno)
	path := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("PUT", path, bytes.NewBuffer(valorJson))
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)
	var alunoMockAtualizado models.Aluno
	json.Unmarshal(res.Body.Bytes(), &alunoMockAtualizado)
	assert.Equal(t, "47123456789", alunoMockAtualizado.CPF)
	assert.Equal(t, "123456700", alunoMockAtualizado.RG)
	assert.Equal(t, "Nome do Aluno Teste", alunoMockAtualizado.Nome)
}
