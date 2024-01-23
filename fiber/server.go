package main

import (
	"os"

	"github.com/gofiber/fiber/v2"

	"github.com/marco-rosner/lightweight-go-server/dbs"
	"github.com/marco-rosner/lightweight-go-server/fiber/person"
)

func main() {
	f := fiber.New()

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

	f.Post("/pessoas", service.AddPerson)
	f.Get("/pessoas", service.SearchPerson)
	f.Get("/pessoas/:id", service.GetPerson)
	f.Get("/contagem-pessoas", service.CountPeople)

	f.Listen(":8080")
}
