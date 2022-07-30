package members

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"

	memberEntity "github.com/fiber-go-sis-app/internal/entity/members"

	memberUC "github.com/fiber-go-sis-app/internal/usecase/services/members"

	customPkg "github.com/fiber-go-sis-app/utils/pkg/custom"
)

func GetAllMemberHandler(ctx *fiber.Ctx) error {
	members, err := memberUC.GetAllMember(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return customPkg.BuildJSONRes(ctx, members)
}

func GetAllDTMemberHandler(ctx *fiber.Ctx) error {
	members, err := memberUC.GetAllMember(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return customPkg.BuildDatatableRes(ctx, int64(len(members)), members)
}

func InsertMemberHandler(ctx *fiber.Ctx) error {
	var member memberEntity.Member

	if err := ctx.BodyParser(&member); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := validator.New().Struct(&member); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := memberUC.InsertMember(ctx, member); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.SendString("Data member berhasil disimpan")
}

func UpdateMemberHandler(ctx *fiber.Ctx) error {
	var member memberEntity.Member

	if err := ctx.BodyParser(&member); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := validator.New().Struct(&member); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := memberUC.UpdateMember(ctx, member); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.SendString("Data member berhasil diubah")
}

func DeleteMemberHandler(ctx *fiber.Ctx) error {
	var member memberEntity.Member

	if err := ctx.BodyParser(&member); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := validator.New().Struct(&member); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := memberUC.DeleteMember(ctx, member.ID); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return ctx.SendString("Data Member berhasil dihapus")
}

func UpsertMemberHandler(ctx *fiber.Ctx) error {
	var member memberEntity.Member

	if err := ctx.BodyParser(&member); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := validator.New().Struct(&member); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := memberUC.UpsertMember(ctx, member); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return customPkg.BuildJSONRes(ctx, member)
}
