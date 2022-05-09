package routes

import (
	"gin-api/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.LoadHTMLGlob("views/*")
	r.Static("/static", "./static")
	r.GET("/:nome", controllers.Saudacao)
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	r.GET("/alunos/:id", controllers.BuscaAlunoPorId)
	r.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoPorCPF)
	r.POST("/alunos", controllers.CriaNovoAluno)
	r.PUT("/alunos/:id", controllers.EditaAluno)
	r.DELETE("/alunos/:id", controllers.DeletaAluno)
	r.GET("/index", controllers.ExibePaginaIndex)
	r.NoRoute(controllers.RotaNaoEncontrada)
	r.Run()
}
