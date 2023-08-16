package person

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo"
)

func (pDB PersonDB) AddPerson(c echo.Context) error {
	var person Person

	if err := c.Bind(&person); err != nil {
		c.JSON(http.StatusBadRequest, "ValidationError: "+err.Error())
	}

	jsonPerson, err := json.Marshal(person)

	if err != nil {
		c.JSON(http.StatusBadRequest, "Bad Request")
	}

	// people = append(people, person)

	pDB.DB.Create(&person)

	return c.JSON(http.StatusCreated, string(jsonPerson))
}
