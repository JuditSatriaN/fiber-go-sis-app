package custom

import (
	"github.com/fiber-go-sis-app/internal/app/constant"
	"github.com/gofiber/fiber/v2"
)

//SetLoginCookie custom package to set login cookie
func SetLoginCookie(ctx *fiber.Ctx, accessToken string, refreshToken string) {
	ctx.Cookie(&fiber.Cookie{
		Name:  constant.JWTAccessCookiesKey,
		Value: accessToken,
	})

	ctx.Cookie(&fiber.Cookie{
		Name:  constant.JWTRefreshCookiesKey,
		Value: refreshToken,
	})
}

//ClearLoginCookie custom package to clear login cookie
func ClearLoginCookie(ctx *fiber.Ctx) {
	ctx.ClearCookie(constant.JWTAccessCookiesKey)
	ctx.ClearCookie(constant.JWTRefreshCookiesKey)
}
