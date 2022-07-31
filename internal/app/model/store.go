package model

const DefaultStoreID = "1"

type StoreStats struct {
	StoreID      string `json:"store_id" db:"store_id"`
	TotalProduct int64  `json:"total_product" db:"total_product"`
}
