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


func (s *ReservationService) CreateReservation(newReservate core.Reservation) error{
	return s.storage.Reservation.CreateReservation(newReservate)
}

func (s *ReservationService) DeleteReservation(userID, cassetteID int) error{
	return s.storage.Reservation.DeleteReservation(userID, cassetteID)
}


func (s *ReservationService) GetUserReservations(userID int) ([]core.Reservation, error){
	res, err := s.storage.Reservation.GetUserReservations(userID)
	return res, err
}