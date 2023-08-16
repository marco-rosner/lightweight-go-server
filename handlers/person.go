package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

type Person struct {
	ID       string    `json:"id"`
	Name     string    `json:"name" binding:"required,len=32"`
	Nickname string    `json:"nickname" binding:"required,len=100"`
	Birth    time.Time `json:"birh" binding:"required"`
}

var people = []Person{
	{ID: "1", Name: "Marco Rosner", Nickname: "Rosner", Birth: time.Now()},
	{ID: "2", Name: "João", Nickname: "João", Birth: time.Now()},
	{ID: "3", Name: "Maria", Nickname: "Maria", Birth: time.Now()},
}

func AddPerson(c echo.Context) error {
	var person Person

	if err := c.Bind(&person); err != nil {
		c.JSON(http.StatusBadRequest, "ValidationError: "+err.Error())
	}

	jsonPerson, err := json.Marshal(person)

	if err != nil {
		c.JSON(http.StatusBadRequest, "Bad Request")
	}

	people = append(people, person)

	return c.JSON(http.StatusCreated, string(jsonPerson))
}

func GetPeople(c echo.Context) error {
	return c.JSON(http.StatusOK, len(people))
}

func GetPerson(c echo.Context) error {
	id := c.Param("id")

	for _, person := range people {
		if person.ID == id {
			c.JSON(http.StatusOK, person)
			return nil
		}
	}

	return c.JSON(http.StatusNotFound, "Person not found")
}
