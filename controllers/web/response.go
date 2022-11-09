package web

import "delivery-api-project/domain"

type Order struct {
	Id           int              `json:"id"`
	ReceiverID   int              `json:"id_receiver"`
	RemitterID   int              `json:"id_remitter"`
	Packages     []domain.Package `json:"package"`
	Delivery     Delivery         `json:"delivery"`
	Status       string           `json:"status"`
	CreationDate string           `json:"creation_date"`
}

type Delivery struct {
	Id              int             `json:"id"`
	OriginLocation  domain.Location `json:"origin_location"`
	DestinyLocation domain.Location `json:"destiny_address"`
	PickUpDate      string          `json:"pick_up_date"`
	DeliveryDate    string          `json:"delivery_date"`
}
