package system_conf

import (
	"fmt"
	"github.com/fiber-go-sis-app/internal/app/model"
	"github.com/gofiber/fiber/v2"

	systemConfRepo "github.com/fiber-go-sis-app/internal/app/repo/system_conf"
)

func GetAllSystemConf(ctx *fiber.Ctx) ([]model.SystemConf, error) {
	data, err := systemConfRepo.GetAllSystemConf(ctx)
	if err != nil {
		return []model.SystemConf{}, err
	}

	return data, nil
}

func GetKeyVoid(ctx *fiber.Ctx) (model.SystemConf, error) {
	data, err := systemConfRepo.GetKeyVoid(ctx)
	if err != nil {
		return model.SystemConf{}, err
	}

	return model.SystemConf{Value: data}, nil
}

func GetSystemConfByID(ctx *fiber.Ctx, ID string) (model.SystemConf, error) {
	data, found, err := systemConfRepo.GetSystemConfByID(ctx, ID)
	if err != nil {
		return model.SystemConf{}, err
	}

	if !found {
		return model.SystemConf{}, fmt.Errorf("system config dengan id : %s tidak ditemukan", ID)
	}

	return data, nil
}

func InsertSystemConf(ctx *fiber.Ctx, data model.SystemConf) error {
	return systemConfRepo.InsertSystemConf(ctx, data)
}

func UpdateSystemConf(ctx *fiber.Ctx, data model.SystemConf) error {
	if _, err := GetSystemConfByID(ctx, data.ID); err != nil {
		return err
	}

	return systemConfRepo.UpdateSystemConf(ctx, data)
}

func DeleteSystemConf(ctx *fiber.Ctx, ID string) error {
	if _, err := GetSystemConfByID(ctx, ID); err != nil {
		return err
	}

	return systemConfRepo.DeleteSystemConf(ctx, ID)
}

func UpsertSystemConf(ctx *fiber.Ctx, data model.SystemConf) error {
	_, found, err := systemConfRepo.GetSystemConfByID(ctx, data.ID)
	if err != nil {
		return err
	}

	if !found || data.ID == "" {
		return systemConfRepo.InsertSystemConf(ctx, data)
	} else {
		return systemConfRepo.UpdateSystemConf(ctx, data)
	}
}
