package cart

import (
	"database/sql"

	"github.com/Eslam-Amin/ecommerce/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) CreateOrder(order types.Order) (int, error) {
	res, err := s.db.Exec(`
	INSERT INTO order 
	(userId, total, status, address)
	VALUES 
(?,?,?,?)
	`, order.UserId, order.Total, order.Status, order.Address)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, nil
	}
	return int(id), nil
}

func (s *Store) CreateOrderItem(orderItem types.OrderItem) error {
	_, err := s.db.Exec(`
	INSERT INTO order_items 
	(orderId, productId, quantity, price)
	VALUES
	(?,?,?,?)`, orderItem.OrderId, orderItem.ProductId, orderItem.Quantity, orderItem.Price)

	return err
}
