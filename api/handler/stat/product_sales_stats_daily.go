package stat

import (
	statUC "github.com/fiber-go-sis-app/internal/app/usecase/stat"
	customPkg "github.com/fiber-go-sis-app/internal/pkg/custom"
	"github.com/gofiber/fiber/v2"
)

func GetTop3ProductSalesDailyHandler(ctx *fiber.Ctx) error {
	results, err := statUC.GetTop3ProductSalesDaily(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return customPkg.BuildJSONRes(ctx, results)
}
