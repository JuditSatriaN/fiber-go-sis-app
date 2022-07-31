package custom

import (
	"fmt"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func ValidateRequest(ctx *fiber.Ctx, dest interface{}) error {
	if err := ctx.BodyParser(dest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := validator.New().Struct(dest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return nil
}

func BuildPageAndLimit(ctx *fiber.Ctx) (int, int, error) {
	page, err := strconv.Atoi(ctx.Query("page", "1"))
	if err != nil {
		return 0, 0, fmt.Errorf("page must be number")
	}

	limit, err := strconv.Atoi(ctx.Query("limit", "10"))
	if err != nil {
		return 0, 0, fmt.Errorf("limit must be number")
	}

	return page, limit, nil
}

func BuildOffset(page int, limit int) int {
	offset := (page - 1) * limit
	return offset
}
