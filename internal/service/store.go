package service

import (
	"github.com/fishmanDK/miet_project/internal/core"
	"github.com/fishmanDK/miet_project/internal/storage"
)

type StoreService struct{
	storage *storage.Storage
}

func newStoreService(storage *storage.Storage) *StoreService{
	return &StoreService{
		storage: storage,
	}
}

func (s *StoreService) CreateStore(newStore core.Store) (int, error) {
	id, err := s.storage.Store.CreateStore(newStore)
	return id, err
}