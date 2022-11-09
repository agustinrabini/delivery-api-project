package packages

import (
	"context"
	"database/sql"
	"delivery-api-project/domain"
)

const getPackagesByOrder = "SELECT * FROM package WHERE id_order = ?"

type repository struct {
	db *sql.DB
}

type Repository interface {
	GetPackagesByOrder(ctx context.Context, id int) ([]domain.Package, error)
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetPackagesByOrder(ctx context.Context, id int) ([]domain.Package, error) {

	ps := []domain.Package{}
	p := domain.Package{}

	rows, err := r.db.Query(getPackagesByOrder)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		_ = rows.Scan(&p.Id, &p.IdOrder, &p.Weight, &p.Size, &p.QuantityItems)
		ps = append(ps, p)
	}

	return ps, nil
}
