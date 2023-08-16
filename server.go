package main

import (
	"github.com/labstack/echo"
	"github.com/marco-rosner/lightweight-go-server/handlers"
)

func main() {
	e := echo.New()

	e.GET("/", HelloWorld)

	e.POST("/pessoas", handlers.AddPerson)
	e.GET("/pessoas/:id", handlers.GetPerson)
	e.GET("/contagem-pessoas", handlers.GetPeople)

	e.Logger.Fatal(e.Start(":8080"))
}
