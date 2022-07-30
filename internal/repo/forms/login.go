package forms

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"

	postgresPkg "github.com/fiber-go-sis-app/utils/pkg/databases/postgres"

	constantsEntity "github.com/fiber-go-sis-app/internal/entity/constants"
	formsEntity "github.com/fiber-go-sis-app/internal/entity/forms"
)

const queryLoginForm = `
	SELECT user_name, password
	FROM users
	WHERE user_name = $1
	LIMIT 1
`

func LoginFormRepo(ctx *fiber.Ctx, userName string) (formsEntity.LoginRequest, error) {
	var user formsEntity.LoginRequest
	db := postgresPkg.GetPgConn()

	if err := db.GetContext(ctx.Context(), &user, queryLoginForm, userName); err != nil {
		if err == sql.ErrNoRows {
			return user, constantsEntity.ErrUserNotFound
		}
		return user, err
	}
	return user, nil
}
