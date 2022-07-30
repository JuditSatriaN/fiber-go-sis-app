package services

import (
	"github.com/gofiber/fiber/v2"

	memberSvc "github.com/fiber-go-sis-app/internal/handler/services/members"
)

// BuildMemberRoutes : Service - service to handle member
func BuildMemberRoutes(service fiber.Router) {
	service.Get("/members", memberSvc.GetAllMemberHandler)
	service.Get("/dt_members", memberSvc.GetAllDTMemberHandler)
	service.Post("/member/insert", memberSvc.InsertMemberHandler)
	service.Post("/member/update", memberSvc.UpdateMemberHandler)
	service.Post("/member/delete", memberSvc.DeleteMemberHandler)
	service.Post("/member/upsert", memberSvc.UpsertMemberHandler)
}
