package login

import (
	"fmt"
	"github.com/fiber-go-sis-app/internal/app/constant"
	"github.com/gofiber/fiber/v2"
)

func WebLoginHandler(ctx *fiber.Ctx) error {
	// Check if user is already login
	if ctx.Cookies(constant.JWTAccessCookiesKey, "") != "" {
		return ctx.Redirect(constant.BaseSISURL)
	}

	return ctx.Render(fmt.Sprintf("%s/login/index", constant.PagesWebDirectory), constant.WebData{
		BaseURL:      constant.BaseURL,
		StaticUrl:    constant.StaticUrl,
		CurrentURL:   constant.WebLoginURL,
		LinkPageList: constant.LinkPageList,
		Title:        constant.WebLoginTitle,
	})
}
