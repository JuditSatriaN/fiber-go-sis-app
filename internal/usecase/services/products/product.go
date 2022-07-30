package products

import (
	"errors"
	"fmt"
	storeEntity "github.com/fiber-go-sis-app/internal/entity/stores"
	customPkg "github.com/fiber-go-sis-app/utils/pkg/custom"

	"github.com/gofiber/fiber/v2"

	constantsEntity "github.com/fiber-go-sis-app/internal/entity/constants"
	productEntity "github.com/fiber-go-sis-app/internal/entity/products"

	productRepo "github.com/fiber-go-sis-app/internal/repo/products"
	storeRepo "github.com/fiber-go-sis-app/internal/repo/stores"
)

// GetDTAllProduct : Get List Of Product for Datatable
func GetDTAllProduct(ctx *fiber.Ctx, page int, limit int, search string) (productEntity.ListProductDataResponse, error) {
	offset := customPkg.BuildOffset(page, limit)

	products, err := productRepo.GetALlProducts(ctx, search, limit, offset)
	if err != nil {
		return productEntity.ListProductDataResponse{}, err
	}

	storeStats, err := storeRepo.GetStoreStats(ctx, storeEntity.DefaultStoreID)
	if err != nil {
		return productEntity.ListProductDataResponse{}, err
	}

	return productEntity.ListProductDataResponse{
		Total: storeStats.CntProduct,
		Data:  products,
	}, nil
}

func GetProductByPLU(ctx *fiber.Ctx, PLU string) (productEntity.Product, error) {
	product, found, err := productRepo.GetProductByPLU(ctx, PLU)
	if err != nil {
		return productEntity.Product{}, err
	}

	if !found {
		return productEntity.Product{}, fmt.Errorf("product dengan nama : %s tidak ditemukan", product.Name)
	}

	return product, nil
}

func GetProductByPLUOrBarcode(ctx *fiber.Ctx, search string) (productEntity.Product, error) {
	product, found, err := productRepo.GetProductByPLUOrBarcode(ctx, search)
	if err != nil {
		return productEntity.Product{}, err
	}

	if !found {
		return productEntity.Product{}, errors.New(constantsEntity.ErrNoDataFound)
	}

	return product, nil
}

func InsertProduct(ctx *fiber.Ctx, product productEntity.Product) error {
	if err := productRepo.InsertProduct(ctx, product); err != nil {
		return err
	}

	return nil
}

func UpdateProduct(ctx *fiber.Ctx, product productEntity.Product) error {
	if _, err := GetProductByPLU(ctx, product.PLU); err != nil {
		return err
	}

	if err := productRepo.UpdateProduct(ctx, product); err != nil {
		return err
	}
	return nil
}

func DeleteProduct(ctx *fiber.Ctx, PLU string) error {
	if _, err := GetProductByPLU(ctx, PLU); err != nil {
		return err
	}

	if err := productRepo.DeleteProduct(ctx, PLU); err != nil {
		return err
	}

	err := storeRepo.UpdateTotalProduct(ctx, storeEntity.StoreStats{
		StoreID:    storeEntity.DefaultStoreID,
		CntProduct: -1,
	})

	return err
}

func UpsertProduct(ctx *fiber.Ctx, product productEntity.Product) error {
	var err error
	_, found, err := productRepo.GetProductByPLU(ctx, product.PLU)
	if err != nil {
		return err
	}

	if !found {
		err = productRepo.InsertProduct(ctx, product)
		err = storeRepo.UpdateTotalProduct(ctx, storeEntity.StoreStats{
			StoreID:    storeEntity.DefaultStoreID,
			CntProduct: 1,
		})
	} else {
		err = productRepo.UpdateProduct(ctx, product)
	}

	if err != nil {
		return err
	}

	return nil
}
