package forms

import (
	"github.com/gofiber/fiber/v2"

	customPkg "github.com/fiber-go-sis-app/utils/pkg/custom"

	constantsEntity "github.com/fiber-go-sis-app/internal/entity/constants"
	formsEntity "github.com/fiber-go-sis-app/internal/entity/forms"

	userRepo "github.com/fiber-go-sis-app/internal/repo/users"
)

func LoginForm(ctx *fiber.Ctx, req formsEntity.LoginRequest) (formsEntity.LoginResponse, error) {
	// Initialization variable
	var res formsEntity.LoginResponse

	data, err := userRepo.GetUserByUserName(ctx, req.UserName)
	if err != nil {
		return res, err
	}

	// check hash password
	if !customPkg.CheckPasswordHash(req.Password, data.Password) {
		return res, constantsEntity.ErrWrongPassword
	}

	token, err := customPkg.CreateJWTToken(formsEntity.JWTRequest{
		UserID:  data.UserID,
		Name:    data.UserName,
		IsAdmin: data.IsAdmin,
	})
	if err != nil {
		return res, err
	}

	return formsEntity.LoginResponse{
		UserID:          data.UserID,
		UserName:        data.UserName,
		FullName:        data.FullName,
		IsAdmin:         data.IsAdmin,
		JWTAccessToken:  token.AccessToken,
		JWTRefreshToken: token.RefreshToken,
	}, nil
}
