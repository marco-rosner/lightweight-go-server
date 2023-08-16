package main

import (
	"github.com/labstack/echo"
	"github.com/marco-rosner/lightweight-go-server/person"
)

func main() {
	e := echo.New()

	e.GET("/", HelloWorld)

	var db = person.NewInMemDB()

	e.POST("/pessoas", person.PersonDB{DB: db}.AddPerson)
	e.GET("/pessoas/:id", person.PersonDB{DB: db}.GetPerson)
	e.GET("/contagem-pessoas", person.PersonDB{DB: db}.CountPeople)

	e.Logger.Fatal(e.Start(":8080"))
}
