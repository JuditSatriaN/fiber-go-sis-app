package custom

import (
	"github.com/fiber-go-sis-app/internal/app/constant"
	"github.com/gofiber/fiber/v2"
)

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
