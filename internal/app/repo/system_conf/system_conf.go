package system_conf

import (
	"database/sql"

	"github.com/fiber-go-sis-app/internal/app/model"
	postgresPkg "github.com/fiber-go-sis-app/internal/pkg/database/postgres"
	"github.com/gofiber/fiber/v2"
)

const queryGetSystemConf = `
	SELECT id, value
	FROM system_conf
	ORDER BY id
`

func GetAllSystemConf(ctx *fiber.Ctx) ([]model.SystemConf, error) {
	var systemConf []model.SystemConf
	db := postgresPkg.GetPgConn()

	if err := db.SelectContext(ctx.Context(), &systemConf, queryGetSystemConf); err != nil {
		return systemConf, err
	}

	return systemConf, nil
}

const queryGetKeyVoid = `
	SELECT value
	FROM system_conf
	WHERE id = 'password_void'
`

func GetKeyVoid(ctx *fiber.Ctx) (string, error) {
	var data string

	db := postgresPkg.GetPgConn()

	if err := db.GetContext(ctx.Context(), &data, queryGetKeyVoid); err != nil {
		if err == sql.ErrNoRows {
			return data, nil
		}

		return data, err
	}

	return data, nil
}

const queryGetSystemConfByID = `
	SELECT id, value
	FROM system_conf
	WHERE id = $1
`

func GetSystemConfByID(ctx *fiber.Ctx, ID string) (model.SystemConf, bool, error) {
	var systemConf model.SystemConf

	db := postgresPkg.GetPgConn()

	if err := db.GetContext(ctx.Context(), &systemConf, queryGetSystemConfByID, ID); err != nil {
		if err == sql.ErrNoRows {
			return systemConf, false, nil
		}

		return systemConf, false, err
	}

	return systemConf, true, nil
}

const queryInsertSystemConf = `
	INSERT INTO system_conf (id, value)
	VALUES (:id, :value)
`

func InsertSystemConf(ctx *fiber.Ctx, data model.SystemConf) error {
	db := postgresPkg.GetPgConn()
	_, err := db.NamedQueryContext(ctx.Context(), queryInsertSystemConf, data)
	return err
}

const queryUpdateSystemConf = `
	UPDATE system_conf 
	SET value = :value
	WHERE id = :id
`

func UpdateSystemConf(ctx *fiber.Ctx, data model.SystemConf) error {
	db := postgresPkg.GetPgConn()
	_, err := db.NamedQueryContext(ctx.Context(), queryUpdateSystemConf, data)
	return err
}

const queryDeleteSystemConf = `
	DELETE FROM system_conf
	WHERE id = $1
`

func DeleteSystemConf(ctx *fiber.Ctx, ID string) error {
	db := postgresPkg.GetPgConn()
	_, err := db.ExecContext(ctx.Context(), queryDeleteSystemConf, ID)
	return err
}
