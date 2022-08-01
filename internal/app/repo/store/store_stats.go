package store

import (
	"database/sql"
	"github.com/fiber-go-sis-app/internal/app/model"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"

	postgresPkg "github.com/fiber-go-sis-app/internal/pkg/databases/postgres"
)

const queryGetStoreStats = `
	SELECT store_id, total_product
	FROM store_stats
	WHERE store_id = $1
`

func GetStoreStats(ctx *fiber.Ctx, storeID string) (model.StoreStats, error) {
	var storeStats model.StoreStats
	db := postgresPkg.GetPgConn()

	if err := db.GetContext(ctx.Context(), &storeStats, queryGetStoreStats, storeID); err != nil {
		if err == sql.ErrNoRows {
			return storeStats, nil
		}

		return storeStats, err
	}

	return storeStats, nil
}

const queryGetTotalProduct = `
	SELECT total_product
	FROM store_stats
	WHERE store_id = $1
`

func GetTotalProduct(ctx *fiber.Ctx, storeID string) (int64, error) {
	var totalProduct int64
	db := postgresPkg.GetPgConn()

	if err := db.GetContext(ctx.Context(), &totalProduct, queryGetTotalProduct, storeID); err != nil {
		if err == sql.ErrNoRows {
			return totalProduct, nil
		}

		return totalProduct, err
	}

	return totalProduct, nil
}

const updateTotalProduct = `
	UPDATE store_stats 
	SET
		total_product = total_product + :total_product,
		update_time = NOW()
	WHERE store_id = :store_id
`

func UpdateTotalProduct(tx *sqlx.Tx, store model.StoreStats) error {
	rows, err := tx.NamedQuery(updateTotalProduct, store)
	defer rows.Close()
	return err
}
