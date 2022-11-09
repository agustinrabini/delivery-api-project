package domain

type Package struct {
	Id            int     `json:"id"`
	IdOrder       int     `json:"id_order"`
	Weight        float32 `json:"weight"`
	Size          string  `json:"size"`
	QuantityItems int     `json:"quantity_items"`
}
