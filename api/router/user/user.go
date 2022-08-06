package user

import (
	"github.com/gofiber/fiber/v2"

	userHandler "github.com/fiber-go-sis-app/api/handler/user"
)

// BuildUserAPI : API to handle user
func BuildUserAPI(api fiber.Router) {
	api.Get("/users", userHandler.GetAllUserHandler)
	api.Get("/dt_users", userHandler.GetAllDTUserHandler)
	api.Post("/user/insert", userHandler.InsertUserHandler)
	api.Post("/user/update", userHandler.UpdateUserHandler)
	api.Post("/user/delete", userHandler.DeleteUserHandler)
	api.Post("/user/upsert", userHandler.UpsertUserHandler)
}
