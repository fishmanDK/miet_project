package core

type Reservation struct {
	Id              int    `json:"id,omitempty" db:"id"`
	ClientId        int `json:"client-id" db:"client_id"`
	CassetteId      int `json:"cassette-id" db:"cassette_id"`
	StoreId         int `json:"store-id" db:"store_id"`
	ReservationDate string `json:"reservation-date,omitempty" db:"reservation_date"`
	Receipt_date    string `json:"receipt-date,omitempty" db:"receipt_date"`
	Return_date     string `json:"return-date,omitempty" db:"return_date"`
	Status          string `json:"status,omitempty" db:"status"`
}
