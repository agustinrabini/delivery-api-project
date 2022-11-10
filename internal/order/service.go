package order

import (
	"context"
	request "delivery-api-project/controllers/web/request"
	response "delivery-api-project/controllers/web/response"
	"delivery-api-project/domain"
	"delivery-api-project/internal/delivery"
	"delivery-api-project/internal/location"
	"delivery-api-project/internal/packages"
	"time"
)

/*
	Creación de orden
	Consulta de orden
	Actualización de estatus de la orden
	Cancelación de la orden con y sin reembolso
*/
type Service interface {
	Get(ctx context.Context, id int) (response.Order, error)
	Create(ctx context.Context, order request.Order) (*int, error)
	UpdateStatus(ctx context.Context, idOrder *int, status string) error
	CancelOrder(ctx context.Context, idOrder *int) (bool, error)
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
		locationRepo: locationRepo,
	}
}

func (s *service) Get(ctx context.Context, id int) (response.Order, error) {

	order, err := s.repository.Get(ctx, id)
	if err != nil {
		return response.Order{}, err
	}

	packages, err := s.packagesRepo.GetPackagesByOrder(ctx, *order.Id)
	if err != nil {
		return response.Order{}, err
	}

	delivery, err := s.delvieryRepo.GetDeliveryByOrder(ctx, *order.Id)
	if err != nil {
		return response.Order{}, err
	}

	receiverLocation, remittentLocation, err := s.locationRepo.GetReceiverAndRemittentLocation(ctx, *order.Id)
	if err != nil {
		return response.Order{}, err
	}

	responseOrder, err := buildResponseOrder(order, packages, delivery, receiverLocation, remittentLocation)
	if err != nil {
		return response.Order{}, err
	}

	return responseOrder, nil
}

func (s *service) Create(ctx context.Context, rrOrder request.Order) (*int, error) {

	receiverLocID, err := s.locationRepo.Create(ctx, rrOrder.Delivery.OriginLocation, "receiver")
	if err != nil {
		return nil, err
	}

	remittentLocID, err := s.locationRepo.Create(ctx, rrOrder.Delivery.OriginLocation, "receiver")
	if err != nil {
		return nil, err
	}

	delivery, err := buildDelivery(*remittentLocID, *receiverLocID)
	if err != nil {
		return nil, err
	}

	deliveryID, err := s.delvieryRepo.Create(ctx, delivery)
	if err != nil {
		return nil, err
	}

	orderId, err := s.repository.Create(ctx, domain.Order{
		Id:           nil,
		IdDelivery:   *deliveryID,
		ReceiverID:   rrOrder.ReceiverID,
		RemitterID:   rrOrder.RemitterID,
		Status:       "creado",
		CreationDate: time.Now().Format("2006-1-2 15:4:5"),
	})

	return orderId, nil
}

func (s *service) UpdateStatus(ctx context.Context, idOrder *int, status string) error {

	err := s.repository.UpdateStatus(ctx, *idOrder, status)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) CancelOrder(ctx context.Context, idOrder *int) (bool, error) {

	order, err := s.repository.Get(ctx, *idOrder)
	if err != nil {
		return false, nil
	}

	dateFormat := "2006-1-2 15:4:5"

	actualDate := time.Now()

	orderDateFormated, err := time.Parse(dateFormat, order.CreationDate)
	if err != nil {
		return false, err
	}

	diff := actualDate.Sub(orderDateFormated)
	if diff.Minutes() < 2 {
		return false, nil
	}

	err = s.repository.UpdateStatus(ctx, *idOrder, "cancelado")
	if err != nil {
		return false, err
	}

	return true, nil
}
