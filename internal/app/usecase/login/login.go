package login

import (
	"github.com/fiber-go-sis-app/internal/app/constant"
	"github.com/fiber-go-sis-app/internal/app/model"
	"github.com/gofiber/fiber/v2"

	userRepo "github.com/fiber-go-sis-app/internal/app/repo/user"
	customPkg "github.com/fiber-go-sis-app/internal/pkg/custom"
)

func ProcessLoginForm(ctx *fiber.Ctx, req model.LoginRequest) (model.LoginResponse, error) {
	// Initialization variable
	var res model.LoginResponse

	data, err := userRepo.GetUserByUserName(ctx, req.UserName)
	if err != nil {
		return res, err
	}

	// Check hash password
	if !customPkg.CheckPasswordHash(req.Password, data.Password) {
		return res, constant.ErrWrongPassword
	}

	// Create login token to set in cookie
	token, err := customPkg.CreateJWTToken(constant.JWTRequest{
		UserID:  data.UserID,
		Name:    data.UserName,
		IsAdmin: data.IsAdmin,
	})
	if err != nil {
		return res, err
	}

	// Set login cookie
	customPkg.SetLoginCookie(ctx, token.AccessToken, token.RefreshToken)

	return model.LoginResponse{
		UserID:          data.UserID,
		UserName:        data.UserName,
		FullName:        data.FullName,
		IsAdmin:         data.IsAdmin,
		JWTAccessToken:  token.AccessToken,
		JWTRefreshToken: token.RefreshToken,
	}, nil
}
