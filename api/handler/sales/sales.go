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
