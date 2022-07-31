package member

import (
	"fmt"

	"github.com/fiber-go-sis-app/internal/app/constant"
	"github.com/gofiber/fiber/v2"
)

func WebSISMemberHandler(ctx *fiber.Ctx) error {

	return ctx.Render(fmt.Sprintf("%s/master/member", constant.PagesSISWebDirectory), constant.WebData{
		BaseURL:      constant.BaseURL,
		StaticUrl:    constant.StaticUrl,
		LinkPageList: constant.LinkPageList,
		CurrentURL:   constant.WebSISMemberURL,
		Title:        constant.WebSISMemberTitle,
	})
}
