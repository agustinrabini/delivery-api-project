package order

import (
	"context"
	response "delivery-api-project/controllers/web/response"
	"delivery-api-project/domain"
	"fmt"
	"time"
)

func buildResponseOrder(order domain.Order, packages []domain.Package, delivery domain.Delivery, receiverLocation, remittentLocation domain.Location) (response.Order, error) {

	respondeDelivery := response.Delivery{
		Id:              *delivery.Id,
		OriginLocation:  remittentLocation,
		DestinyLocation: receiverLocation,
		PickUpDate:      delivery.PickUpDate,
		DeliveryDate:    delivery.DeliveryDate,
	}

	responseOrder := response.Order{
		Id:           *order.Id,
		ReceiverID:   order.ReceiverID,
		RemitterID:   order.RemitterID,
		Packages:     packages,
		Delivery:     respondeDelivery,
		Status:       order.Status,
		CreationDate: order.CreationDate,
	}

	return responseOrder, nil
}

//Returns a domain.Delivery{} with the pick up and delivery date and hour setted.
func buildDelivery(idRemittentLoc, idReceiverLoc int) (domain.Delivery, error) {

	date := time.Now()
	pickUpDate := date.AddDate(0, 0, 1)
	deliveryPickUpDate := date.AddDate(0, 0, 2)

	delivery := domain.Delivery{
		Id:                nil,
		IdOriginLocation:  idRemittentLoc,
		IdDestinyLocation: idReceiverLoc,
		PickUpDate:        pickUpDate.Format("2006-1-2 15:4:5"),
		DeliveryDate:      deliveryPickUpDate.Format("2006-1-2 15:4:5"),
	}

	return delivery, nil
}

func buildPackages(ctx context.Context, s service, pkgs []domain.Package, idOrder *int) error {

	for _, pk := range pkgs {

		if pk.Weight < 5 {
			pk.Size = "S"
		} else if pk.Weight > 5 && pk.Weight < 15 {
			pk.Size = "M"
		} else if pk.Weight > 15 && pk.Weight < 25 {
			pk.Size = "L"
		}

		pk.IdOrder = *idOrder

		err := s.packagesRepo.Create(ctx, pk)
		if err != nil {
			return err
		}
	}

	return nil
}

func cancelOption(ctx context.Context, s service, idOrder *int, status string) error {

	dateFormat := "2006-1-2 15:4:5"
	actualDate := time.Now()

	if status == "cancelado" {

		order, err := s.repository.Get(ctx, *idOrder)
		if err != nil {
			return err
		}

		if order.Status == status4 || order.Status == status5 {
			return fmt.Errorf("no se puede cancelar la orden")
		}

		orderDateFormated, err := time.Parse(dateFormat, order.CreationDate)
		if err != nil {
			return err
		}

		diff := actualDate.Sub(orderDateFormated)

		if diff.Minutes() >= 2 && orderDateFormated.Day() != actualDate.Day() {
			return fmt.Errorf("no se puede cancelar la orden")
		}
	}

	return nil
}
