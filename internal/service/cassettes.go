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

func (s *CassettesService) GetCassettes() ([]core.Cassette, error){
	res, err := s.storage.Cassettes.GetCassettes()
	return res, err
}

func (s *CassettesService) CreateCassette(input core.Cassette) (int, error){
	id, err := s.storage.Cassettes.CreateCassette(input)
	return id, err
}


func (s *CassettesService) CreateCassetteAvailability(newData core.CassetteAvailability) error{
	return s.storage.Cassettes.CreateCassetteAvailability(newData)
}
