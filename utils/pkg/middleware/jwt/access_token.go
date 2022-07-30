package jwt

import (
	"github.com/gofiber/fiber/v2"
	jwtWare "github.com/gofiber/jwt/v3"

	constantsEntity "github.com/fiber-go-sis-app/internal/entity/constants"

	customPkg "github.com/fiber-go-sis-app/utils/pkg/custom"
)

// AccessTokenMiddleware function to handle access token middleware in web
func AccessTokenMiddleware(sourceName string) fiber.Handler {
	// define error handler if source name is web will be return on login if service will be return json
	errHandler := jwtSvcAccessTokenError
	if sourceName == constantsEntity.WebSource {
		errHandler = jwtWebAccessTokenError
	}

	return jwtWare.New(jwtWare.Config{
		ErrorHandler:  errHandler,
		SigningMethod: constantsEntity.JWTMethod,
		SigningKey:    customPkg.GetPrivateKey().Public(),
		TokenLookup:   "header:Authorization,cookie:" + constantsEntity.JWTAccessCookiesKey,
	})
}

func jwtSvcAccessTokenError(ctx *fiber.Ctx, err error) error {
	if err.Error() == constantsEntity.ErrMissingOrMalformedJWT {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": "error", "message": constantsEntity.ErrMissingOrMalformedJWT, "data": nil})
	}

	return ctx.Next()
}

func jwtWebAccessTokenError(ctx *fiber.Ctx, err error) error {
	if err.Error() == constantsEntity.ErrMissingOrMalformedJWT {
		return ctx.Redirect(ctx.BaseURL())
	}

	return ctx.Next()
}
