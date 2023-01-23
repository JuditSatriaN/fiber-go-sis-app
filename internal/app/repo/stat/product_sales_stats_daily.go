package stat

import (
	"database/sql"
	"github.com/fiber-go-sis-app/internal/app/model"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"

	postgresPkg "github.com/fiber-go-sis-app/internal/pkg/database/postgres"
)

const queryUpsertTotalSold = `
	INSERT INTO product_sales_stats_daily (date_sold, plu, total_sold)
	VALUES (:date_sold, :plu, :total_sold)
	ON CONFLICT (date_sold, plu) DO UPDATE
		SET total_sold  = product_sales_stats_daily.total_sold + EXCLUDED.total_sold,
			update_time = NOW();
`

func BulkUpsertTotalSold(ctx *fiber.Ctx, tx *sqlx.Tx, data []model.ProductSalesStatsDaily) error {
	_, err := tx.NamedExecContext(ctx.Context(), queryUpsertTotalSold, data)
	if err != nil {
		return err
	}

	return err
}

const queryGetProductSalesStatsDaily = `
	SELECT pd.plu, p.name, pd.total_sold
	FROM product_sales_stats_daily pd
	JOIN products p ON pd.plu = p.plu
	WHERE pd.date_sold = CURRENT_DATE
	ORDER BY pd.total_sold desc
	LIMIT $1;
`

func GetProductSalesStatsDaily(ctx *fiber.Ctx, limit int) ([]model.ProductSalesStatsDaily, error) {
	var productSales []model.ProductSalesStatsDaily
	db := postgresPkg.GetPgConn()
	if err := db.SelectContext(ctx.Context(), &productSales, queryGetProductSalesStatsDaily, limit); err != nil {
		return []model.ProductSalesStatsDaily{}, err
	}

	return productSales, nil
}

const queryGetTotalProductSoldToday = `
	SELECT COUNT(*)
	FROM product_sales_stats_daily pd
	JOIN products p ON pd.plu = p.plu
	WHERE pd.date_sold = CURRENT_DATE AND pd.total_sold > 0
	LIMIT 1;
`

func GetTotalProductSoldToday(ctx *fiber.Ctx) (int64, error) {
	var totalProduct int64
	db := postgresPkg.GetPgConn()
	if err := db.GetContext(ctx.Context(), &totalProduct, queryGetTotalProductSoldToday); err != nil {
		if err == sql.ErrNoRows {
			return totalProduct, nil
		}
		return totalProduct, nil
	}

	return totalProduct, nil
}

const queryGetTotalPemasukanToday = `
	SELECT SUM(total_pay)
	FROM sales_head
	WHERE create_time::date = CURRENT_DATE
	LIMIT 1;
`

func GetTotalPemasukanToday(ctx *fiber.Ctx) (float64, error) {
	var totalPemasukan float64
	db := postgresPkg.GetPgConn()
	if err := db.GetContext(ctx.Context(), &totalPemasukan, queryGetTotalPemasukanToday); err != nil {
		if err == sql.ErrNoRows {
			return totalPemasukan, nil
		}
		return totalPemasukan, nil
	}

	return totalPemasukan, nil
}

const queryGetTotalTransaksiToday = `
	SELECT COUNT(*)
	FROM sales_head
	WHERE create_time::date = CURRENT_DATE
	LIMIT 1;
`

func GetTotalTransaksiToday(ctx *fiber.Ctx) (float64, error) {
	var totalTransaksi float64
	db := postgresPkg.GetPgConn()
	if err := db.GetContext(ctx.Context(), &totalTransaksi, queryGetTotalTransaksiToday); err != nil {
		if err == sql.ErrNoRows {
			return totalTransaksi, nil
		}
		return totalTransaksi, nil
	}

	return totalTransaksi, nil
}

const queryGetTotalPendapatanToday = `
	SELECT SUM(total_purchase - total_price)
	FROM sales_head
	WHERE create_time::date = CURRENT_DATE
	LIMIT 1;
`

func GetTotalPendapatanToday(ctx *fiber.Ctx) (float64, error) {
	var totalPendapatan float64
	db := postgresPkg.GetPgConn()
	if err := db.GetContext(ctx.Context(), &totalPendapatan, queryGetTotalPendapatanToday); err != nil {
		if err == sql.ErrNoRows {
			return totalPendapatan, nil
		}
		return totalPendapatan, nil
	}

	return totalPendapatan, nil
}
