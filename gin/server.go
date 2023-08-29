package main

import (
	"os"

	"github.com/gin-gonic/gin"

	"github.com/marco-rosner/lightweight-go-server/dbs"
	"github.com/marco-rosner/lightweight-go-server/gin/person"
)

func main() {
	server := gin.Default()

	var db dbs.DB

	switch os.Getenv("DB") {
	case "mongo":
		println("Setting mongoDB")
		db = dbs.NewMongoDB()
	case "postgres":
		println("Setting PostgresDB")
		db = dbs.NewPostgresDB()
	default:
		println("Setting InMemoryDB")
		db = dbs.NewInMemDB()
	}

	service := person.PersonService{DB: db}

	server.POST("/pessoas", service.AddPerson)
	server.GET("/pessoas", service.SearchPerson)
	server.GET("/pessoas/:id", service.GetPerson)
	server.GET("/contagem-pessoas", service.CountPeople)

	server.Run("0.0.0.0:8080")
}
