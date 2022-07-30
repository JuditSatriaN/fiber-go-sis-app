package stores

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"

	storeEntity "github.com/fiber-go-sis-app/internal/entity/stores"

	postgresPkg "github.com/fiber-go-sis-app/utils/pkg/databases/postgres"
)

const queryGetStoreStats = `
	SELECT store_id, cnt_product
	FROM store_stats
	WHERE store_id = $1
`

func GetStoreStats(ctx *fiber.Ctx, storeID string) (storeEntity.StoreStats, error) {
	var storeStats storeEntity.StoreStats
	db := postgresPkg.GetPgConn()

	if err := db.GetContext(ctx.Context(), &storeStats, queryGetStoreStats, storeID); err != nil {
		if err == sql.ErrNoRows {
			return storeStats, nil
		}
		return storeStats, err
	}
	return storeStats, nil
}

const updateTotalProduct = `
	UPDATE store_stats 
	SET
		cnt_product = cnt_product + :cnt_product,
		update_time = NOW()
	WHERE store_id = :store_id
`

func UpdateTotalProduct(ctx *fiber.Ctx, store storeEntity.StoreStats) error {
	db := postgresPkg.GetPgConn()

	_, err := db.NamedQueryContext(ctx.Context(), updateTotalProduct, store)
	return err
}
