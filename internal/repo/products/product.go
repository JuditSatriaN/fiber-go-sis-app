package products

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"

	postgresPkg "github.com/fiber-go-sis-app/utils/pkg/databases/postgres"

	productEntity "github.com/fiber-go-sis-app/internal/entity/product"
)

const queryGetProductByID = `
	SELECT product_id, name, barcode, stock, ppn, price, member_price, discount, category_id
	FROM products
	WHERE product_id = $1
`

func GetProductByID(ctx *fiber.Ctx, productID string) (productEntity.Product, bool, error) {
	var product productEntity.Product
	db := postgresPkg.GetPgConn()

	if err := db.GetContext(ctx.Context(), &product, queryGetProductByID, productID); err != nil {
		if err == sql.ErrNoRows {
			return product, false, nil
		}
		return product, false, err
	}
	return product, true, nil
}

const queryGetProductByIDOrBarcode = `
	(SELECT product_id, name, barcode, stock, ppn, price, member_price, discount, category_id
	FROM products
	WHERE product_id = $1)
	UNION
	(SELECT product_id, name, barcode, stock, ppn, price, member_price, discount, category_id
	FROM products
	WHERE barcode = $1)
`

func GetProductByIDOrBarcode(ctx *fiber.Ctx, search string) (productEntity.Product, bool, error) {
	var product productEntity.Product
	db := postgresPkg.GetPgConn()

	if err := db.GetContext(ctx.Context(), &product, queryGetProductByIDOrBarcode, search); err != nil {
		if err == sql.ErrNoRows {
			return product, false, nil
		}
		return product, false, err
	}
	return product, true, nil
}

const insertProduct = `
	INSERT INTO products (product_id, name, barcode, stock, ppn, price, member_price, discount, category_id)
	VALUES (:product_id, :name, :barcode, :stock, :ppn, :price, :member_price, :discount, :category_id)
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
	    stock = :stock,
	    ppn = :ppn,
	    price = :price,
	    member_price = :member_price,
	    discount = :discount,
	    category_id = :category_id,
		update_time = NOW()
	WHERE product_id = :product_id
`

func UpdateProduct(ctx *fiber.Ctx, product productEntity.Product) error {
	db := postgresPkg.GetPgConn()

	_, err := db.NamedQueryContext(ctx.Context(), updateProduct, product)
	return err
}

const deleteProduct = `
	DELETE FROM products
	WHERE product_id = $1
`

func DeleteProduct(ctx *fiber.Ctx, productID string) error {
	db := postgresPkg.GetPgConn()

	_, err := db.ExecContext(ctx.Context(), deleteProduct, productID)
	return err
}
