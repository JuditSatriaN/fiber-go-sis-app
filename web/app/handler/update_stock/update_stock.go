package update_stock

import (
	"fmt"

	"github.com/fiber-go-sis-app/internal/app/constant"
	"github.com/gofiber/fiber/v2"
)

func WebSISUpdateStockHandler(ctx *fiber.Ctx) error {

	return ctx.Render(fmt.Sprintf("%s/proses/update_stock", constant.PagesSISWebDirectory), constant.WebData{
		BaseURL:      constant.BaseURL,
		StaticUrl:    constant.StaticUrl,
		LinkPageList: constant.LinkPageList,
		CurrentURL:   constant.WebSISUpdateStockURL,
		Title:        constant.WebSISUpdateStockTitle,
	})
}
