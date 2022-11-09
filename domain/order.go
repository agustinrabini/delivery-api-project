package domain

const (
	status1 = "creado"

	status2 = "	recolectado"

	status3 = "en_estacion"

	status4 = "en_ruta"

	status5 = "entregado"

	status6 = "cancelado"
)

type Order struct {
	Id           int       `json:"id"`
	Packages     []Package `json:"package"`
	Delivery     Delivery  `json:"delivery"`
	Status       string    `json:"status"`
	CreationDate string    `json:"creation_date"`
}
