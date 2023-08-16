package person

var people = map[string]*Person{
	"1": {ID: "1", Name: "Marco Rosner", Nickname: "Rosner", Birth: "2000-10-01"},
	"2": {ID: "2", Name: "João", Nickname: "João", Birth: "2000-10-02"},
	"3": {ID: "3", Name: "Maria", Nickname: "Maria", Birth: "2000-10-03"},
}

type InMemDB struct {
	data map[string]*Person
}

func NewInMemDB() *InMemDB {
	return &InMemDB{
		data: people,
	}
}

func (db *InMemDB) Create(p *Person) error {
	db.data[p.ID] = p
	return nil
}

func (db *InMemDB) Get(id string) (*Person, error) {
	p, ok := db.data[id]

	if !ok {
		return nil, ErrNotFound
	}

	return p, nil
}

func (db *InMemDB) Count() (int, error) {
	return len(db.data), nil
}
