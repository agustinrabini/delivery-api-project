package domain

const (
	Status1 = "creado"

	Status2 = "	recolectado"

	Status3 = "en_estacion"

	Status4 = "en_ruta"

	Status5 = "entregado"

	Status6 = "cancelado"
)

type Order struct {
	Id           *int   `json:"id"`
	IdDelivery   int    `json:"id_delivery"`
	ReceiverID   int    `json:"id_receiver"`
	RemitterID   int    `json:"id_remitter"`
	Status       string `json:"status"`
	CreationDate string `json:"creation_date"`
}
