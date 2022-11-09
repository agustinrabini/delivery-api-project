package communesshop

import (
	"database/sql"
	"errors"

	"github.com/agustinrabini/poc-proximity-objects/domain"
)

type Repository interface {
	GetShopByCommune(communeNumber int) (cshs domain.CommunesShops, err error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetShopByCommune(cn int) (cshs domain.CommunesShops, err error) {

	var c domain.CommunesShop

	query := "select * from communesshop where commune_number = ?"

	result, err := r.db.Query(query, cn)
	if err != nil {
		return domain.CommunesShops{}, errors.New(err.Error())
	}

	for result.Next() {
		err := result.Scan(&c.Commune_shop_id, &c.Shop_id, &c.CommuneNumber)
		if err != nil {
			return domain.CommunesShops{}, errors.New(err.Error())
		}
		cshs = append(cshs, c)

	}

	return cshs, nil
}
