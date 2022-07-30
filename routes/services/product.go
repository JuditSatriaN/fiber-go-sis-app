package services

import (
	"github.com/gofiber/fiber/v2"

	productSvc "github.com/fiber-go-sis-app/internal/handler/services/products"
)

// BuildProductRoutes : Service - service to handle product
func BuildProductRoutes(service fiber.Router) {
	service.Get("/dt_products", productSvc.GetDTProductHandler)
	service.Get("/product", productSvc.GetProductByPLUOrBarcode)
	service.Post("/product/insert", productSvc.InsertProductHandler)
	service.Post("/product/update", productSvc.UpdateProductHandler)
	service.Post("/product/delete", productSvc.DeleteProductHandler)
	service.Post("/product/upsert", productSvc.UpsertProductHandler)
}
