package sales

import (
	"github.com/fiber-go-sis-app/internal/app/model"
	"github.com/gofiber/fiber/v2"

	salesUC "github.com/fiber-go-sis-app/internal/app/usecase/sales"
	customPkg "github.com/fiber-go-sis-app/internal/pkg/custom"
)

func InsertSalesHandler(ctx *fiber.Ctx) error {
	var sales model.Sales

	if err := customPkg.ValidateRequest(ctx, &sales); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := salesUC.InsertSales(ctx, sales); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.SendString("Data sales berhasil disimpan")
}

func GetSalesHeadHandler(ctx *fiber.Ctx) error {
	page, limit, err := customPkg.BuildPageAndLimit(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	search := ctx.Query("search", "")

	result, err := salesUC.GetDTAllSalesHead(ctx, page, limit, search)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return customPkg.BuildDatatableRes(ctx, result.Total, result.Data)
}

func GetSalesDetailByInvoiceHandler(ctx *fiber.Ctx) error {
	invoice := ctx.Query("invoice", "")
	if invoice == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invoice tidak boleh kosong",
		})
	}

	result, err := salesUC.GetSalesDetailByInvoice(ctx, invoice)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return customPkg.BuildDatatableRes(ctx, result.Total, result.Data)
}
