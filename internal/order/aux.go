package order

import (
	response "delivery-api-project/controllers/web/response"
	"delivery-api-project/domain"
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
