package main

import (
	"github.com/gin-gonic/gin"

	"github.com/marco-rosner/lightweight-go-server/dbs"
	"github.com/marco-rosner/lightweight-go-server/gin/person"
)

func main() {
	server := gin.Default()

	// db := dbs.NewInMemDB()
	// db := dbs.NewMongoDB()
	db := dbs.NewPostgresDB()

	service := person.PersonService{DB: db}

	server.POST("/pessoas", service.AddPerson)
	server.GET("/pessoas", service.SearchPerson)
	server.GET("/pessoas/:id", service.GetPerson)
	server.GET("/contagem-pessoas", service.CountPeople)

	server.Run("localhost:8080")
}
