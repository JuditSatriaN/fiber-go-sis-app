package stores

const DefaultStoreID = "1"

type StoreStats struct {
	StoreID    string `json:"store_id" db:"store_id"`
	CntProduct int64  `json:"cnt_product" db:"cnt_product"`
}
