package service

import (
	"github.com/fishmanDK/miet_project/internal/core"
	"github.com/fishmanDK/miet_project/internal/storage"
)

type Service struct {
	Clients     Clients
	Cassettes  Cassettes
	Store    Store
	Reservation Reservation
}

type Clients interface{
	CreateClient(newClient core.Client) (int, error)
}
type Cassettes interface{
	GetCassettes() ([]core.Cassette, error)
	CreateCassette(newCassette core.Cassette) (int, error)
	CreateCassetteAvailability(newData core.CassetteAvailability) error
}
type Store interface{
	CreateStore(newStore core.Store) (int, error)
}
type Reservation interface{
	CreateReservation(newReservate core.Reservation) (int, error)
}

func NewSerivce(storage *storage.Storage) *Service {
	return &Service{
		Cassettes: newCassettesService(storage),
		Clients: newClientService(storage),
		Store: newStoreService(storage),
		Reservation: newReservationService(storage),
	}
}