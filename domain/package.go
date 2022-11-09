package domain

type Package struct {
	Id            int     `json:"id"`
	Weight        float32 `json:"weight"`
	Size          string  `json:"size"`
	QuantityItems int     `json:"quantity_items"`
}
