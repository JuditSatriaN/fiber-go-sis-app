package utility

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"github.com/fiber-go-sis-app/internal/app/constant"

	customPkg "github.com/fiber-go-sis-app/internal/pkg/custom"

	userRepo "github.com/fiber-go-sis-app/internal/app/repo/user"
)

func RefreshTokenJWT(ctx *fiber.Ctx, userID string) error {
	data, found, err := userRepo.GetUserByUserID(ctx, userID)
	if err != nil {
		return err
	}

	if !found {
		return fmt.Errorf("user tidak ditemukan silahkan login kembali")
	}

	token, err := customPkg.CreateJWTToken(constant.JWTRequest{
		UserID:  data.UserID,
		Name:    data.UserName,
		IsAdmin: data.IsAdmin,
	})
	if err != nil {
		return err
	}

	ctx.Cookie(&fiber.Cookie{
		Name:  constant.JWTAccessCookiesKey,
		Value: token.AccessToken,
	})

	ctx.Cookie(&fiber.Cookie{
		Name:  constant.JWTRefreshCookiesKey,
		Value: token.RefreshToken,
	})

	return nil
}
