package order

import (
	response "delivery-api-project/controllers/web"
	"delivery-api-project/domain"
)

func buildResponseOrder(order domain.Order, packages []domain.Package, delivery domain.Delivery, receiverLocation, remittentLocation domain.Location) (response.Order, error) {

	respondeDelivery := response.Delivery{
		Id:              delivery.Id,
		OriginLocation:  remittentLocation,
		DestinyLocation: receiverLocation,
		PickUpDate:      delivery.PickUpDate,
		DeliveryDate:    delivery.DeliveryDate,
	}

	responseOrder := response.Order{
		Id:           order.Id,
		ReceiverID:   order.ReceiverID,
		RemitterID:   order.RemitterID,
		Packages:     packages,
		Delivery:     respondeDelivery,
		Status:       order.Status,
		CreationDate: order.CreationDate,
	}

	return responseOrder, nil
}
