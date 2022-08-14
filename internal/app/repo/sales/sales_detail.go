package sales

import (
	"github.com/fiber-go-sis-app/internal/app/model"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

const queryInsertSalesDetail = `
	INSERT INTO sales_detail (invoice, user_id, plu, name, unit_name, barcode, ppn, qty, price, purchase, discount)
	VALUES (:invoice, :user_id, :plu, :name, :unit_name, :barcode, :ppn, :qty, :price, :purchase, :discount)
`

func InsertSalesDetail(ctx *fiber.Ctx, tx *sqlx.Tx, salesDetails []model.SalesDetail) error {
	_, err := tx.NamedExecContext(ctx.Context(), queryInsertSalesDetail, salesDetails)
	return err
}
