package storage

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
	"github.com/fishmanDK/miet_project/internal/core"
	"github.com/jmoiron/sqlx"
)

type CassettesStorage struct {
	db *sqlx.DB
}

func newCassettesStorage(db *sqlx.DB) *CassettesStorage {
	return &CassettesStorage{
		db: db,
	}
}

func toDatasetListCassettes(id *int) *goqu.SelectDataset {
	selectDataset := goqu.From("cassettes")

	if id == nil {
		return selectDataset
	}

	selectDataset = selectDataset.Select("cassettes.id", "cassettes.name", "cassettes.genre", "cassettes.year_of_release").
		Join(goqu.I("cassetteavailability"),
			goqu.On(cassetteavailabilityField("cassette_id").
				Eq(cassetteField("id")))).
		Where(cassetteavailabilityField("store_id").Eq(id))
	return selectDataset
}

func cassetteField(fieldName string) exp.IdentifierExpression {
	return goqu.I(fmt.Sprintf("%s.%s", "cassettes", fieldName))
}

func cassetteavailabilityField(fieldName string) exp.IdentifierExpression {
	return goqu.I(fmt.Sprintf("%s.%s", "cassetteavailability", fieldName))
}

func (s *CassettesStorage) GetCassettes() ([]core.Cassette, error) {
	query, _, err := toDatasetListCassettes(nil).ToSQL()
	if err != nil {
		return []core.Cassette{}, err
	}

	var cassettes []core.Cassette
	if err = s.db.SelectContext(context.Background(), &cassettes, query); err != nil {
		return nil, fmt.Errorf("list cassettes error during execute query: %w", err)
	}

	return cassettes, nil
}

