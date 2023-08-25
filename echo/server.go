package main

import (
	"github.com/labstack/echo"
	"github.com/marco-rosner/lightweight-go-server/dbs"
	"github.com/marco-rosner/lightweight-go-server/echo/person"
)

func main() {
	e := echo.New()

	// db := dbs.NewInMemDB()
	// db := dbs.NewMongoDB()
	db := dbs.NewPostgresDB()

	service := person.PersonService{DB: db}

	e.POST("/pessoas", service.AddPerson)
	e.GET("/pessoas", service.SearchPerson)
	e.GET("/pessoas/:id", service.GetPerson)
	e.GET("/contagem-pessoas", service.CountPeople)

	e.Logger.Fatal(e.Start(":8080"))
}
