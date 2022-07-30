package errors

import (
	"github.com/gofiber/fiber/v2"

	constantsEntity "github.com/fiber-go-sis-app/internal/entity/constants"
)

func Web404NotFoundHandler(ctx *fiber.Ctx) error {

	return ctx.Render("templates/errors/404_not_found", constantsEntity.WebData{
		BaseURL:   constantsEntity.BaseURL,
		StaticUrl: constantsEntity.StaticUrl,
		Title:     constantsEntity.Web404NotFoundTitle,
	})
}
