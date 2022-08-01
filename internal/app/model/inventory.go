package model

type Inventory struct {
	ID          int     `json:"id" db:"id"`
	PLU         string  `json:"plu" db:"plu" validate:"max=20"`
	Name        string  `json:"name" db:"name" validate:"max=255"`
	UnitName    string  `json:"unit_name" db:"unit_name" validate:"max=255"`
	Barcode     string  `json:"barcode" db:"barcode"`
	Ppn         bool    `json:"ppn" db:"ppn"`
	Multiplier  int64   `json:"multiplier" db:"multiplier"`
	Stock       int64   `json:"stock" db:"stock"`
	Price       float32 `json:"price" db:"price"`
	MemberPrice float32 `json:"member_price" db:"member_price"`
	Discount    float32 `json:"discount" db:"discount"`
}

type ListInventoryDataResponse struct {
	Total int64
	Data  []Inventory
}
