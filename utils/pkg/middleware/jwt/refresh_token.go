package jwt

import (
	"github.com/gofiber/fiber/v2"
	jwtWare "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"

	constantsEntity "github.com/fiber-go-sis-app/internal/entity/constants"

	customPkg "github.com/fiber-go-sis-app/utils/pkg/custom"

	formsUC "github.com/fiber-go-sis-app/internal/usecase/services/forms"
)

// RefreshTokenMiddleware function to handle refresh token middleware in web
func RefreshTokenMiddleware(sourceName string) fiber.Handler {
	// define error handler if source name is web will be return on login if service will be return json
	errHandler := jwtSvcRefreshTokenError
	if sourceName == constantsEntity.WebSource {
		errHandler = jwtWebRefreshTokenError
	}

	return jwtWare.New(jwtWare.Config{
		ErrorHandler:   errHandler,
		SuccessHandler: refreshTokenSuccess,
		SigningMethod:  constantsEntity.JWTMethod,
		SigningKey:     customPkg.GetPrivateKey().Public(),
		TokenLookup:    "header:Authorization,cookie:" + constantsEntity.JWTRefreshCookiesKey,
	})
}

func refreshTokenSuccess(ctx *fiber.Ctx) error {
	// Claim data
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	// convert all claims data
	userID := claims["user_id"].(string)

	if err := formsUC.RefreshTokenJWT(ctx, userID); err != nil {
		return err
	}

	return ctx.Next()
}

func jwtSvcRefreshTokenError(ctx *fiber.Ctx, err error) error {
	if err.Error() == constantsEntity.ErrMissingOrMalformedJWT {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": "error", "message": constantsEntity.ErrMissingOrMalformedJWT, "data": nil})
	}
	return ctx.Status(fiber.StatusUnauthorized).
		JSON(fiber.Map{"status": "error", "message": constantsEntity.ErrInvalidORExpiredJWT, "data": nil})
}

func jwtWebRefreshTokenError(ctx *fiber.Ctx, err error) error {
	return ctx.Redirect(constantsEntity.BaseURL)
}
