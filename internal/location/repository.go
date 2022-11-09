package location

import (
	"context"
	"database/sql"
	"delivery-api-project/domain"
	"errors"
)

const getDeliveryByOrder = "SELECT * FROM delivery WHERE id_order = ?"

type repository struct {
	db *sql.DB
}

type Repository interface {
	GetReceiverAndRemittentLocation(ctx context.Context, id int) (receiver domain.Location, remittent domain.Location, err error)
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetReceiverAndRemittentLocation(ctx context.Context, id int) (receiver domain.Location, remittent domain.Location, err error) {

	d := domain.Delivery{}

	result, err := r.db.Query(getDeliveryByOrder, id)
	if err != nil {
		return domain.Delivery{}, err
	}

	for result.Next() {
		err := result.Scan(&d.Id, &d.IdOrder, &d.IdOriginLocation, &d.IdDestinyLocation, &d.PickUpDate, &d.DeliveryDate)
		if err != nil {
			return domain.Delivery{}, errors.New(err.Error())
		}
	}

	return d, nil
}
