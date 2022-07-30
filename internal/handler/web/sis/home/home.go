package home

import (
	"github.com/gofiber/fiber/v2"

	constantsEntity "github.com/fiber-go-sis-app/internal/entity/constants"
)

func WebSISHomeHandler(ctx *fiber.Ctx) error {

	return ctx.Render("templates/sis/pages/home/home", constantsEntity.WebData{
		BaseURL:      constantsEntity.BaseURL,
		StaticUrl:    constantsEntity.StaticUrl,
		LinkPageList: constantsEntity.LinkPageList,
		CurrentURL:   constantsEntity.WebSISHomeURL,
		Title:        constantsEntity.WebSISHomeTitle,
	})
}
