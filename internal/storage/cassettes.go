package storage

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/fishmanDK/miet_project/internal/core"
	"github.com/jmoiron/sqlx"
)

type CassettesStorage struct{
	db *sqlx.DB
}

func newCassettesStorage(db *sqlx.DB) *CassettesStorage{
	return &CassettesStorage{
		db: db,
	}
}


func (s *CassettesStorage) GetCassettes() ([]core.Cassette, error){
	query, _, err := sq.Select("*").From("cassettes").ToSql()
	if err != nil{
		return []core.Cassette{}, err
	}

	rows, err := s.db.Query(query)
	if err != nil{
		return []core.Cassette{}, err
	}
	defer rows.Close()

	var cassettes []core.Cassette
	for rows.Next() {
		var cassette core.Cassette

		err := rows.Scan(&cassette)
		if err != nil {
			return []core.Cassette{}, err
		}
		cassettes = append(cassettes, cassette)
	}

	if err := rows.Err(); err != nil {
		return []core.Cassette{}, err
	}

	return cassettes, nil
}

func (s *CassettesStorage) CreateCassette(newCassette core.Cassette) (int, error){
	query := sq.Insert("Cassettes").
	Columns("name", "genre", "year_of_release").
	Values(newCassette.Name, newCassette.Genre, newCassette.Year).
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

func (s *CassettesStorage) CreateCassetteAvailability(newData core.CassetteAvailability) error{
	query := sq.Insert("CassetteAvailability").
	Columns("cassette_id", "store_id", "total_count").
	Values(newData.CassetteId, newData.StoreId, newData.TotalCount).
	RunWith(s.db).
    PlaceholderFormat(sq.Dollar)

	_, err := query.Exec()
	if err != nil{
		return err
	}

	return nil
}