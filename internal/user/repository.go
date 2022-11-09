package user

import (
	"database/sql"
	"errors"

	"github.com/agustinrabini/delivery-api-project/domain"
)

type Repository interface {
	GetUser(id int) (usr domain.User, err error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetUser(id int) (usr domain.User, err error) {

	query := "select user_id, lat, lng from user where user_id = ?"

	result, err := r.db.Query(query, id)
	if err != nil {
		return domain.User{}, errors.New(err.Error())
	}

	for result.Next() {
		err := result.Scan(&usr.ID, &usr.Lat, &usr.Lng)
		if err != nil {
			return domain.User{}, errors.New(err.Error())
		}
	}

	return usr, nil
}
