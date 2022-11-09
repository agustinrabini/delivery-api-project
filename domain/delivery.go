package domain

type Delivery struct {
	Id              int      `json:"id"`
	OriginLocation  Location `json:"origin_location"`
	DestinyLocation Location `json:"destiny_address"`
	PickUpDate      string   `json:"pick_up_date"`
	DeliveryDate    string   `json:"delivery_date"`
}
