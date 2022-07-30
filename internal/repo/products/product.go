package products

import (
	"database/sql"
	"fmt"

	"github.com/gofiber/fiber/v2"

	postgresPkg "github.com/fiber-go-sis-app/utils/pkg/databases/postgres"

	productEntity "github.com/fiber-go-sis-app/internal/entity/products"
)

const queryGetALlProducts = `
	SELECT plu, name, barcode, ppn
	FROM products
	WHERE $1 = '' OR value_text_search @@ to_tsquery($1)
	LIMIT $2
	OFFSET $3
`

func GetALlProducts(ctx *fiber.Ctx, search string, limit int, offset int) ([]productEntity.Product, error) {
	fmt.Println("SEARCH : ", search)
	var product []productEntity.Product
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

func GetProductByPLU(ctx *fiber.Ctx, PLU string) (productEntity.Product, bool, error) {
	var product productEntity.Product
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

func GetProductByPLUOrBarcode(ctx *fiber.Ctx, search string) (productEntity.Product, bool, error) {
	var product productEntity.Product
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

func InsertProduct(ctx *fiber.Ctx, product productEntity.Product) error {
	db := postgresPkg.GetPgConn()

	_, err := db.NamedQueryContext(ctx.Context(), insertProduct, product)
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

func UpdateProduct(ctx *fiber.Ctx, product productEntity.Product) error {
	db := postgresPkg.GetPgConn()

	_, err := db.NamedQueryContext(ctx.Context(), updateProduct, product)
	return err
}

const deleteProduct = `
	DELETE FROM products
	WHERE plu = $1
`

func DeleteProduct(ctx *fiber.Ctx, productID string) error {
	db := postgresPkg.GetPgConn()

	_, err := db.ExecContext(ctx.Context(), deleteProduct, productID)
	return err
}
