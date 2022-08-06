package product

import (
	"github.com/fiber-go-sis-app/internal/app/model"
	"github.com/gofiber/fiber/v2"

	productUC "github.com/fiber-go-sis-app/internal/app/usecase/product"
	customPkg "github.com/fiber-go-sis-app/internal/pkg/custom"
)

// GetALLDTProductHandler : Get List Of Product for Datatable
func GetALLDTProductHandler(ctx *fiber.Ctx) error {
	page, limit, err := customPkg.BuildPageAndLimit(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	search := ctx.Query("search", "")

	result, err := productUC.GetAllDTProduct(ctx, page, limit, search)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return customPkg.BuildDatatableRes(ctx, result.Total, result.Data)
}

// GetALLProductHandler : Get List Of Product
func GetALLProductHandler(ctx *fiber.Ctx) error {
	page, limit, err := customPkg.BuildPageAndLimit(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	search := ctx.Query("search", "")

	result, err := productUC.GetAllProduct(ctx, page, limit, search)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return customPkg.BuildJSONRes(ctx, result)
}

// GetProductByPLUOrBarcode : Get product data from PLU or barcode
func GetProductByPLUOrBarcode(ctx *fiber.Ctx) error {
	search := ctx.Query("search", "")
	if search == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Inputan tidak boleh kosong!",
		})
	}

	result, err := productUC.GetProductByPLUOrBarcode(ctx, search)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return customPkg.BuildJSONRes(ctx, result)
}

// InsertProductHandler : Insert product data
func InsertProductHandler(ctx *fiber.Ctx) error {
	var product model.Product

	if err := customPkg.ValidateRequest(ctx, &product); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := productUC.InsertProduct(ctx, product); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.SendString("Data product berhasil disimpan")
}

// UpdateProductHandler : Update product data
func UpdateProductHandler(ctx *fiber.Ctx) error {
	var product model.Product

	if err := customPkg.ValidateRequest(ctx, &product); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := productUC.UpdateProduct(ctx, product); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.SendString("Data product berhasil diubah")
}

// DeleteProductHandler : Delete product data
func DeleteProductHandler(ctx *fiber.Ctx) error {
	var product model.Product

	if err := customPkg.ValidateRequest(ctx, &product); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := productUC.DeleteProduct(ctx, product.PLU); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.SendString("Data product berhasil dihapus")
}

// UpsertProductHandler : Upsert product data
func UpsertProductHandler(ctx *fiber.Ctx) error {
	var product model.Product

	if err := customPkg.ValidateRequest(ctx, &product); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := productUC.UpsertProduct(ctx, product); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return customPkg.BuildJSONRes(ctx, product)
}
