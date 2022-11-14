package delivery

import (
	"context"
	"database/sql"
	"delivery-api-project/domain"
	"errors"
)

const (
	createDelivery     = "INSERT INTO delivery (id, id_origin_location, id_destiny_location, pick_up_date, delivery_date) VALUES (?,?,?,?,?)"
	getDeliveryByOrder = "SELECT * FROM delivery WHERE id = ?"
)

type repository struct {
	db *sql.DB
}

type Repository interface {
	Create(ctx context.Context, loc domain.Delivery) (*int, error)
	GetDeliveryByOrder(ctx context.Context, id int) (domain.Delivery, error)
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Create(ctx context.Context, loc domain.Delivery) (*int, error) {

	stmt, err := r.db.Prepare(createDelivery)
	if err != nil {
		return nil, err
	}

	res, err := stmt.Exec(nil, loc.IdOriginLocation, loc.IdDestinyLocation, loc.PickUpDate, loc.DeliveryDate)
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

func (r *repository) GetDeliveryByOrder(ctx context.Context, id int) (domain.Delivery, error) {

	d := domain.Delivery{}

	result, err := r.db.Query(getDeliveryByOrder, id)
	if err != nil {
		return domain.Delivery{}, err
	}

	for result.Next() {
		err := result.Scan(&d.Id, &d.IdOriginLocation, &d.IdDestinyLocation, &d.PickUpDate, &d.DeliveryDate)
		if err != nil {
			return domain.Delivery{}, errors.New(err.Error())
		}
	}

	return d, nil
}
