package person

import (
	"fmt"
)

var (
	ErrNotFound = fmt.Errorf("not found")
)

type DB interface {
	Create(*Person) error
	Get(string) (*Person, error)
	Count() (int, error)
}
