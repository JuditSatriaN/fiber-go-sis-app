package user

import (
	"fmt"

	"github.com/fiber-go-sis-app/internal/app/model"
	"github.com/gofiber/fiber/v2"

	userRepo "github.com/fiber-go-sis-app/internal/app/repo/user"
	customPkg "github.com/fiber-go-sis-app/internal/pkg/custom"
)

func GetAllUser(ctx *fiber.Ctx) ([]model.User, error) {
	users, err := userRepo.GetAllUser(ctx)
	if err != nil {
		return []model.User{}, err
	}

	return users, nil
}

func GetUserByUserID(ctx *fiber.Ctx, userID string) (model.User, error) {
	user, found, err := userRepo.GetUserByUserID(ctx, userID)
	if err != nil {
		return model.User{}, err
	}

	if !found {
		return model.User{}, fmt.Errorf("user dengan nama : %s tidak ditemukan", user.UserName)
	}

	return user, nil
}

func InsertUser(ctx *fiber.Ctx, user model.User) error {
	user.Password, _ = customPkg.HashPassword(user.Password)
	return userRepo.InsertUser(ctx, user)
}

func UpdateUser(ctx *fiber.Ctx, user model.User) error {
	userDB, err := GetUserByUserID(ctx, user.UserID)
	if err != nil {
		return err
	}

	// replace to existing data
	if userDB.Password != user.Password {
		user.Password, _ = customPkg.HashPassword(user.Password)
	}

	return userRepo.UpdateUser(ctx, user)
}

func DeleteUser(ctx *fiber.Ctx, userID string) error {
	if _, err := GetUserByUserID(ctx, userID); err != nil {
		return err
	}

	return userRepo.DeleteUser(ctx, userID)
}

func UpsertUser(ctx *fiber.Ctx, user model.User) error {
	userDB, found, err := userRepo.GetUserByUserID(ctx, user.UserID)
	if err != nil {
		return err
	}

	// replace to existing data
	if userDB.Password != user.Password {
		user.Password, _ = customPkg.HashPassword(user.Password)
	}

	if !found {
		return userRepo.InsertUser(ctx, user)
	} else {
		return userRepo.UpdateUser(ctx, user)
	}
}
