package model

type Inventory struct {
	ID          int32   `json:"id" db:"id"`
	PLU         string  `json:"plu" db:"plu" validate:"max=30"`
	Name        string  `json:"name" db:"name" validate:"max=255"`
	UnitID      int32   `json:"unit_id" db:"unit_id"`
	UnitName    string  `json:"unit_name" db:"unit_name" validate:"max=30"`
	Barcode     string  `json:"barcode" db:"barcode" validate:"max=30"`
	Ppn         bool    `json:"ppn" db:"ppn"`
	Multiplier  int64   `json:"multiplier" db:"multiplier"`
	Stock       int64   `json:"stock" db:"stock"`
	Price       float32 `json:"price" db:"price"`
	MemberPrice float32 `json:"member_price" db:"member_price"`
	Purchase    float32 `json:"purchase" db:"purchase"`
	Discount    float32 `json:"discount" db:"discount"`
}

type ListInventoryDataResponse struct {
	Total int64
	Data  []Inventory
}
