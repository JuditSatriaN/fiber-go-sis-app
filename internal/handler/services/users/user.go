package users

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"

	userEntity "github.com/fiber-go-sis-app/internal/entity/users"

	customPkg "github.com/fiber-go-sis-app/utils/pkg/custom"

	userUC "github.com/fiber-go-sis-app/internal/usecase/services/users"
)

func GetAllUserHandler(ctx *fiber.Ctx) error {
	users, err := userUC.GetAllUser(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return customPkg.BuildJSONRes(ctx, users)
}

func GetAllDTUserHandler(ctx *fiber.Ctx) error {
	users, err := userUC.GetAllUser(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return customPkg.BuildDatatableRes(ctx, int64(len(users)), users)
}

func InsertUserHandler(ctx *fiber.Ctx) error {
	var user userEntity.User

	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := validator.New().Struct(&user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if user.Password != "" {
		user.Password, _ = customPkg.HashPassword(user.Password)
	}

	if err := userUC.InsertUser(ctx, user); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.SendString("Data user berhasil disimpan")
}

func UpdateUserHandler(ctx *fiber.Ctx) error {
	var user userEntity.User
	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := validator.New().Struct(&user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := userUC.UpdateUser(ctx, user); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.SendString("Data user berhasil diubah")
}

func DeleteUserHandler(ctx *fiber.Ctx) error {
	var user userEntity.User

	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := validator.New().Struct(&user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := userUC.DeleteUser(ctx, user.UserID); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.SendString("Data user berhasil dihapus")
}

func UpsertUserHandler(ctx *fiber.Ctx) error {
	var user userEntity.User

	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := validator.New().Struct(&user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := userUC.UpsertUser(ctx, user); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return customPkg.BuildJSONRes(ctx, user)
}
