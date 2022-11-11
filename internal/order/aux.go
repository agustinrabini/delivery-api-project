package order

import (
	"context"
	"delivery-api-project/controllers/web/request"
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

func buildPackages(ctx context.Context, s service, rrpkgs []request.Package, idOrder *int) error {

	for _, rrpk := range rrpkgs {

		pkg := domain.Package{Weight: rrpk.Weight}

		if rrpk.Weight < 5 {
			pkg.Size = "S"
		} else if rrpk.Weight > 5 && rrpk.Weight < 15 {
			pkg.Size = "M"
		} else if rrpk.Weight > 15 && rrpk.Weight < 25 {
			pkg.Size = "L"
		}

		pkg.IdOrder = *idOrder

		err := s.packagesRepo.Create(ctx, pkg)
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
