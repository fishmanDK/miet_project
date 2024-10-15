package storage

import (
	"fmt"

	"github.com/fishmanDK/miet_project/internal/core"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Clients interface{
	CreateClient(newClient core.Client) (int, error)
}
type Cassettes interface{
	GetCassettes() ([]core.Cassette, error)
	CreateCassette(input core.Cassette) (int, error)
	CreateCassetteAvailability(newData core.CassetteAvailability) error
}
type Store interface{
	CreateStore(newStore core.Store) (int, error)
}

type Reservation interface{
	CreateReservation(newReservate core.Reservation) (int, error)
}


type Storage struct{
	Clients Clients
	Cassettes Cassettes
	Store Store
	Reservation Reservation
}

type Config struct{
	User	string 	`yaml:"pg_user"`
	Database	string 	`yaml:"pg_database"`
	Host	string 	`yaml:"pg_host"`
	Port	string 	`yaml:"pg_port"`
	Sslmode	string 	`yaml:"pg_sslmode"`
	Password	string 	`yaml:"pg_password"`
}

func NewStorage(db *sqlx.DB) *Storage {
	return &Storage{
		Cassettes: newCassettesStorage(db),
		Clients: newClientStorage(db),
		Store: newStoreStorage(db),
		Reservation: newReservationStorage(db),
	}
}

func (c *Config) ToString() string{
	s := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", c.Host, c.Port, c.User, c.Database, c.Password, c.Sslmode)
	return s
}