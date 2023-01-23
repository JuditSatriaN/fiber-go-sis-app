package model

type StoreStats struct {
	StoreID        string `json:"store_id" db:"store_id" validate:"max=30"`
	TotalProduct   int64  `json:"total_product" db:"total_product"`
	TotalInventory int64  `json:"total_inventory" db:"total_inventory"`
}

type StatisticDashboard struct {
	TotalProductTerjualHariIni int64   `json:"total_product_terjual_hari_ini"`
	TotalPemasukanHariIni      float64 `json:"total_pemasukan_hari_ini"`
	TotalPendapatanHariIni     float64 `json:"total_pendapatan_hari_ini"`
	TotalTransaksiHariIni      float64 `json:"total_transaksi_hari_ini"`
}
