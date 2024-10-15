package storage

import (
	sq "github.com/Masterminds/squirrel"

	"github.com/fishmanDK/miet_project/internal/core"
	"github.com/jmoiron/sqlx"
)

type ClientStorage struct{
	db *sqlx.DB
}

func newClientStorage(db *sqlx.DB) *ClientStorage{
	return &ClientStorage{
		db: db,
	}
}

func (s *ClientStorage) CreateClient(newClient core.Client) (int, error){
	query := sq.Insert("Clients").
	Columns("name", "email", "password").
	Values(newClient.Name, newClient.Email, newClient.Password).
	Suffix("RETURNING \"id\"").
	RunWith(s.db).
    PlaceholderFormat(sq.Dollar)

	var id int
	err := query.QueryRow().Scan(&id)
	if err != nil{
		return 0, err
	}

	return id, nil
}