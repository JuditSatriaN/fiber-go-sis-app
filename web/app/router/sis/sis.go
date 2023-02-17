package sis

import (
	"github.com/gofiber/fiber/v2"

	// Master Handler
	homeHandler "github.com/fiber-go-sis-app/web/app/handler/home"
	inventoryHandler "github.com/fiber-go-sis-app/web/app/handler/inventory"
	memberHandler "github.com/fiber-go-sis-app/web/app/handler/member"
	productHandler "github.com/fiber-go-sis-app/web/app/handler/product"
	systemConfHandler "github.com/fiber-go-sis-app/web/app/handler/system_conf"
	unitHandler "github.com/fiber-go-sis-app/web/app/handler/unit"
	userHandler "github.com/fiber-go-sis-app/web/app/handler/user"

	// Proses Handler
	updateStockHandler "github.com/fiber-go-sis-app/web/app/handler/update_stock"
)

func BuildSISRoutes(service fiber.Router) {
	service.Get("/", homeHandler.WebSISHomeHandler)
	service.Get("/user", userHandler.WebSISUserHandler)
	service.Get("/unit", unitHandler.WebSISUnitHandler)
	service.Get("/member", memberHandler.WebSISMemberHandler)
	service.Get("/product", productHandler.WebProductHandler)
	service.Get("/inventory", inventoryHandler.WebSISInventoryHandler)
	service.Get("/system_conf", systemConfHandler.WebSISSystemConfHandler)

	service.Get("/update_stock", updateStockHandler.WebSISUpdateStockHandler)
}
