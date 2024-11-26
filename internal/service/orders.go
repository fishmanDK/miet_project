package service

import (
	"github.com/fishmanDK/miet_project/internal/core"
	"github.com/fishmanDK/miet_project/internal/storage"
)

type OrderService struct {
	storage *storage.Storage
}

func NewOrderService(storage *storage.Storage) *OrderService {
	return &OrderService{
		storage: storage,
	}
}

func (s *OrderService) GetUserOrders(userID int) ([]core.Order, error) {
	orders, err := s.storage.Orders.GetUserOrders(userID)
	return orders, err
}

func (s *OrderService) CreateOrder(newOrder core.Order) (int, error) {
	id, err := s.storage.Orders.CreateOrder(newOrder)
	return id, err
}

func (s *OrderService) DeleteOrder(userID, cassetteID int) error {
	return s.storage.Orders.DeleteOrder(userID, cassetteID)
}

func (s *OrderService) GetOrdersForAdmin(cassetteID, storeID int) ([]core.OrdersForAdminResponse, error) {
	res, err := s.storage.Orders.GetOrdersForAdmin(cassetteID, storeID)
	return res, err
}
