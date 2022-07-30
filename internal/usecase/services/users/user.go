package users

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	customPkg "github.com/fiber-go-sis-app/utils/pkg/custom"

	userEntity "github.com/fiber-go-sis-app/internal/entity/users"

	userRepo "github.com/fiber-go-sis-app/internal/repo/users"
)

func GetAllUser(ctx *fiber.Ctx) ([]userEntity.User, error) {
	users, err := userRepo.GetAllUser(ctx)
	if err != nil {
		return []userEntity.User{}, err
	}

	return users, nil
}

func GetUserByUserID(ctx *fiber.Ctx, userID string) (userEntity.User, error) {
	user, found, err := userRepo.GetUserByUserID(ctx, userID)
	if err != nil {
		return userEntity.User{}, err
	}

	if !found {
		return userEntity.User{}, fmt.Errorf("user dengan nama : %s tidak ditemukan", user.UserName)
	}

	return user, nil
}

func InsertUser(ctx *fiber.Ctx, user userEntity.User) error {
	if user.Password != "" {
		user.Password, _ = customPkg.HashPassword(user.Password)
	}

	return userRepo.InsertUser(ctx, user)
}

func UpdateUser(ctx *fiber.Ctx, user userEntity.User) error {
	userDB, err := GetUserByUserID(ctx, user.UserID)
	if err != nil {
		return err
	}

	// replace to existing data
	if user.Password == "" {
		user.Password = userDB.Password
	} else {
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

func UpsertUser(ctx *fiber.Ctx, user userEntity.User) error {
	userDB, found, err := userRepo.GetUserByUserID(ctx, user.UserID)
	if err != nil {
		return err
	}

	if userDB.Password != user.Password {
		user.Password, _ = customPkg.HashPassword(user.Password)
	}

	if !found {
		return userRepo.InsertUser(ctx, user)
	} else {
		return userRepo.UpdateUser(ctx, user)
	}
}
