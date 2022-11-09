package domain

type Delivery struct {
	Id                int    `json:"id"`
	IdOrder           int    `json:"id_order"`
	IdOriginLocation  int    `json:"id_origin_location"`
	IdDestinyLocation int    `json:"id_destiny_address"`
	PickUpDate        string `json:"pick_up_date"`
	DeliveryDate      string `json:"delivery_date"`
}
