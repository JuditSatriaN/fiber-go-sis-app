package system_conf

import (
	"fmt"

	"github.com/fiber-go-sis-app/internal/app/constant"
	"github.com/gofiber/fiber/v2"
)

func WebSISSystemConfHandler(ctx *fiber.Ctx) error {

	return ctx.Render(fmt.Sprintf("%s/master/system_conf", constant.PagesSISWebDirectory), constant.WebData{
		BaseURL:      constant.BaseURL,
		StaticUrl:    constant.StaticUrl,
		LinkPageList: constant.LinkPageList,
		CurrentURL:   constant.WebSISSystemConfURL,
		Title:        constant.WebSISSystemConfTitle,
	})
}
