package jwt

import (
	constant2 "github.com/fiber-go-sis-app/internal/app/constant"
	customPkg "github.com/fiber-go-sis-app/internal/pkg/custom"
	"github.com/gofiber/fiber/v2"
	jwtWare "github.com/gofiber/jwt/v3"
)

// AccessTokenMiddleware function to handle access token middleware in web
func AccessTokenMiddleware(sourceName string) fiber.Handler {
	// define error handler if source name is web will be return on login if service will be return json
	errHandler := jwtSvcAccessTokenError
	if sourceName == constant2.WebSource {
		errHandler = jwtWebAccessTokenError
	}

	return jwtWare.New(jwtWare.Config{
		ErrorHandler:  errHandler,
		SigningMethod: constant2.JWTMethod,
		SigningKey:    customPkg.GetPrivateKey().Public(),
		TokenLookup:   "header:Authorization,cookie:" + constant2.JWTAccessCookiesKey,
	})
}

func jwtSvcAccessTokenError(ctx *fiber.Ctx, err error) error {
	if err.Error() == constant2.ErrMissingOrMalformedJWT {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": "error", "message": constant2.ErrMissingOrMalformedJWT, "data": nil})
	}

	return ctx.Next()
}

func jwtWebAccessTokenError(ctx *fiber.Ctx, err error) error {
	if err.Error() == constant2.ErrMissingOrMalformedJWT {
		return ctx.Redirect(ctx.BaseURL())
	}

	return ctx.Next()
}
