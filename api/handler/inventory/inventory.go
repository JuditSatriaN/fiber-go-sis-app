package inventory

import (
	"github.com/fiber-go-sis-app/internal/app/model"
	"github.com/gofiber/fiber/v2"

	inventoryUC "github.com/fiber-go-sis-app/internal/app/usecase/inventory"
	customPkg "github.com/fiber-go-sis-app/internal/pkg/custom"
)

// GetDTInventoryHandler : Get List Inventory for Datatable
func GetDTInventoryHandler(ctx *fiber.Ctx) error {
	page, limit, err := customPkg.BuildPageAndLimit(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	search := ctx.Query("search", "")

	result, err := inventoryUC.GetDTAllInventory(ctx, page, limit, search)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return customPkg.BuildDatatableRes(ctx, result.Total, result.Data)
}

func UpdateStockInventoryHandler(ctx *fiber.Ctx) error {
	var inventory model.Inventory

	if err := customPkg.ValidateRequest(ctx, &inventory); err != nil {
		return err
	}

	if err := inventoryUC.UpdateStockInventory(ctx, inventory); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.SendString("Data update stock berhasil diubah")
}
