package storage

import (
	"fmt"

	"github.com/fishmanDK/miet_project/internal/core"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Orders interface {
	GetUserOrders(userID int) ([]core.Order, error)
	CreateOrder(newOrder core.Order) (int, error)
	DeleteOrder(userID, cassetteID int) error
	GetOrdersForAdmin(cassetteID, storeID int) ([]core.OrdersForAdminResponse, error)
}

type Auth interface {
	Authentication(user core.Client) (core.AuthResult, error)
	CreateSession(userId int, session core.Session) error
	CreateUser(newUser core.Client) (int, error)
}
type Cassettes interface {
	GetCassette(cassetteID int) (core.Cassette, error)
	GetCassettes() ([]core.Cassette, error)
	GetCassettesByStoreID(id int) ([]core.Cassette, error)
	GetCassetteDetails(cassetteID, userID int) (core.CassetteAvailability, error)
	CreateCassette(input core.CreateCassetteReq) (int, error)
	CreateCassetteAvailability(newData core.CassetteAvailability) error
	DeleteCasseteByID(cassetteID int) error
	SaveCassetteChanges(changes core.ChangeCassette) error
}
type Store interface {
	GetStores() ([]core.Store, error)
	CreateStore(newStore core.Store) (int, error)
}

type Reservation interface {
	CreateReservation(newReservate core.Reservation) error
	DeleteReservation(userID, cassetteID int) error
	GetUserReservations(userID int) ([]core.Reservation, error)
}

type Storage struct {
	Auth        Auth
	Cassettes   Cassettes
	Store       Store
	Reservation Reservation
	Orders      Orders
}

type Config struct {
	User     string `yaml:"pg_user"`
	Database string `yaml:"pg_database"`
	Host     string `yaml:"pg_host"`
	Port     string `yaml:"pg_port"`
	Sslmode  string `yaml:"pg_sslmode"`
	Password string `yaml:"pg_password"`
}

func NewStorage(db *sqlx.DB) *Storage {
	return &Storage{
		Cassettes:   newCassettesStorage(db),
		Auth:        newAuthStorage(db),
		Store:       newStoreStorage(db),
		Reservation: newReservationStorage(db),
		Orders:      newOrdersStorage(db),
	}
}

func (c *Config) ToString() string {
	s := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", c.Host, c.Port, c.User, c.Database, c.Password, c.Sslmode)
	return s
}
