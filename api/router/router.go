package router

import (
	"github.com/gofiber/fiber/v2"

	inventoryRoute "github.com/fiber-go-sis-app/api/router/inventory"
	loginRoute "github.com/fiber-go-sis-app/api/router/login"
	memberRoute "github.com/fiber-go-sis-app/api/router/member"
	productRoute "github.com/fiber-go-sis-app/api/router/product"
	userRoute "github.com/fiber-go-sis-app/api/router/user"
)

// BuildAPIRouter : Function to handle all API in this project
func BuildAPIRouter(app *fiber.App) {
	apiGroup := app.Group("/api")
	userRoute.BuildUserAPI(apiGroup)
	loginRoute.BuildLoginAPI(apiGroup)
	memberRoute.BuildMemberAPI(apiGroup)
	productRoute.BuildProductAPI(apiGroup)
	inventoryRoute.BuildInventoryAPI(apiGroup)
}
