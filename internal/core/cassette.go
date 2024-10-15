package core

type Cassette struct {
	Id         int `json:"id,omitempty" db:"id"`
	Name       string `json:"name" db:"name"`
	Genre      string `json:"genre,omitempty" db:"genre"`
	Year       int    `json:"year,omitempty" db:"year_of_release"`
}

type CassetteAvailability struct {
	CassetteId   int `json:"cassetteId" db:"cassette_id"`
	StoreId      int `json:"storeId" db:"store_id"`
	TotalCount   int    `json:"totalCount" db:"total_count"`
	RentedCount  int    `json:"rentedCount,omitempty" db:"rented_count"`
}