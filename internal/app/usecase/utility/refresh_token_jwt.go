package utility

import (
	"fmt"

	"github.com/fiber-go-sis-app/internal/app/constant"
	"github.com/gofiber/fiber/v2"

	userRepo "github.com/fiber-go-sis-app/internal/app/repo/user"
	customPkg "github.com/fiber-go-sis-app/internal/pkg/custom"
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

	customPkg.SetLoginCookie(ctx, token.AccessToken, token.RefreshToken)

	return nil
}
