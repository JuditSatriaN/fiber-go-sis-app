package home

import (
	"fmt"

	"github.com/fiber-go-sis-app/internal/app/constant"
	"github.com/gofiber/fiber/v2"
)

func WebSISHomeHandler(ctx *fiber.Ctx) error {

	return ctx.Render(fmt.Sprintf("%s/home/home", constant.PagesSISWebDirectory), constant.WebData{
		BaseURL:      constant.BaseURL,
		StaticUrl:    constant.StaticUrl,
		LinkPageList: constant.LinkPageList,
		CurrentURL:   constant.WebSISHomeURL,
		Title:        constant.WebSISHomeTitle,
	})
}
