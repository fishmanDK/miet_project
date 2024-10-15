package service

import (
	"github.com/fishmanDK/miet_project/internal/core"
	"github.com/fishmanDK/miet_project/internal/storage"
)

type ClientService struct{
	storage *storage.Storage
}

func newClientService(storage *storage.Storage) *ClientService{
	return &ClientService{
		storage: storage,
	}
}

func (s *ClientService) CreateClient(newClient core.Client) (int, error){
	id, err := s.storage.Clients.CreateClient(newClient)
	return id, err
}