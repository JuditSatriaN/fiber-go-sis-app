package sis

import (
	"github.com/gofiber/fiber/v2"

	homeHandler "github.com/fiber-go-sis-app/web/app/handler/home"
	inventoryHandler "github.com/fiber-go-sis-app/web/app/handler/inventory"
	memberHandler "github.com/fiber-go-sis-app/web/app/handler/member"
	productHandler "github.com/fiber-go-sis-app/web/app/handler/product"
	unitHandler "github.com/fiber-go-sis-app/web/app/handler/unit"
	userHandler "github.com/fiber-go-sis-app/web/app/handler/user"

	updateStockHandler "github.com/fiber-go-sis-app/web/app/handler/update_stock"
)

func BuildSISRoutes(service fiber.Router) {
	service.Get("/", homeHandler.WebSISHomeHandler)
	service.Get("/user", userHandler.WebSISUserHandler)
	service.Get("/unit", unitHandler.WebSISUnitHandler)
	service.Get("/member", memberHandler.WebSISMemberHandler)
	service.Get("/product", productHandler.WebProductHandler)
	service.Get("/inventory", inventoryHandler.WebSISInventoryHandler)

	service.Get("/update_stock", updateStockHandler.WebSISUpdateStockHandler)
}
