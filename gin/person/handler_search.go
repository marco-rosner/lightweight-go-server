package person

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marco-rosner/lightweight-go-server/dbs"
)

func (pDB PersonService) SearchPerson(c *gin.Context) {
	term := c.Query("t")

	if term == "" {
		c.Error(errors.New("Error: "))
	}

	p, err := pDB.DB.Search(term)

	switch err {
	case nil:
		c.JSON(http.StatusOK, p)
	case dbs.ErrNotFound:
		c.JSON(http.StatusNotFound, err.Error())
	default:
		c.JSON(http.StatusInternalServerError, err.Error())
	}
}
