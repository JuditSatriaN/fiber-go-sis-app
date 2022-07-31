package member

import (
	"database/sql"

	"github.com/fiber-go-sis-app/internal/app/model"
	"github.com/gofiber/fiber/v2"

	postgresPkg "github.com/fiber-go-sis-app/internal/pkg/databases/postgres"
)

const queryGetAllMember = `
	SELECT id, name, phone
	FROM members
	ORDER BY id
`

func GetAllMember(ctx *fiber.Ctx) ([]model.Member, error) {
	var members []model.Member
	db := postgresPkg.GetPgConn()

	if err := db.SelectContext(ctx.Context(), &members, queryGetAllMember); err != nil {
		return members, err
	}

	return members, nil
}

const queryGetMemberByID = `
	SELECT id, name, phone
	FROM members
	WHERE id = $1
`

func GetMemberByID(ctx *fiber.Ctx, ID int) (model.Member, bool, error) {
	var member model.Member

	db := postgresPkg.GetPgConn()

	if err := db.GetContext(ctx.Context(), &member, queryGetMemberByID, ID); err != nil {
		if err == sql.ErrNoRows {
			return member, false, nil
		}

		return member, false, err
	}

	return member, true, nil
}

const insertMember = `
	INSERT INTO members (name, phone)
	VALUES (:name, :phone)
`

func InsertMember(ctx *fiber.Ctx, member model.Member) error {
	db := postgresPkg.GetPgConn()

	_, err := db.NamedQueryContext(ctx.Context(), insertMember, member)
	return err
}

const updateMember = `
	UPDATE members SET
		name = :name,
	    phone = :phone,
		update_time = NOW()
	WHERE id = :id
`

func UpdateMember(ctx *fiber.Ctx, member model.Member) error {
	db := postgresPkg.GetPgConn()

	_, err := db.NamedQueryContext(ctx.Context(), updateMember, member)
	return err
}

const deleteMember = `
	DELETE FROM members
	WHERE id = $1
`

func DeleteMember(ctx *fiber.Ctx, ID int) error {
	db := postgresPkg.GetPgConn()

	_, err := db.ExecContext(ctx.Context(), deleteMember, ID)
	return err
}
