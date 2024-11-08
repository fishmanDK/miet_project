package storage

import (
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

func (s *ReservationStorage) CreateReservation(newReservate core.Reservation) error {
	query := sq.Insert("reserve_pool").
		Columns("user_id", "cassette_id").
		Values(newReservate.UserId, newReservate.CassetteId).
		RunWith(s.db).
		PlaceholderFormat(sq.Dollar)

	_, err := query.Exec()
	if err != nil{
		return err
	}

	return nil
}

func (s *ReservationStorage) DeleteReservation(userID, cassetteID int) error{
	query := sq.Delete("reserve_pool").
		Where(sq.Eq{"user_id": userID, "cassette_id": cassetteID}).
		RunWith(s.db).
		PlaceholderFormat(sq.Dollar)

	_, err := query.Exec()
	if err != nil {
		return err
	}

	return nil
}

func (s *ReservationStorage) GetUserReservations(userID int) ([]core.Reservation, error){
	query, _, err := sq.Select("cassette_id", "name").
		From("reserve_pool").
		Join("cassettes ON reserve_pool.cassette_id = cassettes.id").
		Where(sq.Eq{"user_id": userID}).
		RunWith(s.db).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return []core.Reservation{}, err
	}

	rows, err := s.db.Query(query, userID)
	if err != nil {
		return []core.Reservation{}, err
	}
	defer rows.Close()

	var res []core.Reservation
	for rows.Next() {
		var newVal core.Reservation

		err := rows.Scan(&newVal.CassetteId, &newVal.Name)
		if err != nil {
			return []core.Reservation{}, err
		}
		res = append(res, newVal)
	}

	if err := rows.Err(); err != nil {
		return []core.Reservation{}, err
	}

    return res, nil
}

func (s *ReservationStorage) GetReservationsForAdmin(cassetteID, storeID int)  ([]core.ReservationsForAdminResponse, error) {
	query, _, err := sq.Select("reservation_date", "email").
		From("clients").
		Join("reservations ON clients.id = reservations.client_id").
		Where(sq.Eq{"cassette_id": cassetteID}).
		Where(sq.Eq{"store_id": storeID}).
		RunWith(s.db).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return []core.ReservationsForAdminResponse{}, err
	}

	rows, err := s.db.Query(query, cassetteID, storeID)
	if err != nil {
		return []core.ReservationsForAdminResponse{}, err
	}
	defer rows.Close()

	var res []core.ReservationsForAdminResponse
	for rows.Next() {
		var newVal core.ReservationsForAdminResponse

		err := rows.Scan(&newVal.ReservationDate, &newVal.Email)
		if err != nil {
			return []core.ReservationsForAdminResponse{}, err
		}
		res = append(res, newVal)
	}

	if err := rows.Err(); err != nil {
		return []core.ReservationsForAdminResponse{}, err
	}

    return res, nil
}
