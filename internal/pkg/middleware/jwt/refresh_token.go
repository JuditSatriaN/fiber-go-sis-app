package jwt

import (
	"github.com/fiber-go-sis-app/internal/app/constant"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"

	formsUC "github.com/fiber-go-sis-app/internal/app/usecase/utility"
	customPkg "github.com/fiber-go-sis-app/internal/pkg/custom"
	jwtWare "github.com/gofiber/jwt/v3"
)

// RefreshTokenMiddleware function to handle refresh token middleware in web
func RefreshTokenMiddleware(sourceName string) fiber.Handler {
	// Define error handler if source name is web will be return on login if service will be return json
	errHandler := jwtSvcRefreshTokenError
	if sourceName == constant.WebSource {
		errHandler = jwtWebRefreshTokenError
	}

	return jwtWare.New(jwtWare.Config{
		ErrorHandler:   errHandler,
		SigningMethod:  constant.JWTMethod,
		SuccessHandler: refreshTokenSuccess,
		SigningKey:     customPkg.GetPrivateKey().Public(),
		TokenLookup:    "header:Authorization,cookie:" + constant.JWTRefreshCookiesKey,
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
	if err.Error() == constant.ErrMissingOrMalformedJWT {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": "error", "message": constant.ErrMissingOrMalformedJWT, "data": nil})
	}
	return ctx.Status(fiber.StatusUnauthorized).
		JSON(fiber.Map{"status": "error", "message": constant.ErrInvalidORExpiredJWT, "data": nil})
}

func jwtWebRefreshTokenError(ctx *fiber.Ctx, err error) error {
	customPkg.ClearLoginCookie(ctx)
	return ctx.Redirect(constant.BaseURL)
}
