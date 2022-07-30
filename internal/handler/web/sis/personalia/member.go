package personalia

import (
	"github.com/gofiber/fiber/v2"

	constantsEntity "github.com/fiber-go-sis-app/internal/entity/constants"
)

func WebSISMemberHandler(ctx *fiber.Ctx) error {

	return ctx.Render("templates/sis/pages/master/member", constantsEntity.WebData{
		BaseURL:      constantsEntity.BaseURL,
		StaticUrl:    constantsEntity.StaticUrl,
		LinkPageList: constantsEntity.LinkPageList,
		CurrentURL:   constantsEntity.WebSISMemberURL,
		Title:        constantsEntity.WebSISMemberTitle,
	})
}
