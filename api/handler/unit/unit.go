package unit

import (
	"github.com/fiber-go-sis-app/internal/app/model"
	unitUC "github.com/fiber-go-sis-app/internal/app/usecase/unit"
	customPkg "github.com/fiber-go-sis-app/internal/pkg/custom"
	"github.com/gofiber/fiber/v2"
)

func GetAllUnitHandler(ctx *fiber.Ctx) error {
	units, err := unitUC.GetAllUnit(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return customPkg.BuildJSONRes(ctx, units)
}

func GetAllDTUnitHandler(ctx *fiber.Ctx) error {
	units, err := unitUC.GetAllUnit(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return customPkg.BuildDatatableRes(ctx, int64(len(units)), units)
}

func InsertUnitHandler(ctx *fiber.Ctx) error {
	var unit model.Unit

	if err := customPkg.ValidateRequest(ctx, &unit); err != nil {
		return err
	}

	if err := unitUC.InsertUnit(ctx, unit); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.SendString("Data unit berhasil disimpan")
}

func UpdateUnitHandler(ctx *fiber.Ctx) error {
	var unit model.Unit

	if err := customPkg.ValidateRequest(ctx, &unit); err != nil {
		return err
	}

	if err := unitUC.UpdateUnit(ctx, unit); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.SendString("Data unit berhasil diubah")
}

func DeleteUnitHandler(ctx *fiber.Ctx) error {
	var unit model.Unit

	if err := customPkg.ValidateRequest(ctx, &unit); err != nil {
		return err
	}

	if err := unitUC.DeleteUnit(ctx, unit.ID); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return ctx.SendString("Data unit berhasil dihapus")
}

func UpsertUnitHandler(ctx *fiber.Ctx) error {
	var unit model.Unit

	if err := customPkg.ValidateRequest(ctx, &unit); err != nil {
		return err
	}

	if err := unitUC.UpsertUnit(ctx, unit); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return customPkg.BuildJSONRes(ctx, unit)
}
