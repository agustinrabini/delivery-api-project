package request

import "delivery-api-project/domain"

type Order struct {
	ReceiverID   int              `json:"id_receiver" validate:"required"`
	RemitterID   int              `json:"id_remitter" validate:"required"`
	Packages     []domain.Package `json:"package" validate:"required"`
	Delivery     Delivery         `json:"delivery" validate:"required"`
	Status       string           `json:"status" validate:"required"`
	CreationDate string           `json:"creation_date" validate:"required"`
}

type Delivery struct {
	OriginLocation  domain.Location `json:"origin_location" validate:"required"`
	DestinyLocation domain.Location `json:"destiny_address" validate:"required"`
}
