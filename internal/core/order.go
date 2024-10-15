package core

type Order struct {
	Id        string `json:"id,omitempty" db:"id"`
	ClientId  string `json:"clientId" db:"client_id"`
	CassetteId string `json:"cassetteId" db:"cassette_id"`
	StoreId   string `json:"storeId" db:"store_id"`
	OrderDate string `json:"orderDate,omitempty" db:"order_date"`
	Status    string `json:"status,omitempty" db:"status"`
}

type OrderExecution struct {
	OrderId     string `json:"orderId" db:"order_id"` 
	ReturnDate  string `json:"returnDate,omitempty" db:"return_date"` 
	IssueDate   string `json:"issueDate,omitempty" db:"issue_date"` 
}