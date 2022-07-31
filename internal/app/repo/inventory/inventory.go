package inventory

import (
	"github.com/fiber-go-sis-app/internal/app/model"
	"github.com/gofiber/fiber/v2"

	postgresPkg "github.com/fiber-go-sis-app/internal/pkg/databases/postgres"
)

const queryGetALlInventory = `
	SELECT p.plu, p.name, p.barcode, p.ppn, pd.multiplier, pd.stock, 
	       pd.price, pd.member_price, pd.discount
	FROM product_details pd
	INNER JOIN products p on pd.plu = p.plu
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
