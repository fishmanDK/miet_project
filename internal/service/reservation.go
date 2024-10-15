package service

import (
	"github.com/fishmanDK/miet_project/internal/core"
	"github.com/fishmanDK/miet_project/internal/storage"
)

type ReservationService struct{
	storage *storage.Storage
}

func newReservationService(storage *storage.Storage) *ReservationService{
	return &ReservationService{
		storage: storage,
	}
}


func (s *ReservationService) CreateReservation(newReservate core.Reservation) (int, error){
	id, err := s.storage.Reservation.CreateReservation(newReservate)
	return id, err
}