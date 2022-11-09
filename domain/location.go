package domain

type Location struct {
	Id          int     `json:"id"`
	Province    string  `json:"province"`
	City        string  `json:"city"`
	Commune     string  `json:"commune"`
	FullAddress string  `json:"full_address"`
	Lat         float64 `json:"lat"`
	Lng         float64 `json:"lng"`
}
