package inventory

import (
	"github.com/gofiber/fiber/v2"

	inventoryHandler "github.com/fiber-go-sis-app/api/handler/inventory"
)

// BuildInventoryAPI : API to handle inventory product
func BuildInventoryAPI(api fiber.Router) {
	api.Get("/inventory", inventoryHandler.GetDTInventoryHandler)
	api.Post("/inventory/update_stock", inventoryHandler.UpdateStockInventoryHandler)
}
