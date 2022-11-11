package packages

import (
	"context"
	"database/sql"
	"delivery-api-project/domain"
)

const (
	createLocation     = "INSERT INTO location (id, id_order, weight, size, quantity_items) VALUES (?,?,?,?,?)"
	getPackagesByOrder = "SELECT * FROM package WHERE id_order = ?"
)

type repository struct {
	db *sql.DB
}

/* Id            int     `json:"id"`
IdOrder       int     `json:"id_order" validate:"required"`
Weight        float32 `json:"weight"validate:"required,number,min=1,max=25"`
Size          string  `json:"size" validate:"required"`
QuantityItems int     `json:"quantity_items" validate:"required"`
*/

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
	stmt, err := r.db.Prepare(createLocation)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(nil, pk.IdOrder, pk.Weight, pk.Size, pk.QuantityItems)
	if err != nil {
		return err
	}

	return nil
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
