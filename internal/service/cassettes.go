package service

import (
	"github.com/fishmanDK/miet_project/internal/core"
	"github.com/fishmanDK/miet_project/internal/storage"
)

type CassettesService struct{
	storage *storage.Storage
}

func newCassettesService(storage *storage.Storage) *CassettesService{
	return &CassettesService{
		storage: storage,
	}
}

func (s *CassettesService) GetCassette(cassetteID int) (core.Cassette, error){
	res, err := s.storage.Cassettes.GetCassette(cassetteID)
	return res, err
}

func (s *CassettesService) GetCassettes() ([]core.Cassette, error){
	res, err := s.storage.Cassettes.GetCassettes()
	return res, err
}

func (s *CassettesService) GetCassettesByStoreID(id int) ([]core.Cassette, error){
	res, err := s.storage.Cassettes.GetCassettesByStoreID(id)
	return res, err
}

func (s *CassettesService) GetCassetteDetails(cassetteID, userID int) (core.CassetteAvailability, error){
	res, err := s.storage.Cassettes.GetCassetteDetails(cassetteID, userID)
	return res, err
}


func (s *CassettesService) CreateCassette(input core.CreateCassetteReq) (int, error){
	id, err := s.storage.Cassettes.CreateCassette(input)
	return id, err
}


func (s *CassettesService) CreateCassetteAvailability(newData core.CassetteAvailability) error{
	return s.storage.Cassettes.CreateCassetteAvailability(newData)
}

func (s *CassettesService) DeleteCasseteByID(cassetteID int) error{
	return s.storage.Cassettes.DeleteCasseteByID(cassetteID)
}

func (s *CassettesService) SaveCassetteChanges(changes core.ChangeCassette) error{
	return s.storage.Cassettes.SaveCassetteChanges(changes)
}