package unit

import (
	"github.com/gofiber/fiber/v2"

	unitHandler "github.com/fiber-go-sis-app/api/handler/unit"
)

// BuildUnitAPI : API to handle unit
func BuildUnitAPI(api fiber.Router) {
	api.Get("/units", unitHandler.GetAllUnitHandler)
	api.Get("/dt_units", unitHandler.GetAllDTUnitHandler)
	api.Post("/unit/insert", unitHandler.InsertUnitHandler)
	api.Post("/unit/update", unitHandler.UpdateUnitHandler)
	api.Post("/unit/delete", unitHandler.DeleteUnitHandler)
	api.Post("/unit/upsert", unitHandler.UpsertUnitHandler)
}
