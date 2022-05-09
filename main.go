package main

import (
	"gin-api/database"
	"gin-api/routes"
)

func main() {

	database.ConectaComBancoDeDados()

	routes.HandleRequests()
}
