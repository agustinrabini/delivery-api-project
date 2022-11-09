package location

import (
	"context"
	"database/sql"
	"delivery-api-project/domain"
)

const (
	getReceiverLocationyByOrder  = "SELECT * FROM delivery WHERE id_order = ? and type = receiver"
	getRemittentLocationyByOrder = "SELECT * FROM delivery WHERE id_order = ? and type = remitent"
)

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

//Returns the recevier location and the remitter location of an order.
func (r *repository) GetReceiverAndRemittentLocation(ctx context.Context, id int) (receiver domain.Location, remittent domain.Location, err error) {

	receiverResult, err := r.db.Query(getReceiverLocationyByOrder, id)
	if err != nil {
		return domain.Location{}, domain.Location{}, err
	}

	remittentResult, err := r.db.Query(getRemittentLocationyByOrder, id)
	if err != nil {
		return domain.Location{}, domain.Location{}, err
	}

	for receiverResult.Next() {
		err := receiverResult.Scan(&receiver.Id, &receiver.IdOrder, &receiver.Type, &receiver.Province, &receiver.City, &receiver.Commune, &receiver.FullAddress, &receiver.Lat, &receiver.Lng)
		if err != nil {
			return domain.Location{}, domain.Location{}, err
		}
	}

	for remittentResult.Next() {
		err := remittentResult.Scan(&remittent.Id, &remittent.IdOrder, &remittent.Type, &remittent.Province, &remittent.City, &remittent.Commune, &remittent.FullAddress, &remittent.Lat, &remittent.Lng)
		if err != nil {
			return domain.Location{}, domain.Location{}, err
		}
	}

	return receiver, remittent, nil
}
