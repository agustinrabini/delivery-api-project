package shop

import (
	"database/sql"
	"errors"

	"github.com/agustinrabini/delivery-api-project/domain"
)

type Repository interface {
	GetAll() (s domain.Shops, err error)
	GetOne(id int) (sh domain.Shop, err error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll() (s domain.Shops, err error) {

	var sh domain.Shop

	query := "select * from shops"

	result, err := r.db.Query(query)
	if err != nil {
		return domain.Shops{}, errors.New(err.Error())
	}

	for result.Next() {
		err := result.Scan(&sh.Shop_id, &sh.Commune, &sh.Lat, &sh.Lng)
		if err != nil {
			return domain.Shops{}, errors.New(err.Error())
		}

		s = append(s, sh)

	}

	return s, nil
}

func (r *repository) GetOne(id int) (sh domain.Shop, err error) {

	query := "select * from shops where shop_id = ?"

	result, err := r.db.Query(query, id)
	if err != nil {
		return domain.Shop{}, errors.New(err.Error())
	}

	for result.Next() {
		err := result.Scan(&sh.Shop_id, &sh.Commune, &sh.Lat, &sh.Lng)
		if err != nil {
			return domain.Shop{}, errors.New(err.Error())
		}
	}

	return sh, nil
}
