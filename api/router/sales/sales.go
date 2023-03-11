package sales

import (
	"github.com/gofiber/fiber/v2"

	salesHandler "github.com/fiber-go-sis-app/api/handler/sales"
)

// BuildSalesAPI : API to handle sales
func BuildSalesAPI(api fiber.Router) {
	api.Post("/sales", salesHandler.InsertSalesHandler)
	api.Get("/list_sales_head", salesHandler.GetSalesHeadHandler)
	api.Get("/list_sales_detail_by_invoice", salesHandler.GetSalesDetailByInvoiceHandler)
}
