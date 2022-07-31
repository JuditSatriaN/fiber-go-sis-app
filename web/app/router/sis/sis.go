package sis

import (
	"github.com/gofiber/fiber/v2"

	homeHandler "github.com/fiber-go-sis-app/web/app/handler/home"
	memberHandler "github.com/fiber-go-sis-app/web/app/handler/member"
	productHandler "github.com/fiber-go-sis-app/web/app/handler/product"
	userHandler "github.com/fiber-go-sis-app/web/app/handler/user"
)

func BuildSISRoutes(service fiber.Router) {
	service.Get("/", homeHandler.WebSISHomeHandler)
	service.Get("/user", userHandler.WebSISUserHandler)
	service.Get("/member", memberHandler.WebSISMemberHandler)
	service.Get("/product", productHandler.WebProductHandler)
}
