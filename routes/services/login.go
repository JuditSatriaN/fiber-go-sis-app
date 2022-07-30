package services

import (
	"github.com/gofiber/fiber/v2"

	formsSvc "github.com/fiber-go-sis-app/internal/handler/services/forms"
)

// BuildLoginRoutes : Service - service to handle login
func BuildLoginRoutes(service fiber.Router) {
	service.Post("/login", formsSvc.LoginHandler)
}
