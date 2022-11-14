package domain

type Location struct {
	Id          *int    `json:"id"`
	Type        string  `json:"type" validate:"required"`
	Province    string  `json:"province" validate:"required"`
	City        string  `json:"city" validate:"required"`
	Commune     string  `json:"commune" validate:"required"`
	FullAddress string  `json:"full_address" validate:"required"`
	Lat         float64 `json:"lat" validate:"required"`
	Lng         float64 `json:"lng" validate:"required"`
}
