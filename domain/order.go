package domain

type Order struct {
	Id           *int    `json:"id"`
	IdDelivery   int     `json:"id_delivery"`
	ReceiverID   int     `json:"id_receiver"`
	RemitterID   int     `json:"id_remitter"`
	Status       string  `json:"status"`
	CreationDate string  `json:"creation_date"`
	Price        float32 `json:"price"`
}
