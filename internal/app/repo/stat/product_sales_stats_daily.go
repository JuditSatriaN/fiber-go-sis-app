package stat

import (
	"github.com/fiber-go-sis-app/internal/app/model"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
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
