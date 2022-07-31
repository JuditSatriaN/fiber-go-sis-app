package login

import (
	"github.com/fiber-go-sis-app/internal/app/model"
	"github.com/gofiber/fiber/v2"

	loginUC "github.com/fiber-go-sis-app/internal/app/usecase/login"
	customPkg "github.com/fiber-go-sis-app/internal/pkg/custom"
)

func ProcessLoginHandler(ctx *fiber.Ctx) error {
	var loginRequest model.LoginRequest

	if err := customPkg.ValidateRequest(ctx, &loginRequest); err != nil {
		return err
	}

	data, err := loginUC.ProcessLoginForm(ctx, loginRequest)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return customPkg.BuildJSONRes(ctx, data)
}
