package product

import (
	"fmt"
	"github.com/fiber-go-sis-app/internal/app/model"
	"github.com/gofiber/fiber/v2"

	inventoryRepo "github.com/fiber-go-sis-app/internal/app/repo/inventory"
	storeRepo "github.com/fiber-go-sis-app/internal/app/repo/store"
	customPkg "github.com/fiber-go-sis-app/internal/pkg/custom"
)

// GetDTAllInventory : Get List Of Inventory for Datatable
func GetDTAllInventory(ctx *fiber.Ctx, page int, limit int, search string) (model.ListInventoryDataResponse, error) {
	offset := customPkg.BuildOffset(page, limit)

	inventories, err := inventoryRepo.GetALlInventory(ctx, search, limit, offset)
	if err != nil {
		return model.ListInventoryDataResponse{}, err
	}

	totalProduct, err := storeRepo.GetTotalProduct(ctx, model.DefaultStoreID)
	if err != nil {
		return model.ListInventoryDataResponse{}, err
	}

	return model.ListInventoryDataResponse{
		Total: totalProduct,
		Data:  inventories,
	}, nil
}

func GetInventoryByID(ctx *fiber.Ctx, ID int) (model.Inventory, error) {
	inventory, found, err := inventoryRepo.GetInventoryByID(ctx, ID)
	if err != nil {
		return model.Inventory{}, err
	}

	if !found {
		return model.Inventory{}, fmt.Errorf("inventory dengan id : %d tidak ditemukan", ID)
	}

	return inventory, nil
}

func UpdateStockInventory(ctx *fiber.Ctx, inventory model.Inventory) error {
	if _, err := GetInventoryByID(ctx, inventory.ID); err != nil {
		return err
	}

	return inventoryRepo.UpdateStockInventory(ctx, inventory)
}
