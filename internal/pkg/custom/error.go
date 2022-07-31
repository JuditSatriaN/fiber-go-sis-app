package custom

import "github.com/gofiber/fiber/v2"

var CustomErrorHandler = func(c *fiber.Ctx, err error) error {
	// Default 500 status-code
	code := fiber.StatusInternalServerError

	if e, ok := err.(*fiber.Error); ok {
		// Override status code if fiber.Error type
		code = e.Code
	}

	// Set Content-Type: text/plain; charset=utf-8
	c.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)

	if code == fiber.StatusNotFound {
		return c.Status(code).Redirect("/404_not_found")
	}

	// Return status-code with error message
	return c.Status(code).SendString(err.Error())
}
