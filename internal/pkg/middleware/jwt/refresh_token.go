package jwt

import (
	constant2 "github.com/fiber-go-sis-app/internal/app/constant"
	formsUC "github.com/fiber-go-sis-app/internal/app/usecase/utility"
	customPkg "github.com/fiber-go-sis-app/internal/pkg/custom"
	"github.com/gofiber/fiber/v2"
	jwtWare "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
)

// RefreshTokenMiddleware function to handle refresh token middleware in web
func RefreshTokenMiddleware(sourceName string) fiber.Handler {
	// define error handler if source name is web will be return on login if service will be return json
	errHandler := jwtSvcRefreshTokenError
	if sourceName == constant2.WebSource {
		errHandler = jwtWebRefreshTokenError
	}

	return jwtWare.New(jwtWare.Config{
		ErrorHandler:   errHandler,
		SuccessHandler: refreshTokenSuccess,
		SigningMethod:  constant2.JWTMethod,
		SigningKey:     customPkg.GetPrivateKey().Public(),
		TokenLookup:    "header:Authorization,cookie:" + constant2.JWTRefreshCookiesKey,
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
	if err.Error() == constant2.ErrMissingOrMalformedJWT {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": "error", "message": constant2.ErrMissingOrMalformedJWT, "data": nil})
	}
	return ctx.Status(fiber.StatusUnauthorized).
		JSON(fiber.Map{"status": "error", "message": constant2.ErrInvalidORExpiredJWT, "data": nil})
}

func jwtWebRefreshTokenError(ctx *fiber.Ctx, err error) error {
	return ctx.Redirect(constant2.BaseURL)
}
