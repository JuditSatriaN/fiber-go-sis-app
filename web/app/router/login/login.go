package login

import (
	"github.com/gofiber/fiber/v2"

	loginHandler "github.com/fiber-go-sis-app/web/app/handler/login"
)

func BuildLoginRoutes(service fiber.Router) {
	service.Get("/", loginHandler.WebLoginHandler)
}
