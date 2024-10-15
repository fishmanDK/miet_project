package core

type Client struct {
	Id               int `json:"id,omitempty" db:"id"`
	Name             string `json:"name" db:"name"`
	Email            string `json:"email,omitempty" db:"email"`
	Password         string `json:"password" db:"password"`
	RegistrationDate string `json:"registrationDate,omitempty" db:"registration_date"`
}
