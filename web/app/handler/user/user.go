package user

import (
	"fmt"

	"github.com/fiber-go-sis-app/internal/app/constant"
	"github.com/gofiber/fiber/v2"
)

func WebSISUserHandler(ctx *fiber.Ctx) error {

	return ctx.Render(fmt.Sprintf("%s/master/user", constant.PagesSISWebDirectory), constant.WebData{
		BaseURL:      constant.BaseURL,
		StaticUrl:    constant.StaticUrl,
		LinkPageList: constant.LinkPageList,
		CurrentURL:   constant.WebSISUserURL,
		Title:        constant.WebSISUserTitle,
	})
}
