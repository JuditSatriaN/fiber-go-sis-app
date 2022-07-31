package error

import (
	"github.com/gofiber/fiber/v2"

	errorHandler "github.com/fiber-go-sis-app/web/app/handler/error"
)

func BuildError404NotFound(service fiber.Router) {
	service.Get("/404_not_found", errorHandler.Web404NotFoundHandler)
}
