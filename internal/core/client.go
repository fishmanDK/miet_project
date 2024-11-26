package core

type Client struct {
	Id               int    `json:"id,omitempty" db:"id"`
	Email            string `json:"email" db:"email"`
	Password         string `json:"password" db:"password"`
	Role             string `json:"role" db:"role"`
	RegistrationDate string `json:"registrationDate,omitempty" db:"registration_date" goqu:"skipinsert"`
}

type AuthResult struct {
	Id    int    `db:"id"`
	Email string `db:"email"`
	Role  string `db:"role"`
}
