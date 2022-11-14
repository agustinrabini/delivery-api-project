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
	status1 string = "creado"
	status2 string = "recolectado"
	status3 string = "en_estacion"
	status4 string = "en_ruta"
	status5 string = "entregado"
	status6 string = "cancelado"
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
	if order.Id == nil {
		return response.Order{}, nil
	}
	if err != nil {
		return response.Order{}, fmt.Errorf("error getting the order: %v", err)
	}

	packages, err := s.packagesRepo.GetPackagesByOrder(ctx, *order.Id)
	if err != nil {
		return response.Order{}, fmt.Errorf("error getting the packages order: %v", err)
	}

	delivery, err := s.deliveryRepo.GetDeliveryByOrder(ctx, order.IdDelivery)
	if err != nil {
		return response.Order{}, fmt.Errorf("error getting the delivery order: %v", err)
	}

	receiverLocation, remittentLocation, err := s.locationRepo.GetReceiverAndRemittentLocation(ctx, delivery.IdOriginLocation, delivery.IdDestinyLocation)
	if err != nil {
		return response.Order{}, fmt.Errorf("error getting the delivery locations order: %v", err)
	}

	responseOrder, err := buildResponseOrder(order, packages, delivery, receiverLocation, remittentLocation)
	if err != nil {
		return response.Order{}, fmt.Errorf("error building the response order: %v", err)
	}

	return responseOrder, nil
}

func (s *service) Create(ctx context.Context, rrOrder request.Order) (*int, error) {

	receiverLocID, err := s.locationRepo.Create(ctx, rrOrder.Delivery.DestinyLocation, "receiver")
	if err != nil {
		return nil, fmt.Errorf("error creating the receiver location: %v", err)
	}

	remittentLocID, err := s.locationRepo.Create(ctx, rrOrder.Delivery.OriginLocation, "remitter")
	if err != nil {
		return nil, fmt.Errorf("error creating the remittent location: %v", err)
	}

	delivery, err := buildDelivery(*remittentLocID, *receiverLocID)
	if err != nil {
		return nil, fmt.Errorf("error building the delivery: %v", err)
	}

	deliveryID, err := s.deliveryRepo.Create(ctx, delivery)
	if err != nil {
		return nil, fmt.Errorf("error creating the delivery order: %v", err)
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
		return nil, fmt.Errorf("error creating the order: %v", err)
	}

	err = buildPackages(ctx, *s, rrOrder.Packages, orderId)
	if err != nil {
		return nil, fmt.Errorf("error building the packages order: %v", err)
	}

	return orderId, nil
}

func (s *service) UpdateStatus(ctx context.Context, idOrder *int, status string) error {

	if status != status1 && status != status2 && status != status3 && status != status4 && status != status5 && status != status6 {
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
