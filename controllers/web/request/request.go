package request

import "delivery-api-project/domain"

type Order struct {
	ReceiverID   int              `json:"id_receiver"`
	RemitterID   int              `json:"id_remitter"`
	Packages     []domain.Package `json:"package"`
	Delivery     Delivery         `json:"delivery"`
	Status       string           `json:"status"`
	CreationDate string           `json:"creation_date"`
}

type Delivery struct {
	OriginLocation  domain.Location `json:"origin_location"`
	DestinyLocation domain.Location `json:"destiny_address"`
}
