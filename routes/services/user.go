package services

import (
	"github.com/gofiber/fiber/v2"

	userSvc "github.com/fiber-go-sis-app/internal/handler/services/users"
)

func BuildUserRoutes(service fiber.Router) {
	service.Get("/users", userSvc.GetAllUserHandler)
	service.Get("/dt_users", userSvc.GetAllDTUserHandler)
	service.Post("/user/insert", userSvc.InsertUserHandler)
	service.Post("/user/update", userSvc.UpdateUserHandler)
	service.Post("/user/delete", userSvc.DeleteUserHandler)
	service.Post("/user/upsert", userSvc.UpsertUserHandler)
}
