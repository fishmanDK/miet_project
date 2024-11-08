package service

import (
	"github.com/fishmanDK/miet_project/internal/core"
	"github.com/fishmanDK/miet_project/internal/storage"
)

type Service struct {
	Auth        Auth
	Cassettes   Cassettes
	Store       Store
	Reservation Reservation
	Orders      Orders
}

type Orders interface {
	GetUserOrders(userID int) ([]core.Order, error)
	CreateOrder(newOrder core.Order) (int, error)
	DeleteOrder(delOrder core.DeleteOrder) error
}

type Auth interface {
	Authentication(user core.Client) (core.Tokens, error)
	ParseToken(accessToken string) (*ParseDataUser, error)
	CreateUser(newUser core.Client) (int, error)
}
type Cassettes interface {
	GetCassette(cassetteID int) (core.Cassette, error)
	GetCassettes() ([]core.Cassette, error)
	GetCassettesByStoreID(id int) ([]core.Cassette, error)
	GetCassetteDetails(cassetteID, userID int) (core.CassetteAvailability, error)
	CreateCassette(newCassette core.CreateCassetteReq) (int, error)
	CreateCassetteAvailability(newData core.CassetteAvailability) error
	DeleteCasseteByID(cassetteID int) error
}
type Store interface {
	GetStores() ([]core.Store, error)
	CreateStore(newStore core.Store) (int, error)
}
type Reservation interface {
	CreateReservation(newReservate core.Reservation) error
	DeleteReservation(userID, cassetteID int) error
	GetReservationsForAdmin(cassetteID, storeID int) ([]core.ReservationsForAdminResponse, error)
	GetUserReservations(userID int) ([]core.Reservation, error)
}

func NewSerivce(storage *storage.Storage) *Service {
	return &Service{
		Cassettes:   newCassettesService(storage),
		Auth:        newAuthService(storage),
		Store:       newStoreService(storage),
		Reservation: newReservationService(storage),
		Orders:      newOrderService(storage),
	}
}
