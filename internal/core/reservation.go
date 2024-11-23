package core

type Reservation struct {
	UserId          int    `json:"user_id" db:"user_id"`
	CassetteId      int    `json:"cassette_id" db:"cassette_id"`
	ReservationDate string `json:"reservation_date: db:"reservation_date"`
	Cassette
}

type DeleteReservation struct {
	CassetteId int `json:"cassette_id" db:"cassette_id"`
}

