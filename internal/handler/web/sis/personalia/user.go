package personalia

import (
	"github.com/gofiber/fiber/v2"

	constantsEntity "github.com/fiber-go-sis-app/internal/entity/constants"
)

func WebSISUserHandler(ctx *fiber.Ctx) error {

	return ctx.Render("templates/sis/pages/master/user", constantsEntity.WebData{
		BaseURL:      constantsEntity.BaseURL,
		StaticUrl:    constantsEntity.StaticUrl,
		LinkPageList: constantsEntity.LinkPageList,
		CurrentURL:   constantsEntity.WebSISUserURL,
		Title:        constantsEntity.WebSISUserTitle,
	})
}
