package core

type Cassette struct {
	Id    int    `json:"id,omitempty" db:"id"`
	Name  string `json:"name" db:"name"`
	Genre string `json:"genre,omitempty" db:"genre"`
	Year  string    `json:"year,omitempty" db:"year_of_release"`
	TotalCount   int  `json:"totalCount" db:"total_count"`
	RentedCount  int  `json:"rentedCount,omitempty" db:"rented_count"`
}

type CassetteAvailability struct {
	UserID       int  `json:"user_id,omitempty" db:"user_id"`
	IsOrdered    bool `json:"isOrdered" db:"isOrdered"`
	IsReservated bool `json:"isReservated" db:"isReservated"`
	CassetteId   int  `json:"cassetteId,omitempty" db:"cassette_id"`
	StoreId      int  `json:"storeId,omitempty" db:"store_id"`
	TotalCount   int  `json:"totalCount" db:"total_count"`
	RentedCount  int  `json:"rentedCount,omitempty" db:"rented_count"`
}

type CreateCassetteReq struct {
	Name       string `json:"name" db:"name"`
	Genre      string `json:"genre,omitempty" db:"genre"`
	Year       int    `json:"year,omitempty" db:"year_of_release"`
	StoreId    int    `json:"storeId,omitempty" db:"store_id"`
	TotalCount int    `json:"totalCount" db:"total_count"`
}
