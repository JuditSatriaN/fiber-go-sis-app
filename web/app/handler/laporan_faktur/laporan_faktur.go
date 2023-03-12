package laporan_faktur

import (
	"fmt"

	"github.com/fiber-go-sis-app/internal/app/constant"
	"github.com/gofiber/fiber/v2"
)

func WebSISLaporanFakturHandler(ctx *fiber.Ctx) error {

	return ctx.Render(fmt.Sprintf("%s/laporan/laporan_faktur", constant.PagesSISWebDirectory), constant.WebData{
		BaseURL:      constant.BaseURL,
		StaticUrl:    constant.StaticUrl,
		LinkPageList: constant.LinkPageList,
		CurrentURL:   constant.WebSISLaporanFakturURL,
		Title:        constant.WebSISLaporanFakturTitle,
	})
}
