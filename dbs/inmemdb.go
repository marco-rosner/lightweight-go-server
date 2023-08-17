package dbs

import "github.com/marco-rosner/lightweight-go-server/models"

var people = map[string]*models.Person{
	"1": {ID: "1", Name: "Marco Rosner", Nickname: "Rosner", Birth: "2000-10-01"},
	"2": {ID: "2", Name: "João", Nickname: "João", Birth: "2000-10-02"},
	"3": {ID: "3", Name: "Maria", Nickname: "Maria", Birth: "2000-10-03"},
}

type InMemDB struct {
	data map[string]*models.Person
}

func NewInMemDB() *InMemDB {
	return &InMemDB{
		data: people,
	}
}

func (db *InMemDB) Create(p *models.Person) error {
	db.data[p.ID] = p
	return nil
}

func (db *InMemDB) Get(id string) (*models.Person, error) {
	p, ok := db.data[id]

	if !ok {
		return nil, ErrNotFound
	}

	return p, nil
}

func (db *InMemDB) Count() (int, error) {
	return len(db.data), nil
}
