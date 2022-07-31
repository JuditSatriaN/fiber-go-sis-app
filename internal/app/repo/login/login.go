package login

import (
	"database/sql"

	"github.com/fiber-go-sis-app/internal/app/constant"
	"github.com/fiber-go-sis-app/internal/app/model"
	"github.com/gofiber/fiber/v2"

	postgresPkg "github.com/fiber-go-sis-app/internal/pkg/databases/postgres"
)

const queryLoginForm = `
	SELECT user_name, password
	FROM users
	WHERE user_name = $1
	LIMIT 1
`

func ProcessLoginForm(ctx *fiber.Ctx, userName string) (model.LoginRequest, error) {
	var user model.LoginRequest
	db := postgresPkg.GetPgConn()

	if err := db.GetContext(ctx.Context(), &user, queryLoginForm, userName); err != nil {
		if err == sql.ErrNoRows {
			return user, constant.ErrUserNotFound
		}
		return user, err
	}
	return user, nil
}
