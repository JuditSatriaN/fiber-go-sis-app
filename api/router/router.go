package router

import (
	"github.com/gofiber/fiber/v2"

	inventoryRoute "github.com/fiber-go-sis-app/api/router/inventory"
	loginRoute "github.com/fiber-go-sis-app/api/router/login"
	memberRoute "github.com/fiber-go-sis-app/api/router/member"
	productRoute "github.com/fiber-go-sis-app/api/router/product"
	salesRoute "github.com/fiber-go-sis-app/api/router/sales"
	unitRoute "github.com/fiber-go-sis-app/api/router/unit"
	userRoute "github.com/fiber-go-sis-app/api/router/user"
)

// BuildAPIRouter : Function to handle all API in this project
func BuildAPIRouter(app *fiber.App) {
	apiGroup := app.Group("/api")
	userRoute.BuildUserAPI(apiGroup)
	unitRoute.BuildUnitAPI(apiGroup)
	loginRoute.BuildLoginAPI(apiGroup)
	salesRoute.BuildSalesAPI(apiGroup)
	memberRoute.BuildMemberAPI(apiGroup)
	productRoute.BuildProductAPI(apiGroup)
	inventoryRoute.BuildInventoryAPI(apiGroup)
}
