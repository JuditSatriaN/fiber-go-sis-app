package system_conf

import (
	"github.com/fiber-go-sis-app/internal/app/model"
	"github.com/gofiber/fiber/v2"

	systemConfUC "github.com/fiber-go-sis-app/internal/app/usecase/system_conf"
	customPkg "github.com/fiber-go-sis-app/internal/pkg/custom"
)

func GetAllSystemConfHandler(ctx *fiber.Ctx) error {
	data, err := systemConfUC.GetAllSystemConf(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return customPkg.BuildJSONRes(ctx, data)
}

func GetKeyVoidHandler(ctx *fiber.Ctx) error {
	data, err := systemConfUC.GetKeyVoid(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return customPkg.BuildJSONRes(ctx, data)
}

func GetSystemConfByIDHandler(ctx *fiber.Ctx) error {
	var param model.SystemConf

	if err := customPkg.ValidateRequest(ctx, &param); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	data, err := systemConfUC.GetSystemConfByID(ctx, param.ID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return customPkg.BuildJSONRes(ctx, data)
}

func GetAllDTSystemConfHandler(ctx *fiber.Ctx) error {
	data, err := systemConfUC.GetAllSystemConf(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return customPkg.BuildDatatableRes(ctx, int64(len(data)), data)
}

func InsertSystemConfHandler(ctx *fiber.Ctx) error {
	var data model.SystemConf

	if err := customPkg.ValidateRequest(ctx, &data); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := systemConfUC.InsertSystemConf(ctx, data); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.SendString("Data system config berhasil disimpan")
}

func UpdateSystemConfHandler(ctx *fiber.Ctx) error {
	var data model.SystemConf

	if err := customPkg.ValidateRequest(ctx, &data); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := systemConfUC.UpdateSystemConf(ctx, data); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.SendString("Data system config berhasil diubah")
}

func DeleteSystemConfHandler(ctx *fiber.Ctx) error {
	var data model.SystemConf

	if err := customPkg.ValidateRequest(ctx, &data); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := systemConfUC.DeleteSystemConf(ctx, data.ID); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return ctx.SendString("Data System Config berhasil dihapus")
}

func UpsertSystemConfHandler(ctx *fiber.Ctx) error {
	var data model.SystemConf

	if err := customPkg.ValidateRequest(ctx, &data); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := systemConfUC.UpsertSystemConf(ctx, data); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return customPkg.BuildJSONRes(ctx, data)
}
