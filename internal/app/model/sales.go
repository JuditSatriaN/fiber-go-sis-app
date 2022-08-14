package model

type Sales struct {
	Detail []SalesDetail `json:"sales_detail"`
}

type SalesDetail struct {
	PLU      string  `json:"plu" db:"plu" validate:"max=20"`
	UnitID   int32   `json:"unit_id" db:"unit_id"`
	Barcode  string  `json:"barcode" db:"barcode"`
	Ppn      bool    `json:"ppn" db:"ppn"`
	Qty      int64   `json:"qty" db:"qty"`
	Price    float32 `json:"price" db:"price"`
	Purchase float32 `json:"purchase" db:"purchase"`
	Discount float32 `json:"discount" db:"discount"`
}
