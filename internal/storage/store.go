package storage

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/fishmanDK/miet_project/internal/core"
	"github.com/jmoiron/sqlx"
)

type StoreStorage struct{
	db *sqlx.DB
}

func newStoreStorage(db *sqlx.DB) *StoreStorage{
	return &StoreStorage{
		db: db,
	}
}

func (s *StoreStorage) CreateStore(newStore core.Store) (int, error){
	query := sq.Insert("Stores").
	Columns("address").
	Values(newStore.Address).
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