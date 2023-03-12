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
	                          discount, member_id, inventory_id)
	VALUES (:invoice, :user_id, :plu, :name, :unit_name, :barcode, :ppn, :qty, :price, :purchase, 
	        :discount, :member_id, :inventory_id)
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
	       total_tax, total_discount, total_pay, create_time
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
	SELECT id, invoice, user_id, plu, name, unit_name, barcode, ppn, qty, price, purchase, 
	       discount, member_id, inventory_id, create_time
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

const queryGetSalesHeadByInvoice = `
	SELECT id, invoice, user_id, total_item, total_price, total_purchase,
	       total_tax, total_discount, total_pay, create_time
	FROM sales_head
	WHERE invoice = $1
`

func GetSalesHeadByInvoice(ctx *fiber.Ctx, invoice string) (model.SalesHead, bool, error) {
	var data model.SalesHead

	db := postgresPkg.GetPgConn()

	if err := db.GetContext(ctx.Context(), &data, queryGetSalesHeadByInvoice, invoice); err != nil {
		if err == sql.ErrNoRows {
			return data, false, nil
		}

		return data, false, err
	}

	return data, true, nil
}

const queryInsertVoidHead = `
	INSERT INTO void_head (invoice, user_id, total_item, total_price, total_purchase, 
	                       total_tax, total_discount, total_pay)
	VALUES (:invoice, :user_id, :total_item, :total_price, :total_purchase, 
	        :total_tax, :total_discount, :total_pay)
`

func InsertVoidHead(ctx *fiber.Ctx, tx *sqlx.Tx, voidHead model.VoidHead) error {
	_, err := tx.NamedExecContext(ctx.Context(), queryInsertVoidHead, voidHead)
	return err
}

const queryInsertVoidDetail = `
	INSERT INTO void_detail (invoice, user_id, plu, name, unit_name, barcode, ppn, qty, price, purchase, 
	                          discount, member_id, inventory_id)
	VALUES (:invoice, :user_id, :plu, :name, :unit_name, :barcode, :ppn, :qty, :price, :purchase, 
	        :discount, :member_id, :inventory_id)
`

func InsertVoidDetail(ctx *fiber.Ctx, tx *sqlx.Tx, voidDetails []model.VoidDetail) error {
	_, err := tx.NamedExecContext(ctx.Context(), queryInsertVoidDetail, voidDetails)
	return err
}

const queryDeleteSalesDetailByInvoice = `
	DELETE FROM sales_detail
	WHERE invoice = $1
`

func DeleteSalesDetailByInvoice(ctx *fiber.Ctx, tx *sqlx.Tx, salesDetails []model.SalesDetail) error {
	for _, v := range salesDetails {
		_, err := tx.ExecContext(ctx.Context(), queryDeleteSalesDetailByInvoice, v.Invoice)
		if err != nil {
			return err
		}
	}
	return nil
}

const queryDeleteSalesHeadByInvoice = `
	DELETE FROM sales_head
	WHERE invoice = :invoice
`

func DeleteSalesHeadByInvoice(ctx *fiber.Ctx, tx *sqlx.Tx, salesHead model.SalesHead) error {
	_, err := tx.NamedExecContext(ctx.Context(), queryDeleteSalesHeadByInvoice, salesHead)
	return err
}
