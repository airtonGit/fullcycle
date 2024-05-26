package database

import (
	"database/sql"

	"github.com/airtongit/fullcycle/cleanarch/internal/entity"
)

type OrderRepository struct {
	Db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{Db: db}
}

func (r *OrderRepository) Save(order *entity.Order) error {
	stmt, err := r.Db.Prepare("INSERT INTO orders (id, price, tax, final_price) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(order.ID, order.Price, order.Tax, order.FinalPrice)
	if err != nil {
		return err
	}
	return nil
}

func (r *OrderRepository) List() ([]*entity.Order, error) {

	resultSet, err := r.Db.Query("SELECT id, price, tax, final_price FROM orders")
	if err != nil {
		return nil, err
	}
	defer resultSet.Close()
	var orders []*entity.Order
	for resultSet.Next() {
		var order entity.Order
		if err := resultSet.Scan(&order.ID, &order.Price, &order.Tax, &order.FinalPrice); err != nil {
			return nil, err
		}
		orders = append(orders, &order)
	}

	return orders, nil
}

func (r *OrderRepository) GetTotal() (int, error) {
	var total int
	err := r.Db.QueryRow("Select count(*) from orders").Scan(&total)
	if err != nil {
		return 0, err
	}
	return total, nil
}