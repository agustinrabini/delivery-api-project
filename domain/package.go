package domain

type Package struct {
	Id            int     `json:"id"`
	IdOrder       int     `json:"id_order" validate:"required"`
	Weight        float32 `json:"weight"validate:"required,number,min=1,max=25"`
	Size          string  `json:"size" validate:"required"`
	QuantityItems int     `json:"quantity_items" validate:"required"`
}
