package forms

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	constantsEntity "github.com/fiber-go-sis-app/internal/entity/constants"
	formsEntity "github.com/fiber-go-sis-app/internal/entity/forms"

	customPkg "github.com/fiber-go-sis-app/utils/pkg/custom"

	userRepo "github.com/fiber-go-sis-app/internal/repo/users"
)

func RefreshTokenJWT(ctx *fiber.Ctx, userID string) error {
	data, found, err := userRepo.GetUserByUserID(ctx, userID)
	if err != nil {
		return err
	}

	if !found {
		return fmt.Errorf("user tidak ditemukan silahkan login kembali")
	}

	token, err := customPkg.CreateJWTToken(formsEntity.JWTRequest{
		UserID:  data.UserID,
		Name:    data.UserName,
		IsAdmin: data.IsAdmin,
	})
	if err != nil {
		return err
	}

	ctx.Cookie(&fiber.Cookie{
		Name:  constantsEntity.JWTAccessCookiesKey,
		Value: token.AccessToken,
	})

	ctx.Cookie(&fiber.Cookie{
		Name:  constantsEntity.JWTRefreshCookiesKey,
		Value: token.RefreshToken,
	})

	return nil
}
