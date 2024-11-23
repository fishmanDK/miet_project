package storage

import (
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/fishmanDK/miet_project/internal/core"
	"github.com/jmoiron/sqlx"
)

type OrdersStorage struct {
	db *sqlx.DB
}

func newOrdersStorage(db *sqlx.DB) *OrdersStorage {
	return &OrdersStorage{
		db: db,
	}
}

func (s *OrdersStorage) GetUserOrders(userID int) ([]core.Order, error) {
	query, args, err := sq.Select("c.id", "o.order_date", "c.name", "s.address").
		From("orders o").
		Join("cassettes c ON o.cassette_id = c.id").
		Join("stores s ON o.store_id = s.id").
		Where(sq.Eq{"user_id": userID}).
		RunWith(s.db).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	fmt.Println(query, args)

	if err != nil {
		return nil, err
	}

	rows, err := s.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []core.Order
	for rows.Next() {
		var order core.Order
		err := rows.Scan(&order.ID, &order.OrderDate, &order.NameCassette, &order.StoreAddress)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return orders, nil
}

func (s *OrdersStorage) CreateOrder(newOrder core.Order) (int, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	var totalCount, rentedCount int
	selectQuery := sq.Select("total_count", "rented_count").
		From("cassetteAvailability").
		Where(sq.Eq{"cassette_id": newOrder.CassetteId, "store_id": newOrder.StoreID}).
		RunWith(tx).
		PlaceholderFormat(sq.Dollar)

	fmt.Println(newOrder)
	fmt.Println(selectQuery.ToSql())
	err = selectQuery.QueryRow().Scan(&totalCount, &rentedCount)
	if err != nil {
		fmt.Println(1)
		return 0, err
	}

	if totalCount <= 0 {
		return 0, fmt.Errorf("кассета недоступна для заказа, нет в наличии")
	}

	insertQuery := sq.Insert("orders").
		Columns("user_id", "store_id", "cassette_id").
		Values(newOrder.UserId, newOrder.StoreID, newOrder.CassetteId).
		Suffix("RETURNING \"id\"").
		RunWith(tx).
		PlaceholderFormat(sq.Dollar)

	var id int
	err = insertQuery.QueryRow().Scan(&id)
	if err != nil {
		fmt.Println(2)
		return 0, err
	}

	updateQuery := sq.Update("cassetteAvailability").
		Set("total_count", totalCount-1).
		Set("rented_count", rentedCount+1).
		Where(sq.Eq{"cassette_id": newOrder.CassetteId, "store_id": newOrder.StoreID}).
		RunWith(tx).
		PlaceholderFormat(sq.Dollar)

	_, err = updateQuery.Exec()
	if err != nil {
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *OrdersStorage) DeleteOrder(userID, cassetteID int) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	queryGetStoreID := sq.Select("store_id").
		From("orders").
		Where(sq.Eq{"cassette_id": cassetteID}).
		Where(sq.Eq{"user_id": userID}).
		RunWith(tx).
		PlaceholderFormat(sq.Dollar)

	var storeID int
	err = queryGetStoreID.QueryRow().Scan(&storeID)
	if err != nil || storeID == 0{
		return err
	}

	queryDeleteOrder := sq.Delete("orders").
		Where(sq.Eq{"cassette_id": cassetteID}).
		Where(sq.Eq{"user_id": userID}).
		RunWith(tx).
		PlaceholderFormat(sq.Dollar)

	_, err = queryDeleteOrder.Exec()
	if err != nil {
		return err
	}

	var totalCount, rentedCount int
	selectQuery := sq.Select("total_count", "rented_count").
		From("cassetteAvailability").
		Where(sq.Eq{"cassette_id": cassetteID, "store_id": storeID}).
		RunWith(tx).
		PlaceholderFormat(sq.Dollar)

	err = selectQuery.QueryRow().Scan(&totalCount, &rentedCount)
	if err != nil {
		return err
	}

	if totalCount <= 0 {
		return fmt.Errorf("кассета недоступна для заказа, нет в наличии")
	}

	updateQuery := sq.Update("cassetteAvailability").
		Set("total_count", totalCount+1).
		Set("rented_count", rentedCount-1).
		Where(sq.Eq{"cassette_id": cassetteID, "store_id": storeID}).
		RunWith(tx).
		PlaceholderFormat(sq.Dollar)

	_, err = updateQuery.Exec()
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil

}

func (s *OrdersStorage) GetOrdersForAdmin(cassetteID, storeID int) ([]core.OrdersForAdminResponse, error) {
	query, args, err := sq.Select("order_date", "email", "cassette_id", "user_id").
		From("users").
		Join("orders ON users.id = orders.user_id").
		Where(sq.Eq{"cassette_id": cassetteID}).
		Where(sq.Eq{"store_id": storeID}).
		RunWith(s.db).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	fmt.Println(query, args)
	if err != nil {
		return []core.OrdersForAdminResponse{}, err
	}

	rows, err := s.db.Query(query, cassetteID, storeID)
	if err != nil {
		return []core.OrdersForAdminResponse{}, err
	}
	defer rows.Close()

	var res []core.OrdersForAdminResponse
	for rows.Next() {
		var newVal core.OrdersForAdminResponse

		err := rows.Scan(&newVal.ReservationDate, &newVal.Email, &newVal.CassetteID, &newVal.UserID)
		if err != nil {
			return []core.OrdersForAdminResponse{}, err
		}
		res = append(res, newVal)
	}

	if err := rows.Err(); err != nil {
		return []core.OrdersForAdminResponse{}, err
	}

	return res, nil
}
