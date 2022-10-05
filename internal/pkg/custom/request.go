package custom

import (
	"fmt"
	"strconv"

	"github.com/fiber-go-sis-app/internal/app/constant"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// ValidateRequest global function to validate request
func ValidateRequest(ctx *fiber.Ctx, dest interface{}) error {
	if err := ctx.BodyParser(dest); err != nil {
		return err
	}

	if err := validator.New().Struct(dest); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			if err.Tag() == constant.ErrFieldStartsWith {
				return ConvertErrorStartswith(err.Field(), err.Param())
			}
		}
		return err
	}

	return nil
}

// BuildPageAndLimit global function to build page and limit
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

// BuildOffset global function to build offset test reset 1 branch 1
func BuildOffset(page int, limit int) int {
	offset := (page - 1) * limit
	return offset
}
