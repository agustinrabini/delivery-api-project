package packages

import (
	"context"
	"database/sql"
	"delivery-api-project/domain"
)

const (
	createPkg          = "INSERT INTO package (id, id_order, weight, size, quantity_items) VALUES (?,?,?,?,?)"
	getPackagesByOrder = "SELECT * FROM package WHERE id_order = ?"
)

type repository struct {
	db *sql.DB
}

type Repository interface {
	Create(ctx context.Context, pk domain.Package) error
	GetPackagesByOrder(ctx context.Context, id int) ([]domain.Package, error)
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Create(ctx context.Context, pk domain.Package) error {
	stmt, err := r.db.Prepare(createPkg)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(nil, pk.IdOrder, pk.Weight, pk.Size, pk.QuantityItems)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) GetPackagesByOrder(ctx context.Context, idOrder int) ([]domain.Package, error) {

	ps := []domain.Package{}
	p := domain.Package{}

	rows, err := r.db.Query(getPackagesByOrder, idOrder)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&p.Id, &p.IdOrder, &p.Weight, &p.Size, &p.QuantityItems)
		if err != nil {
			return nil, err
		}
		ps = append(ps, p)
	}

	return ps, nil
}
