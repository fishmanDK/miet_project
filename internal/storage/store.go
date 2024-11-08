package storage

import (
	"fmt"

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

func (s *StoreStorage) GetStores() ([]core.Store, error) {
	query := sq.Select("id", "address").From("Stores").RunWith(s.db)
	fmt.Println(query.ToSql())
	rows, err := query.Query()
	if err != nil{
		return []core.Store{}, err
	}
	defer rows.Close()

	var res []core.Store
	for rows.Next(){
		var store core.Store
		err := rows.Scan(&store.Id, &store.Address)
		if err != nil {
			return []core.Store{}, err
		}
		
		res = append(res, store)
	}

	return res, nil
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