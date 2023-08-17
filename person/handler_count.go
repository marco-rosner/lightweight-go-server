package person

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/marco-rosner/lightweight-go-server/dbs"
)

func (pDB PersonService) CountPeople(c echo.Context) error {
	num, err := pDB.DB.Count()

	if err != nil {
		c.JSON(http.StatusNotFound, dbs.ErrNotFound)
	}

	return c.JSON(http.StatusOK, num)
}
