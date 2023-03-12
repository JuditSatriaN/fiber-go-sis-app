package invoice

import (
	"fmt"
	"github.com/fiber-go-sis-app/internal/app/model"
	invoiceRepo "github.com/fiber-go-sis-app/internal/app/repo/invoice"
	"github.com/gofiber/fiber/v2"
	"time"
)

func getHeadFak() string {
	now := time.Now()
	headFakMonth := model.TranslateMonthHeadFak.Replace(now.Month().String())
	lastTwoDigits := now.Year() % 1e2
	headFakYear := fmt.Sprintf("%02d", lastTwoDigits)
	return fmt.Sprintf("INV-%s-%s", headFakMonth, headFakYear)
}

func generateInvoice(headFak string, lastCounter int) string {
	return fmt.Sprintf("%s-%06d", headFak, lastCounter+1)
}

func GetInvoice(ctx *fiber.Ctx) (model.InvoiceResp, error) {

	headFak := getHeadFak()
	invoiceData, _, err := invoiceRepo.GetInvoice(ctx, headFak)

	if err != nil {
		return model.InvoiceResp{}, err
	}

	return model.InvoiceResp{Data: generateInvoice(headFak, invoiceData.LastCounter)}, nil
}

func UpdateInvoice(ctx *fiber.Ctx) error {
	headFak := getHeadFak()
	err := invoiceRepo.UpsertInvoice(ctx, model.Invoice{
		HeadFak: headFak,
	})
	if err != nil {
		return err
	}

	return nil
}
