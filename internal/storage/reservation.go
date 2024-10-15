package storage

import (
	"database/sql"
	"fmt"
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/fishmanDK/miet_project/internal/core"
	"github.com/jmoiron/sqlx"
)

type ReservationStorage struct {
	db *sqlx.DB
}

func newReservationStorage(db *sqlx.DB) *ReservationStorage {
	return &ReservationStorage{
		db: db,
	}
}

func (s *ReservationStorage) CreateReservation(newReservate core.Reservation) (int, error) {
    tx, err := s.db.Begin()
    if err != nil {
        return 0, err
    }
    defer tx.Rollback()

    // Вставка новой резервации
    query := sq.Insert("Reservations").
        Columns("client_id", "cassette_id", "store_id", "status").
        Values(newReservate.ClientId, newReservate.CassetteId, newReservate.StoreId, "reserve").
        Suffix("RETURNING \"id\"").
        RunWith(tx). // Используем tx вместо s.db для работы внутри транзакции
        PlaceholderFormat(sq.Dollar)

    var id int
    err = query.QueryRow().Scan(&id)
    if err != nil {
        return 0, err
    }

    // Проверка наличия кассеты и значения total_count
    checkQuery := sq.Select("total_count").
        From("CassetteAvailability").
        Where(sq.Eq{"cassette_id": newReservate.CassetteId}).
        Where(sq.Eq{"store_id": newReservate.StoreId}).
        RunWith(tx). // Используем tx для выполнения внутри транзакции
        PlaceholderFormat(sq.Dollar)

    var totalCount int
    err = checkQuery.QueryRow().Scan(&totalCount)
    if err != nil {
        if err == sql.ErrNoRows {
            // Запись не найдена, возвращаем ошибку
            return 0, fmt.Errorf("CassetteAvailability not found")
        }
        return 0, fmt.Errorf("failed to check CassetteAvailability: %w", err)
    }

    // Если total_count равно 0, возвращаем ошибку
    if totalCount == 0 {
        return 0, fmt.Errorf("total_count cannot be zero")
    }

    // Обновление наличия кассеты
    updateQuery := sq.Update("CassetteAvailability").
        Set("total_count", sq.Expr("total_count - 1")).
        Set("rented_count", sq.Expr("rented_count + 1")).
        Where(sq.Eq{"cassette_id": newReservate.CassetteId}).
        Where(sq.Eq{"store_id": newReservate.StoreId}).
        RunWith(tx). // Используем tx для выполнения внутри транзакции
        PlaceholderFormat(sq.Dollar)

    // Логирование запроса для отладки
    q, args, err := updateQuery.ToSql()
    if err != nil {
        return 0, fmt.Errorf("failed to convert update query to SQL: %w", err)
    }
    log.Printf("Update SQL query: %s %v", q, args)

    _, err = updateQuery.Exec()
    if err != nil {
        return 0, fmt.Errorf("failed to update CassetteAvailability: %w", err)
    }

    // Подтверждение транзакции
    err = tx.Commit()
    if err != nil {
        return 0, err
    }

    return id, nil
}

// func (s *ReservationStorage) CreateReservation(newReservate core.Reservation) (int, error) {
//     tx, err := s.db.Begin()
//     if err != nil {
//         return 0, err
//     }
//     defer tx.Rollback()

//     query := sq.Insert("Reservations").
//         Columns("client_id", "cassette_id", "store_id", "status").
//         Values(newReservate.ClientId, newReservate.CassetteId, newReservate.StoreId, "reserve").
//         Suffix("RETURNING \"id\"").
//         RunWith(s.db).
//         PlaceholderFormat(sq.Dollar)

//     var id int
//     err = query.QueryRow().Scan(&id)
//     if err != nil {
//         return 0, err
//     }

//     updateQuery := sq.Update("CassetteAvailability").
//         Set("total_count", sq.Expr("total_count - 1")).
//         Set("rented_count", sq.Expr("rented_count + 1")).
//         Where(sq.Eq{"cassette_id": newReservate.CassetteId}).
//         Where(sq.Eq{"store_id": newReservate.StoreId})

//     _, err = updateQuery.RunWith(s.db).Exec()
//     if err != nil {
//         return 0, fmt.Errorf("failed to update CassetteAvailability: %w", err)
//     }

//     err = tx.Commit()
//     if err != nil {
//         return 0, err
//     }

//     return id, nil
// }


    // checkQuery := sq.Select("total_count").From("CassetteAvailability").
    //     Where(sq.Eq{"cassette_id": newReservate.CassetteId}).
    //     Where(sq.Eq{"store_id": newReservate.StoreId})

    // q, args, err := checkQuery.ToSql()
    // log.Printf("SQL query: %s %v", q, args) // Исправлено для правильного вывода аргументов

    // var totalCount int
    // err = checkQuery.RunWith(s.db).QueryRow().Scan(&totalCount)
    // if err != nil {
    //     if err == sql.ErrNoRows {
    //         return 0, fmt.Errorf("CassetteAvailability not found")
    //     }
    //     return 0, fmt.Errorf("failed to check CassetteAvailability: %w", err)
    // }

    // if totalCount == 0 {
    //     return 0, fmt.Errorf("total_count cannot be zero")
    // }


// func (s *ReservationStorage) CreateReservation(newReservate core.Reservation) (int, error) {
// 	tx, err := s.db.Begin()
// 	if err != nil {
// 		return 0, err
// 	}
// 	defer tx.Rollback()

// 	query := sq.Insert("Reservations").
// 		Columns("client_id", "cassette_id", "store_id", "status").
// 		Values(newReservate.ClientId, newReservate.CassetteId, newReservate.StoreId, "reserve").
// 		Suffix("RETURNING \"id\"").
// 		RunWith(s.db).
// 		PlaceholderFormat(sq.Dollar)

// 	var id int
// 	err = query.QueryRow().Scan(&id)
// 	if err != nil {
// 		return 0, err
// 	}

// 	checkQuery := sq.Select("total_count").From("CassetteAvailability").
// 		Where(sq.Eq{"cassette_id": newReservate.CassetteId, "store_id": newReservate.StoreId}).RunWith(s.db)

// 	var totalCount int
// 	err = checkQuery.QueryRow().Scan(&totalCount)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			// Запись не найдена, возвращаем ошибку
// 			return 0, fmt.Errorf("CassetteAvailability not found")
// 		}
// 		return 0, fmt.Errorf("failed to check CassetteAvailability: %w", err)
// 	}

// 	// Если total_count равно 0, возвращаем ошибку
// 	if totalCount == 0 {
// 		return 0, fmt.Errorf("total_count cannot be zero")
// 	}

// 	// Обновляем total_count и rented_count
// 	updateQuery := sq.Update("CassetteAvailability").
// 		Set("total_count", sq.Expr("total_count - 1")).
// 		Set("rented_count", sq.Expr("rented_count + 1")).
// 		Where(sq.Eq{"cassette_id": newReservate.CassetteId, "store_id": newReservate.StoreId}).
// 		RunWith(s.db).
// 		PlaceholderFormat(sq.Dollar)

// 	_, err = updateQuery.Exec()
// 	if err != nil {
// 		return 0, fmt.Errorf("failed to update CassetteAvailability: %w", err)
// 	}

// 	err = tx.Commit()
// 	if err != nil {
// 		return 0, err
// 	}

// 	return id, nil
// }
