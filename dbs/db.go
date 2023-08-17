package dbs

import (
	"fmt"

	"github.com/marco-rosner/lightweight-go-server/models"
)

var (
	ErrNotFound = fmt.Errorf("not found")
)

type DB interface {
	Create(*models.Person) error
	Get(string) (*models.Person, error)
	Count() (int, error)
}
