package core

type Order struct {
	ID           string `json:"id,omitempty" db:"id"`
	UserId       int    `json:"userId" db:"user_id"`
	CassetteId   int    `json:"cassetteId" db:"cassette_id"`
	StoreID      int    `json:"storeId" db:"store_id"`
	OrderDate    string `json:"orderDate,omitempty" db:"order_date"`
	NameCassette string `json:"name" db:"name"`
	StoreAddress string `json:"address" db:"address"`
}

type OrderExecution struct {
	OrderId    string `json:"orderId" db:"order_id"`
	ReturnDate string `json:"returnDate,omitempty" db:"return_date"`
	IssueDate  string `json:"issueDate,omitempty" db:"issue_date"`
}

type DeleteOrder struct {
	UserId     string `json:"userId" db:"user_id"`
	CassetteId string `json:"cassetteId" db:"cassette_id"`
	StoreID    string `json:"storeId" db:"store_id"`
}
