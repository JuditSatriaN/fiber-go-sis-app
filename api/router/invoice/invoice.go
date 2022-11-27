package member

import (
	"github.com/gofiber/fiber/v2"

	invoiceHandler "github.com/fiber-go-sis-app/api/handler/invoice"
)

// BuildInvoiceAPI : API to handle member
func BuildInvoiceAPI(api fiber.Router) {
	api.Get("invoice/generate_invoice", invoiceHandler.GetInvoiceHandler)
	api.Post("invoice/update_invoice", invoiceHandler.UpdateInvoiceHandler)
}
