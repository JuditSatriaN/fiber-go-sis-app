package jwt

import (
	"github.com/fiber-go-sis-app/internal/app/constant"
	"github.com/gofiber/fiber/v2"

	customPkg "github.com/fiber-go-sis-app/internal/pkg/custom"
	jwtWare "github.com/gofiber/jwt/v3"
)

// AccessTokenMiddleware function to handle access token middleware in web
func AccessTokenMiddleware(sourceName string) fiber.Handler {
	// Define error handler if source name is web will be return on login if service will be return json
	errHandler := jwtSvcAccessTokenError
	if sourceName == constant.WebSource {
		errHandler = jwtWebAccessTokenError
	}

	return jwtWare.New(jwtWare.Config{
		ErrorHandler:  errHandler,
		SigningMethod: constant.JWTMethod,
		SigningKey:    customPkg.GetPrivateKey().Public(),
		TokenLookup:   "header:Authorization,cookie:" + constant.JWTAccessCookiesKey,
	})
}

//jwtSvcAccessTokenError custom package to handle jwt service token error
func jwtSvcAccessTokenError(ctx *fiber.Ctx, err error) error {
	if err.Error() == constant.ErrMissingOrMalformedJWT {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": "error", "message": constant.ErrMissingOrMalformedJWT, "data": nil})
	}

	return ctx.Next()
}

//jwtWebAccessTokenError custom package to handle jwt web token error
func jwtWebAccessTokenError(ctx *fiber.Ctx, err error) error {
	if err.Error() == constant.ErrMissingOrMalformedJWT {
		return ctx.Redirect(ctx.BaseURL())
	}

	return ctx.Next()
}
