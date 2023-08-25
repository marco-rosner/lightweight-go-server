package main

import (
	"github.com/gofiber/fiber/v2"

	"github.com/marco-rosner/lightweight-go-server/dbs"
	"github.com/marco-rosner/lightweight-go-server/fiber/person"
)

func main() {
	f := fiber.New()

	// db := dbs.NewInMemDB()
	// db := dbs.NewMongoDB()
	db := dbs.NewPostgresDB()

	service := person.PersonService{DB: db}

	f.Post("/pessoas", service.AddPerson)
	f.Get("/pessoas", service.SearchPerson)
	f.Get("/pessoas/:id", service.GetPerson)
	f.Get("/contagem-pessoas", service.CountPeople)

	f.Listen(":8080")
}
