package order

import (
	"context"
	"database/sql"
	"delivery-api-project/domain"
	"fmt"
)

const (
	getOrder     = "SELECT * FROM orders WHERE id = ?"
	create       = "INSERT INTO orders (id, id_delivery, id_receiver, id_remitter, status, creation_date) VALUES (?,?,?,?,?,?)"
	updateStauts = "UPDATE orders SET status = ? WHERE id = ? "
)

type repository struct {
	db *sql.DB
}

type Repository interface {
	Get(ctx context.Context, id int) (domain.Order, error)
	Create(ctx context.Context, order domain.Order) (*int, error)
	UpdateStatus(ctx context.Context, idOrder int, newStatus string) error
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Get(ctx context.Context, id int) (domain.Order, error) {

	order := domain.Order{}

	result, err := r.db.Query(getOrder, id)
	if err != nil {
		return domain.Order{}, err
	}

	for result.Next() {
		err := result.Scan(&order.Id, &order.IdDelivery, &order.ReceiverID, &order.RemitterID, &order.Status, &order.CreationDate)
		if err != nil {
			return domain.Order{}, err
		}

		if order.Id == nil && err == nil {
			return domain.Order{}, nil
		}
	}

	return order, nil
}

func (r *repository) Create(ctx context.Context, order domain.Order) (*int, error) {
	stmt, err := r.db.Prepare(create)
	if err != nil {
		return nil, err
	}

	res, err := stmt.Exec(nil, order.IdDelivery, order.ReceiverID, order.RemitterID, order.Status, order.CreationDate)
	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	finalId := int(id)

	return &finalId, nil
}

func (r *repository) UpdateStatus(ctx context.Context, idOrder int, newStatus string) error {
	stmt, err := r.db.Prepare(updateStauts)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(newStatus, idOrder)
	if err != nil {
		return err
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if affect < 1 {
		return fmt.Errorf("object not updated, check if it exists")
	}

	return nil
}
