package void

import (
	"fmt"

	"github.com/fiber-go-sis-app/internal/app/constant"
	"github.com/gofiber/fiber/v2"
)

func WebSISVoidTransactionHandler(ctx *fiber.Ctx) error {

	return ctx.Render(fmt.Sprintf("%s/proses/void_transaction", constant.PagesSISWebDirectory), constant.WebData{
		BaseURL:      constant.BaseURL,
		StaticUrl:    constant.StaticUrl,
		LinkPageList: constant.LinkPageList,
		CurrentURL:   constant.WebSISVoidTransactionURL,
		Title:        constant.WebSISVoidTransactionTitle,
	})
}
