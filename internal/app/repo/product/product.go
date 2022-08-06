package product

import (
	"database/sql"

	"github.com/fiber-go-sis-app/internal/app/model"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"

	postgresPkg "github.com/fiber-go-sis-app/internal/pkg/databases/postgres"
)

const queryGetALlProducts = `
	SELECT plu, name, barcode, ppn
	FROM products
	WHERE $1 = '' OR value_text_search @@ to_tsquery($1)
	LIMIT $2
	OFFSET $3
`

func GetALlProducts(ctx *fiber.Ctx, search string, limit int, offset int) ([]model.Product, error) {
	var product []model.Product
	db := postgresPkg.GetPgConn()

	if err := db.SelectContext(ctx.Context(), &product, queryGetALlProducts, search, limit, offset); err != nil {
		return product, err
	}

	return product, nil
}

const queryGetProductByPLU = `
	SELECT plu, name, barcode, ppn
	FROM products
	WHERE plu = $1
`

func GetProductByPLU(ctx *fiber.Ctx, PLU string) (model.Product, bool, error) {
	var product model.Product
	db := postgresPkg.GetPgConn()

	if err := db.GetContext(ctx.Context(), &product, queryGetProductByPLU, PLU); err != nil {
		if err == sql.ErrNoRows {
			return product, false, nil
		}

		return product, false, err
	}

	return product, true, nil
}

const queryGetProductByPLUOrBarcode = `
	(SELECT plu, name, barcode, ppn
	FROM products
	WHERE plu = $1)
	UNION
	(SELECT plu, name, barcode, ppn
	FROM products
	WHERE barcode = $1)
`

func GetProductByPLUOrBarcode(ctx *fiber.Ctx, search string) (model.Product, bool, error) {
	var product model.Product
	db := postgresPkg.GetPgConn()

	if err := db.GetContext(ctx.Context(), &product, queryGetProductByPLUOrBarcode, search); err != nil {
		if err == sql.ErrNoRows {
			return product, false, nil
		}

		return product, false, err
	}

	return product, true, nil
}

const insertProduct = `
	INSERT INTO products (plu, name, barcode, ppn)
	VALUES (:plu, :name, :barcode, :ppn)
`

func InsertProduct(tx *sqlx.Tx, product model.Product) error {
	_, err := tx.NamedQuery(insertProduct, product)
	return err
}

const updateProduct = `
	UPDATE products SET
		name = :name,
	    barcode = :barcode,
	    ppn = :ppn,
		update_time = NOW()
	WHERE plu = :plu
`

func UpdateProduct(ctx *fiber.Ctx, product model.Product) error {
	db := postgresPkg.GetPgConn()

	_, err := db.NamedQueryContext(ctx.Context(), updateProduct, product)
	return err
}

const deleteProduct = `
	DELETE FROM products
	WHERE plu = $1
`

func DeleteProduct(tx *sqlx.Tx, productID string) error {
	_, err := tx.Exec(deleteProduct, productID)
	return err
}
