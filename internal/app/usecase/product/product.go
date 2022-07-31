package product

import (
	"errors"
	"fmt"

	"github.com/fiber-go-sis-app/internal/app/constant"
	"github.com/fiber-go-sis-app/internal/app/model"
	"github.com/gofiber/fiber/v2"

	productRepo "github.com/fiber-go-sis-app/internal/app/repo/product"
	storeRepo "github.com/fiber-go-sis-app/internal/app/repo/store"
	customPkg "github.com/fiber-go-sis-app/internal/pkg/custom"
	postgresPkg "github.com/fiber-go-sis-app/internal/pkg/databases/postgres"
)

// GetAllDTProduct : Get List Of Product for Datatable
func GetAllDTProduct(ctx *fiber.Ctx, page int, limit int, search string) (model.ListProductDataResponse, error) {
	offset := customPkg.BuildOffset(page, limit)

	products, err := productRepo.GetALlProducts(ctx, search, limit, offset)
	if err != nil {
		return model.ListProductDataResponse{}, err
	}

	totalProduct, err := storeRepo.GetTotalProduct(ctx, model.DefaultStoreID)
	if err != nil {
		return model.ListProductDataResponse{}, err
	}

	return model.ListProductDataResponse{
		Total: totalProduct,
		Data:  products,
	}, nil
}

func GetProductByPLU(ctx *fiber.Ctx, PLU string) (model.Product, error) {
	product, found, err := productRepo.GetProductByPLU(ctx, PLU)
	if err != nil {
		return model.Product{}, err
	}

	if !found {
		return model.Product{}, fmt.Errorf("product dengan nama : %s tidak ditemukan", product.Name)
	}

	return product, nil
}

func GetProductByPLUOrBarcode(ctx *fiber.Ctx, search string) (model.Product, error) {
	product, found, err := productRepo.GetProductByPLUOrBarcode(ctx, search)
	if err != nil {
		return model.Product{}, err
	}

	if !found {
		return model.Product{}, errors.New(constant.ErrNoDataFound)
	}

	return product, nil
}

func InsertProduct(ctx *fiber.Ctx, product model.Product) error {
	tx, err := postgresPkg.BeginTxx(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if err := productRepo.InsertProduct(tx, product); err != nil {
		return err
	}

	if err := storeRepo.UpdateTotalProduct(tx, model.StoreStats{
		StoreID:      model.DefaultStoreID,
		TotalProduct: 1,
	}); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func UpdateProduct(ctx *fiber.Ctx, product model.Product) error {
	if _, err := GetProductByPLU(ctx, product.PLU); err != nil {
		return err
	}
	return productRepo.UpdateProduct(ctx, product)
}

func DeleteProduct(ctx *fiber.Ctx, PLU string) error {
	if _, err := GetProductByPLU(ctx, PLU); err != nil {
		return err
	}

	tx, err := postgresPkg.BeginTxx(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if err := productRepo.DeleteProduct(tx, PLU); err != nil {
		return err
	}

	if err := storeRepo.UpdateTotalProduct(tx, model.StoreStats{
		StoreID:      model.DefaultStoreID,
		TotalProduct: -1,
	}); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func UpsertProduct(ctx *fiber.Ctx, product model.Product) error {
	_, found, err := productRepo.GetProductByPLU(ctx, product.PLU)
	if err != nil {
		return err
	}

	if !found {
		if err := InsertProduct(ctx, product); err != nil {
			return err
		}
	} else {
		if err := UpdateProduct(ctx, product); err != nil {
			return err
		}
	}

	return nil
}
