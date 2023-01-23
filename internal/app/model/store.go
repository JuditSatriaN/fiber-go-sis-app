package model

import "time"

type StoreStats struct {
	StoreID        string `json:"store_id" db:"store_id" validate:"max=30"`
	TotalProduct   int64  `json:"total_product" db:"total_product"`
	TotalInventory int64  `json:"total_inventory" db:"total_inventory"`
}

type StatisticDashboard struct {
	TotalProductTerjualHariIni int64                    `json:"total_product_terjual_hari_ini"`
	TotalPemasukanHariIni      float64                  `json:"total_pemasukan_hari_ini"`
	TotalPendapatanHariIni     float64                  `json:"total_pendapatan_hari_ini"`
	TotalTransaksiHariIni      float64                  `json:"total_transaksi_hari_ini"`
	TotalPendapatanMonthly     []TotalPendapatanMonthly `json:"total_pendapatan_monthly"`
}

type TotalPendapatanMonthly struct {
	TimeTx         time.Time `db:"time_tx" json:"time_tx"`
	TotalPenjualan float64   `db:"total_penjualan" json:"total_penjualan"`
}
