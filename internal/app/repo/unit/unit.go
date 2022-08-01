package unit

import (
	"database/sql"

	"github.com/fiber-go-sis-app/internal/app/model"
	"github.com/gofiber/fiber/v2"

	postgresPkg "github.com/fiber-go-sis-app/internal/pkg/databases/postgres"
)

const queryGetAllUnit = `
	SELECT id, name
	FROM units
	ORDER BY id
`

func GetAllUnit(ctx *fiber.Ctx) ([]model.Unit, error) {
	var units []model.Unit
	db := postgresPkg.GetPgConn()

	if err := db.SelectContext(ctx.Context(), &units, queryGetAllUnit); err != nil {
		return units, err
	}

	return units, nil
}

const queryGetUnitByID = `
	SELECT id, name
	FROM units
	WHERE id = $1
`

func GetUnitByID(ctx *fiber.Ctx, ID int) (model.Unit, bool, error) {
	var unit model.Unit

	db := postgresPkg.GetPgConn()

	if err := db.GetContext(ctx.Context(), &unit, queryGetUnitByID, ID); err != nil {
		if err == sql.ErrNoRows {
			return unit, false, nil
		}

		return unit, false, err
	}

	return unit, true, nil
}

const insertUnit = `
	INSERT INTO units (name)
	VALUES (:name)
`

func InsertUnit(ctx *fiber.Ctx, unit model.Unit) error {
	db := postgresPkg.GetPgConn()

	_, err := db.NamedQueryContext(ctx.Context(), insertUnit, unit)
	return err
}

const updateUnit = `
	UPDATE units SET
		name = :name,
		update_time = NOW()
	WHERE id = :id
`

func UpdateUnit(ctx *fiber.Ctx, unit model.Unit) error {
	db := postgresPkg.GetPgConn()

	_, err := db.NamedQueryContext(ctx.Context(), updateUnit, unit)
	return err
}

const deleteUnit = `
	DELETE FROM units
	WHERE id = $1
`

func DeleteUnit(ctx *fiber.Ctx, ID int) error {
	db := postgresPkg.GetPgConn()

	_, err := db.ExecContext(ctx.Context(), deleteUnit, ID)
	return err
}
