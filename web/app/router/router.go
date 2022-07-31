package router

import (
	"github.com/fiber-go-sis-app/internal/app/constant"
	"github.com/gofiber/fiber/v2"

	jwtMiddleware "github.com/fiber-go-sis-app/internal/pkg/middleware/jwt"

	errorRouter "github.com/fiber-go-sis-app/web/app/router/error"
	loginRouter "github.com/fiber-go-sis-app/web/app/router/login"
	sisRouter "github.com/fiber-go-sis-app/web/app/router/sis"
)

// BuildWebRouter : Function to handle all web router in this project
func BuildWebRouter(app *fiber.App) {
	loginRouter.BuildLoginRoutes(app)
	errorRouter.BuildError404NotFound(app)

	sisGroup := app.Group("/sis")
	sisGroup.Use(jwtMiddleware.AccessTokenMiddleware(constant.WebSource))
	sisGroup.Use(jwtMiddleware.RefreshTokenMiddleware(constant.WebSource))
	sisRouter.BuildSISRoutes(sisGroup)
}
