package location

import (
	"context"
	"database/sql"
	"delivery-api-project/domain"
)

const (
	getReceiverLocationyByOrder  = "SELECT * FROM location WHERE id_order = ? and type = receiver"
	getRemittentLocationyByOrder = "SELECT * FROM location WHERE id_order = ? and type = remitent"
	createLocation               = "INSERT INTO location (id, id_order, type, province, city, commune, full_address, lat, lng) VALUES (?,?,?,?,?,?,?,?,?)"
)

type repository struct {
	db *sql.DB
}

type Repository interface {
	Create(ctx context.Context, loc domain.Location, typeLoc string) (*int, error)
	GetReceiverAndRemittentLocation(ctx context.Context, id int) (receiver domain.Location, remittent domain.Location, err error)
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Create(ctx context.Context, loc domain.Location, typeLoc string) (*int, error) {

	stmt, err := r.db.Prepare(createLocation)
	if err != nil {
		return nil, err
	}

	res, err := stmt.Exec(nil, typeLoc, loc.Province, loc.City, loc.Commune, loc.FullAddress, loc.Lat, loc.Lng)
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
		err := receiverResult.Scan(&receiver.Id, &receiver.Type, &receiver.Province, &receiver.City, &receiver.Commune, &receiver.FullAddress, &receiver.Lat, &receiver.Lng)
		if err != nil {
			return domain.Location{}, domain.Location{}, err
		}
	}

	for remittentResult.Next() {
		err := remittentResult.Scan(&remittent.Id, &remittent.Type, &remittent.Province, &remittent.City, &remittent.Commune, &remittent.FullAddress, &remittent.Lat, &remittent.Lng)
		if err != nil {
			return domain.Location{}, domain.Location{}, err
		}
	}

	return receiver, remittent, nil
}