func (s *CassettesStorage) CreateCassette(newCassette core.CreateCassetteReq) (int, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	query := sq.Insert("Cassettes").
		Columns("name", "genre", "year_of_release").
		Values(newCassette.Name, newCassette.Genre, newCassette.Year).
		Suffix("RETURNING \"id\"").
		RunWith(tx).
		PlaceholderFormat(sq.Dollar)

	var id int
	err = query.QueryRow().Scan(&id)
	if err != nil {
		return 0, err
	}

	fmt.Println(id, newCassette)

	query = sq.Insert("cassetteAvailability").
		Columns("cassette_id", "store_id", "total_count").
		Values(id, newCassette.StoreId, newCassette.TotalCount).
		RunWith(tx).
		PlaceholderFormat(sq.Dollar)

	_, err = query.Exec()
	if err != nil {
		fmt.Println(2, err.Error())
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *CassettesStorage) GetCassettesByStoreID(id int) ([]core.Cassette, error) {
	query, _, err := toDatasetListCassettes(&id).ToSQL()
	if err != nil {
		return []core.Cassette{}, err
	}

	var cassettes []core.Cassette
	if err = s.db.SelectContext(context.Background(), &cassettes, query); err != nil {
		return nil, fmt.Errorf("list cassettes error during execute query: %w", err)
	}

	return cassettes, nil
}

func (s *CassettesStorage) CreateCassetteAvailability(newData core.CassetteAvailability) error {
	query := sq.Insert("CassetteAvailability").
		Columns("cassette_id", "store_id", "total_count").
		Values(newData.CassetteId, newData.StoreId, newData.TotalCount).
		RunWith(s.db).
		PlaceholderFormat(sq.Dollar)

	_, err := query.Exec()
	if err != nil {
		return err
	}

	return nil
}

func (s *CassettesStorage) SaveCassetteChanges(changes core.ChangeCassette) error {
	tx, err := s.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	query := sq.Update("cassetteAvailability").
		Set("total_count", changes.TotalCount).
		Where(sq.Eq{"cassette_id": changes.CassetteID}).
		PlaceholderFormat(sq.Dollar)

	_, err = query.RunWith(tx).Exec()
	if err != nil {
		return fmt.Errorf("failed to update cassette availability: %w", err)
	}

	if changes.Name != "" || changes.Ganre != "" {
		updateCassetteQuery := sq.Update("Cassettes").
			Set("name", changes.Name).
			Set("genre", changes.Ganre).
			Where(sq.Eq{"id": changes.CassetteID}).
			PlaceholderFormat(sq.Dollar)

		// Выполняем обновление кассеты
		_, err = updateCassetteQuery.RunWith(tx).Exec() // Выполняем обновление кассеты
		if err != nil {
			return fmt.Errorf("failed to update cassette: %w", err)
		}
	}

	// Коммит транзакции
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (s *CassettesStorage) GetCassette(cassetteID int) (core.Cassette, error) {
	query := sq.Select("c.name", "c.genre", "c.year_of_release", "ca.total_count", "ca.rented_count").
		From("cassettes c").
		Join("cassetteavailability ca ON c.id = ca.cassette_id").
		Where(sq.Eq{"c.id": cassetteID}).
		RunWith(s.db).
		PlaceholderFormat(sq.Dollar)

	var cassette core.Cassette
	err := query.QueryRow().Scan(&cassette.Name, &cassette.Genre, &cassette.Year, &cassette.TotalCount, &cassette.RentedCount)
	fmt.Println(cassette)
	return cassette, err
}

func (s *CassettesStorage) GetCassetteDetails(cassetteID, userID int) (core.CassetteAvailability, error) {
	// Запрос на получение информации о кассете (total_count и rented_count)
	query := sq.Select("ca.total_count", "ca.rented_count").
		From("cassetteavailability ca").
		Join("cassettes c ON ca.cassette_id = c.id").
		Where(sq.Eq{"ca.cassette_id": cassetteID}).
		RunWith(s.db).
		PlaceholderFormat(sq.Dollar)

	var res core.CassetteAvailability
	err := query.QueryRow().Scan(&res.TotalCount, &res.RentedCount)
	if err != nil {
		return core.CassetteAvailability{}, err
	}

	// Проверка на наличие заказа
	queryCheckOrder := sq.Select("1").
		From("orders").
		Where(sq.Eq{"cassette_id": cassetteID}).
		Where(sq.Eq{"user_id": userID}).
		RunWith(s.db).
		PlaceholderFormat(sq.Dollar)

	var orderExists int
	errCheckOrder := queryCheckOrder.QueryRow().Scan(&orderExists)

	// Если произошла ошибка, кроме migrations.ErrNoRows, возвращаем ошибку
	if errCheckOrder != nil && errCheckOrder != sql.ErrNoRows {
		return res, fmt.Errorf("error querying order status: %v", errCheckOrder)
	}

	// Если запись существует, устанавливаем флаг IsOrdered = true
	if errCheckOrder == nil {
		res.IsOrdered = true
	}

	// Проверка на наличие бронирования
	queryCheckReservation := sq.Select("1").
		From("reserve_pool").
		Where(sq.Eq{"cassette_id": cassetteID}).
		Where(sq.Eq{"user_id": userID}).
		RunWith(s.db).
		PlaceholderFormat(sq.Dollar)

	var reservationExists int
	errCheckReservation := queryCheckReservation.QueryRow().Scan(&reservationExists)

	// Если произошла ошибка, кроме migrations.ErrNoRows, возвращаем ошибку
	if errCheckReservation != nil && errCheckReservation != sql.ErrNoRows {
		return res, fmt.Errorf("error querying reservation status: %v", errCheckReservation)
	}

	// Если запись существует, устанавливаем флаг IsReservated = true
	if errCheckReservation == nil {
		res.IsReservated = true
	}

	// Возвращаем результат
	return res, nil
}

func (s *CassettesStorage) DeleteCasseteByID(cassetteID int) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	query := sq.Delete("cassetteavailability").
		Where(sq.Eq{"cassette_id": cassetteID}).
		RunWith(tx).
		PlaceholderFormat(sq.Dollar)

	sql, args, err := query.ToSql() // Получаем SQL и аргументы
	if err != nil {
		return err
	}
	log.Printf("Executing query: %s, args: %v", sql, args) // Логируем запрос

	_, err = query.Exec()
	if err != nil {
		return err
	}

	_, err = query.Exec()
	if err != nil {
		return err
	}

	query = sq.Delete("cassettes").
		Where(sq.Eq{"id": cassetteID}).
		RunWith(tx).
		PlaceholderFormat(sq.Dollar)

	sql, args, err = query.ToSql() // Получаем SQL и аргументы
	if err != nil {
		return err
	}
	log.Printf("Executing query: %s, args: %v", sql, args) // Логируем запрос

	_, err = query.Exec()
	if err != nil {
		return err
	}
	_, err = query.Exec()
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
