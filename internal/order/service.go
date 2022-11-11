package order

import (
	"context"
	request "delivery-api-project/controllers/web/request"
	response "delivery-api-project/controllers/web/response"
	"delivery-api-project/domain"
	"delivery-api-project/internal/delivery"
	"delivery-api-project/internal/location"
	"delivery-api-project/internal/packages"
	"fmt"
	"time"
)

const (
	status1 = "creado"
	status2 = "recolectado"
	status3 = "en_estacion"
	status4 = "en_ruta"
	status5 = "entregado"
	status6 = "cancelado"
)

type Service interface {
	Get(ctx context.Context, id int) (response.Order, error)
	Create(ctx context.Context, order request.Order) (*int, error)
	UpdateStatus(ctx context.Context, idOrder *int, status string) error
}

type service struct {
	repository   Repository
	packagesRepo packages.Repository
	deliveryRepo delivery.Repository
	locationRepo location.Repository
}

func NewService(repository Repository, packagesRepo packages.Repository, deliveryRepo delivery.Repository, locationRepo location.Repository) Service {
	return &service{
		repository:   repository,
		packagesRepo: packagesRepo,
		deliveryRepo: deliveryRepo,
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

	delivery, err := s.deliveryRepo.GetDeliveryByOrder(ctx, *order.Id)
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

	deliveryID, err := s.deliveryRepo.Create(ctx, delivery)
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
	if err != nil {
		return nil, err
	}

	err = buildPackages(ctx, *s, rrOrder.Packages, orderId)
	if err != nil {
		return nil, err
	}

	return orderId, nil
}

func (s *service) UpdateStatus(ctx context.Context, idOrder *int, status string) error {

	if status != status1 || status != status2 || status != status3 || status != status4 || status != status5 || status != status6 {
		return fmt.Errorf("status invalido")
	}

	err := cancelOption(ctx, *s, idOrder, status)
	if err != nil {
		return err
	}

	err = s.repository.UpdateStatus(ctx, *idOrder, status)
	if err != nil {
		return err
	}

	return nil
}
