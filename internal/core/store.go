package core

type Store struct {
	Id      int `json:"id,omitempty" db:"id"`
	Address string `json:"address,omitempty" db:"address"`
}