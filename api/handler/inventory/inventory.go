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

// SearchInventoryHandler : Search all inventory by param
func SearchInventoryHandler(ctx *fiber.Ctx) error {
	search := ctx.Query("search", "")

	results, err := inventoryUC.SearchInventoryByParam(ctx, search)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return customPkg.BuildJSONRes(ctx, results)
}

func InsertInventoryHandler(ctx *fiber.Ctx) error {
	var inventory model.Inventory

	if err := customPkg.ValidateRequest(ctx, &inventory); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := inventoryUC.InsertInventory(ctx, inventory); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.SendString("Data inventory berhasil disimpan")
}

func UpdateInventoryHandler(ctx *fiber.Ctx) error {
	var inventory model.Inventory

	if err := customPkg.ValidateRequest(ctx, &inventory); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := inventoryUC.UpdateInventory(ctx, inventory); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.SendString("Data inventory berhasil diubah")
}

func DeleteInventoryHandler(ctx *fiber.Ctx) error {
	var inventory model.Inventory

	if err := customPkg.ValidateRequest(ctx, &inventory); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := inventoryUC.DeleteInventory(ctx, inventory.ID); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return ctx.SendString("Data inventory berhasil dihapus")
}

func UpsertInventoryHandler(ctx *fiber.Ctx) error {
	var inventory model.Inventory

	if err := customPkg.ValidateRequest(ctx, &inventory); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := inventoryUC.UpsertInventory(ctx, inventory); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return customPkg.BuildJSONRes(ctx, inventory)
}

func UpdateStockInventoryHandler(ctx *fiber.Ctx) error {
	var inventory model.Inventory

	if err := customPkg.ValidateRequest(ctx, &inventory); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := inventoryUC.UpdateStockInventory(ctx, inventory); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.SendString("Data update stock berhasil diubah")
}
