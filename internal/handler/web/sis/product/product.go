package product

import (
	"github.com/gofiber/fiber/v2"

	constantsEntity "github.com/fiber-go-sis-app/internal/entity/constants"
)

func WebProductHandler(ctx *fiber.Ctx) error {

	return ctx.Render("templates/sis/pages/master/product", constantsEntity.WebData{
		BaseURL:      constantsEntity.BaseURL,
		StaticUrl:    constantsEntity.StaticUrl,
		LinkPageList: constantsEntity.LinkPageList,
		CurrentURL:   constantsEntity.WebSISProductURL,
		Title:        constantsEntity.WebSISProductTitle,
	})
}
