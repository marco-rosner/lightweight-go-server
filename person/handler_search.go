package person

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/marco-rosner/lightweight-go-server/dbs"
)

func (pDB PersonService) SearchPerson(c echo.Context) error {
	term := c.QueryParam("t")

	if term == "" {
		return echo.ErrBadRequest
	}

	p, err := pDB.DB.Search(term)

	switch err {
	case nil:
		return c.JSON(http.StatusOK, p)
	case dbs.ErrNotFound:
		return echo.ErrNotFound
	default:
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
}
