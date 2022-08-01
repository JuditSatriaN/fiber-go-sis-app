package inventory

import (
	"database/sql"
	"github.com/fiber-go-sis-app/internal/app/model"
	"github.com/gofiber/fiber/v2"

	postgresPkg "github.com/fiber-go-sis-app/internal/pkg/databases/postgres"
)

const queryGetALlInventory = `
	SELECT i.id, p.plu, p.name, u.name AS unit_name, p.barcode, p.ppn, i.multiplier, i.stock, 
	       i.price, i.member_price, i.discount
	FROM inventory i
	INNER JOIN products p on i.plu = p.plu
	INNER JOIN units u on i.unit_id = u.id
	WHERE $1 = '' OR p.value_text_search @@ to_tsquery($1)
	LIMIT $2
	OFFSET $3
`

func GetALlInventory(ctx *fiber.Ctx, search string, limit int, offset int) ([]model.Inventory, error) {
	var inventory []model.Inventory
	db := postgresPkg.GetPgConn()

	if err := db.SelectContext(ctx.Context(), &inventory, queryGetALlInventory, search, limit, offset); err != nil {
		return inventory, err
	}
	return inventory, nil
}

const queryGetInventoryByID = `
	SELECT id, plu
	FROM inventory
	WHERE id = $1
`

func GetInventoryByID(ctx *fiber.Ctx, ID int) (model.Inventory, bool, error) {
	var inventory model.Inventory

	db := postgresPkg.GetPgConn()

	if err := db.GetContext(ctx.Context(), &inventory, queryGetInventoryByID, ID); err != nil {
		if err == sql.ErrNoRows {
			return inventory, false, nil
		}

		return inventory, false, err
	}

	return inventory, true, nil
}

const updateStockInventory = `
	UPDATE inventory 
	SET stock = :stock
	WHERE id = :id
`

func UpdateStockInventory(ctx *fiber.Ctx, inventory model.Inventory) error {
	db := postgresPkg.GetPgConn()

	_, err := db.NamedQueryContext(ctx.Context(), updateStockInventory, inventory)
	return err
}
