package invoice

import (
	invoiceUC "github.com/fiber-go-sis-app/internal/app/usecase/invoice"
	customPkg "github.com/fiber-go-sis-app/internal/pkg/custom"
	"github.com/gofiber/fiber/v2"
)

func GetInvoiceHandler(ctx *fiber.Ctx) error {
	invoiceData, err := invoiceUC.GetInvoice(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return customPkg.BuildJSONRes(ctx, invoiceData)
}

func UpdateInvoiceHandler(ctx *fiber.Ctx) error {
	if err := invoiceUC.UpdateInvoice(ctx); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.SendString("Data invoice berhasil diupdate")
}
