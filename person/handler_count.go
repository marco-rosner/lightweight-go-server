package person

import (
	"net/http"

	"github.com/labstack/echo"
)

func (pDB PersonDB) CountPeople(c echo.Context) error {
	num, err := pDB.DB.Count()

	if err != nil {
		c.JSON(http.StatusNotFound, ErrNotFound)
	}

	return c.JSON(http.StatusOK, num)
}
