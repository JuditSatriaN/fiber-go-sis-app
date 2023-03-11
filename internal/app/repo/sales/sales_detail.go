package sales

import (
	"database/sql"
	"github.com/fiber-go-sis-app/internal/app/model"
	postgresPkg "github.com/fiber-go-sis-app/internal/pkg/database/postgres"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

const queryInsertSalesDetail = `
	INSERT INTO sales_detail (invoice, user_id, plu, name, unit_name, barcode, ppn, qty, price, purchase, 
	                          discount, member_id)
	VALUES (:invoice, :user_id, :plu, :name, :unit_name, :barcode, :ppn, :qty, :price, :purchase, :discount, :member_id)
`

func InsertSalesDetail(ctx *fiber.Ctx, tx *sqlx.Tx, salesDetails []model.SalesDetail) error {
	_, err := tx.NamedExecContext(ctx.Context(), queryInsertSalesDetail, salesDetails)
	return err
}

const queryGetTotalSalesHead = `
	SELECT COUNT(*) OVER (ROWS BETWEEN CURRENT ROW AND 1000 FOLLOWING) AS total_count
	FROM sales_head
	WHERE $1 = '' OR value_text_search @@ plainto_tsquery($1)
`

func GetTotalSalesHead(ctx *fiber.Ctx, search string) (int64, error) {
	var totalSalesHead int64
	db := postgresPkg.GetPgConn()
	if err := db.GetContext(ctx.Context(), &totalSalesHead, queryGetTotalSalesHead, search); err != nil {
		if err == sql.ErrNoRows {
			return totalSalesHead, nil
		}

		return totalSalesHead, err
	}

	return totalSalesHead, nil
}

const queryGetALlSalesHead = `
	SELECT id, invoice, user_id, total_item, total_price, total_purchase,
	       total_tax, total_discount, total_pay
	FROM sales_head
	WHERE $1 = '' OR value_text_search @@ plainto_tsquery($1)
	ORDER BY id DESC 
	LIMIT $2
	OFFSET $3
`

func GetALlSalesHead(ctx *fiber.Ctx, search string, limit int, offset int) ([]model.SalesHead, error) {
	var salesHead []model.SalesHead
	db := postgresPkg.GetPgConn()
	if err := db.SelectContext(ctx.Context(), &salesHead, queryGetALlSalesHead, search, limit, offset); err != nil {
		return salesHead, err
	}

	return salesHead, nil
}

const queryGetALlSalesDetailByInvoice = `
	SELECT id, invoice, user_id, plu, name, unit_name, barcode, ppn, qty, price, purchase, discount, member_id
	FROM sales_detail
	WHERE invoice = $1
	ORDER BY id
`

func GetALlSalesDetailByInvoice(ctx *fiber.Ctx, invoice string) ([]model.SalesDetail, error) {
	var salesDetail []model.SalesDetail
	db := postgresPkg.GetPgConn()
	if err := db.SelectContext(ctx.Context(), &salesDetail, queryGetALlSalesDetailByInvoice, invoice); err != nil {
		return salesDetail, err
	}

	return salesDetail, nil
}
