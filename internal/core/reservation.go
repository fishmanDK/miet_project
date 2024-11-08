package core

type Reservation struct {
	UserId        int    `json:"user_id" db:"user_id"`
	CassetteId      int    `json:"cassette_id" db:"cassette_id"`
	Cassette
}

type DeleteReservation struct {
	CassetteId      int    `json:"cassette_id" db:"cassette_id"`
}

type ReservationsForAdminResponse struct{
	ReservationDate string `json:"reservation_date,omitempty" db:"reservation_date"`
	Email string `json:"email,omitempty" db:"email"`
}