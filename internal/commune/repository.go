package commune

import (
	"database/sql"
	"errors"
	"log"

	"github.com/agustinrabini/delivery-api-project/domain"
)

type Repository interface {
	GetAll() (cms domain.Communes, err error)
	GetOne(number int) (c domain.Commune, err error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll() (cms domain.Communes, err error) {

	var cm domain.Commune

	query := "select * from communes"

	result, err := r.db.Query(query)
	if err != nil {
		return domain.Communes{}, errors.New(err.Error())
	}

	for result.Next() {
		err := result.Scan(&cm.Number, &cm.Lat, &cm.Lng)
		if err != nil {
			return domain.Communes{}, errors.New(err.Error())
		}

		cms = append(cms, cm)
	}

	return cms, nil
}

func (r *repository) GetOne(number int) (c domain.Commune, err error) {

	query := "select number, lat, lng from communes where number = ?"

	result, err := r.db.Query(query, number)
	if err != nil {
		log.Fatal("unnable to run the query: ", err)
	}

	for result.Next() {
		err := result.Scan(&c.Number, &c.Lat, &c.Lng)
		if err != nil {
			return domain.Commune{}, errors.New(err.Error() + "aca")
		}
	}

	return c, nil
}
