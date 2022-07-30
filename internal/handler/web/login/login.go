package login

import (
	"github.com/gofiber/fiber/v2"

	constantsEntity "github.com/fiber-go-sis-app/internal/entity/constants"
)

func WebLoginHandler(ctx *fiber.Ctx) error {
	// cek if user is already login
	if ctx.Cookies(constantsEntity.JWTAccessCookiesKey, "") != "" {
		return ctx.Redirect(constantsEntity.BaseSISURL)
	}

	return ctx.Render("templates/login/index", constantsEntity.WebData{
		BaseURL:      constantsEntity.BaseURL,
		StaticUrl:    constantsEntity.StaticUrl,
		CurrentURL:   constantsEntity.WebLoginURL,
		LinkPageList: constantsEntity.LinkPageList,
		Title:        constantsEntity.WebLoginTitle,
	})
}
