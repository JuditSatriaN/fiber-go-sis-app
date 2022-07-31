package error

import (
	"fmt"
	"github.com/fiber-go-sis-app/internal/app/constant"
	"github.com/gofiber/fiber/v2"
)

func Web404NotFoundHandler(ctx *fiber.Ctx) error {

	return ctx.Render(fmt.Sprintf("%s/error/404_not_found", constant.PagesWebDirectory), constant.WebData{
		BaseURL:   constant.BaseURL,
		StaticUrl: constant.StaticUrl,
		Title:     constant.Web404NotFoundTitle,
	})
}
