package request

import "delivery-api-project/domain"

type Order struct {
	ReceiverID int       `json:"id_receiver" validate:"required"`
	RemitterID int       `json:"id_remitter" validate:"required"`
	Packages   []Package `json:"packages" validate:"required"`
	Delivery   Delivery  `json:"delivery" validate:"required"`
}

type Delivery struct {
	OriginLocation  domain.Location `json:"origin_location" validate:"required"`
	DestinyLocation domain.Location `json:"destiny_location" validate:"required"`
}

type Package struct {
	Weight        float32 `json:"weight" validate:"required,number,min=1,max=25"`
	QuantityItems int     `json:"quantity_items" validate:"required,min=1"`
}
