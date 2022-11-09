package order

import (
	"context"
	request "delivery-api-project/controllers/web"
	response "delivery-api-project/controllers/web"
	"delivery-api-project/internal/delivery"
	"delivery-api-project/internal/location"
	"delivery-api-project/internal/packages"
)

/*
	Creación de orden
	Consulta de orden
	Actualización de estatus de la orden
	Cancelación de la orden con y sin reembolso
*/
type Service interface {
	Get(ctx context.Context, id int) (response.Order, error)
	Create(ctx context.Context, order request.Order) error
	UpdateStatus(ctx context.Context, idOrder int) error
	CancelOrder(ctx context.Context, idOrder int) (bool, error)
}

type service struct {
	repository   Repository
	packagesRepo packages.Repository
	delvieryRepo delivery.Repository
	locationRepo location.Repository
}

func NewService(repository Repository, packagesRepo packages.Repository, delvieryRepo delivery.Repository, locationRepo location.Repository) Service {
	return &service{
		repository:   repository,
		packagesRepo: packagesRepo,
		delvieryRepo: delvieryRepo,
		locationRepo: location.Repository,
	}
}

func (s *service) Get(ctx context.Context, id int) (response.Order, error) {

	order, err := s.repository.Get(ctx, id)
	if err != nil {
		return response.Order{}, err
	}

	packages, err := s.packagesRepo.GetPackagesByOrder(ctx, id)
	if err != nil {
		return response.Order{}, err
	}

	delivery, err := s.delvieryRepo.GetDeliveryByOrder(ctx, id)
	if err != nil {
		return response.Order{}, err
	}

	receiverLocation, remittentLocation, err := s.locationRepo.GetReceiverAndRemittentLocation(ctx, id)
	if err != nil {
		return response.Order{}, err
	}

	responseOrder, err := buildResponseOrder(order, packages, delivery, receiverLocation, remittentLocation)
	if err != nil {
		return response.Order{}, err
	}

	return responseOrder, nil
}
