package custom

import (
	"fmt"
	"github.com/fiber-go-sis-app/internal/app/constant"
	"github.com/gofiber/fiber/v2"
)

//ErrorHandler function custom package to error handler
var ErrorHandler = func(ctx *fiber.Ctx, err error) error {
	// Default 500 status-code
	code := fiber.StatusInternalServerError

	if e, ok := err.(*fiber.Error); ok {
		// Override status code if fiber.Error type
		code = e.Code
	}

	// Set Content-Type: text/plain; charset=utf-8
	ctx.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)

	if code == fiber.StatusNotFound {
		return ctx.Status(code).Redirect(constant.Web404NotFoundURL)
	}

	// Return status-code with error message
	return ctx.Status(code).SendString(err.Error())
}

func ConvertErrorStartswith(field string, param string) error {
	return fmt.Errorf("Field %s harus dimulai dengan kata %s ", field, param)
}
