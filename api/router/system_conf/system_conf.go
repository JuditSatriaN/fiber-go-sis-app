package system_conf

import (
	"github.com/gofiber/fiber/v2"

	systemConfHandler "github.com/fiber-go-sis-app/api/handler/system_conf"
)

// BuildSystemConfAPI : API to handle system config
func BuildSystemConfAPI(api fiber.Router) {
	api.Get("/systems_conf", systemConfHandler.GetAllSystemConfHandler)
	api.Get("/get_key_void", systemConfHandler.GetKeyVoidHandler)
	api.Get("/dt_system_conf", systemConfHandler.GetAllDTSystemConfHandler)
	api.Post("/system_conf/insert", systemConfHandler.InsertSystemConfHandler)
	api.Post("/system_conf/update", systemConfHandler.UpdateSystemConfHandler)
	api.Post("/system_conf/delete", systemConfHandler.DeleteSystemConfHandler)
	api.Post("/system_conf/upsert", systemConfHandler.UpsertSystemConfHandler)
}
