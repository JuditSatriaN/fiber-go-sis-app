package member

import (
	"github.com/fiber-go-sis-app/internal/app/model"
	"github.com/gofiber/fiber/v2"

	memberUC "github.com/fiber-go-sis-app/internal/app/usecase/member"
	customPkg "github.com/fiber-go-sis-app/internal/pkg/custom"
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
	var member model.Member

	if err := customPkg.ValidateRequest(ctx, &member); err != nil {
		return err
	}

	if err := memberUC.InsertMember(ctx, member); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.SendString("Data member berhasil disimpan")
}

func UpdateMemberHandler(ctx *fiber.Ctx) error {
	var member model.Member

	if err := customPkg.ValidateRequest(ctx, &member); err != nil {
		return err
	}

	if err := memberUC.UpdateMember(ctx, member); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.SendString("Data member berhasil diubah")
}

func DeleteMemberHandler(ctx *fiber.Ctx) error {
	var member model.Member

	if err := customPkg.ValidateRequest(ctx, &member); err != nil {
		return err
	}

	if err := memberUC.DeleteMember(ctx, member.ID); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return ctx.SendString("Data Member berhasil dihapus")
}

func UpsertMemberHandler(ctx *fiber.Ctx) error {
	var member model.Member

	if err := customPkg.ValidateRequest(ctx, &member); err != nil {
		return err
	}

	if err := memberUC.UpsertMember(ctx, member); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return customPkg.BuildJSONRes(ctx, member)
}
