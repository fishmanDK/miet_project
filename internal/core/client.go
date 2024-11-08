package core

type Client struct {
	Id               int    `json:"id,omitempty" db:"id"`
	Email            string `json:"email,omitempty" db:"email"`
	Password         string `json:"password" db:"password"`
	RegistrationDate string `json:"registrationDate,omitempty" db:"registration_date"`
}

type AuthResult struct {
	Id    int  `db:"id"`
	Email string `db:"email"`
	Role  string `db:"role"`
}
