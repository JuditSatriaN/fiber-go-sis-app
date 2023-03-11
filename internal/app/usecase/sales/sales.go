package sales

import (
	"fmt"
	inventoryRepo "github.com/fiber-go-sis-app/internal/app/repo/inventory"
	customPkg "github.com/fiber-go-sis-app/internal/pkg/custom"
	"time"

	"github.com/fiber-go-sis-app/internal/app/model"
	"github.com/gofiber/fiber/v2"

	invoiceRepo "github.com/fiber-go-sis-app/internal/app/repo/invoice"
	salesRepo "github.com/fiber-go-sis-app/internal/app/repo/sales"
	statRepo "github.com/fiber-go-sis-app/internal/app/repo/stat"
	postgresPkg "github.com/fiber-go-sis-app/internal/pkg/database/postgres"
)

func getHeadFak() string {
	now := time.Now()
	headFakMonth := model.TranslateMonthHeadFak.Replace(now.Month().String())
	lastTwoDigits := now.Year() % 1e2
	headFakYear := fmt.Sprintf("%02d", lastTwoDigits)
	return fmt.Sprintf("INV-%s-%s", headFakMonth, headFakYear)
}

func InsertSales(ctx *fiber.Ctx, sales model.Sales) error {
	tx, err := postgresPkg.BeginTxx(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if err := salesRepo.InsertSalesHead(ctx, tx, sales.Head); err != nil {
		return err
	}

	if err := salesRepo.InsertSalesDetail(ctx, tx, sales.Detail); err != nil {
		return err
	}

	productSalesStats := make([]model.ProductSalesStatsDaily, len(sales.Detail))
	for idx, data := range sales.Detail {
		salesData := model.UpdateStockAfterSalesData{
			ID:  data.InventoryID,
			Qty: data.Qty,
		}
		if err := inventoryRepo.UpdateStockAfterSales(tx, salesData); err != nil {
			return err
		}

		productSalesStats[idx] = model.ProductSalesStatsDaily{
			DateSold:  time.Now(),
			PLU:       data.PLU,
			TotalSold: data.Qty,
		}
	}

	if err := statRepo.BulkUpsertTotalSold(ctx, tx, productSalesStats); err != nil {
		return err
	}

	if err := invoiceRepo.UpsertInvoice(ctx, model.Invoice{
		HeadFak: getHeadFak(),
	}); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

// GetDTAllSalesHead : Get List Of Sales Head For Datatable
func GetDTAllSalesHead(ctx *fiber.Ctx, page int, limit int, search string) (model.ListSalesHeadDataResponse, error) {
	offset := customPkg.BuildOffset(page, limit)

	salesHead, err := salesRepo.GetALlSalesHead(ctx, search, limit, offset)
	if err != nil {
		return model.ListSalesHeadDataResponse{}, err
	}

	totalSalesHead, err := salesRepo.GetTotalSalesHead(ctx, search)
	if err != nil {
		return model.ListSalesHeadDataResponse{}, err
	}

	return model.ListSalesHeadDataResponse{
		Total: totalSalesHead,
		Data:  salesHead,
	}, nil
}

func GetSalesDetailByInvoice(ctx *fiber.Ctx, invoice string) (model.ListSalesDetailDataResponse, error) {
	data, err := salesRepo.GetALlSalesDetailByInvoice(ctx, invoice)
	if err != nil {
		return model.ListSalesDetailDataResponse{}, err
	}

	return model.ListSalesDetailDataResponse{
		Total: int64(len(data)),
		Data:  data,
	}, nil
}
