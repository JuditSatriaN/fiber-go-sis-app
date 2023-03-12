package sales

import (
	"errors"
	"fmt"
	"time"

	"github.com/fiber-go-sis-app/internal/app/model"
	"github.com/gofiber/fiber/v2"

	inventoryRepo "github.com/fiber-go-sis-app/internal/app/repo/inventory"
	invoiceRepo "github.com/fiber-go-sis-app/internal/app/repo/invoice"
	salesRepo "github.com/fiber-go-sis-app/internal/app/repo/sales"
	statRepo "github.com/fiber-go-sis-app/internal/app/repo/stat"
	customPkg "github.com/fiber-go-sis-app/internal/pkg/custom"
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

	var res []model.SalesHeadResponse
	for _, v := range salesHead {
		res = append(res, model.SalesHeadResponse{
			ID:            v.ID,
			Invoice:       v.Invoice,
			UserID:        v.UserID,
			TotalItem:     v.TotalItem,
			TotalPrice:    v.TotalPrice,
			TotalPurchase: v.TotalPurchase,
			TotalTax:      v.TotalTax,
			TotalDiscount: v.TotalDiscount,
			TotalPay:      v.TotalPay,
			CreateTimeStr: v.CreateTime.Format("2006-01-02 15:04:05"),
		})
	}

	return model.ListSalesHeadDataResponse{
		Total: totalSalesHead,
		Data:  res,
	}, nil
}

func GetSalesDetailByInvoice(ctx *fiber.Ctx, invoice string) (model.ListSalesDetailDataResponse, error) {
	data, err := salesRepo.GetALlSalesDetailByInvoice(ctx, invoice)
	if err != nil {
		return model.ListSalesDetailDataResponse{}, err
	}

	var res []model.SalesDetailResponse
	for _, v := range data {
		res = append(res, model.SalesDetailResponse{
			ID:            v.ID,
			Invoice:       v.Invoice,
			UserID:        v.UserID,
			PLU:           v.PLU,
			Name:          v.Name,
			UnitName:      v.UnitName,
			Barcode:       v.Barcode,
			Ppn:           v.Ppn,
			Qty:           v.Qty,
			Price:         v.Price,
			Purchase:      v.Purchase,
			Discount:      v.Discount,
			MemberID:      v.MemberID,
			InventoryID:   v.InventoryID,
			CreateTimeStr: v.CreateTime.Format("2006-01-02 15:04:05"),
		})
	}

	return model.ListSalesDetailDataResponse{
		Total: int64(len(data)),
		Data:  res,
	}, nil
}

func InsertVoid(ctx *fiber.Ctx, invoice string) error {
	salesHead, found, err := salesRepo.GetSalesHeadByInvoice(ctx, invoice)
	if err != nil {
		return err
	}

	if !found {
		return errors.New("Data Invoice tidak ditemukan")
	}

	salesDetail, err := salesRepo.GetALlSalesDetailByInvoice(ctx, invoice)
	if err != nil {
		return err
	}

	tx, err := postgresPkg.BeginTxx(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if err := salesRepo.InsertVoidHead(ctx, tx, model.VoidHead{
		Invoice:       salesHead.Invoice,
		UserID:        salesHead.UserID,
		TotalItem:     salesHead.TotalItem,
		TotalPrice:    salesHead.TotalPrice,
		TotalPurchase: salesHead.TotalPurchase,
		TotalTax:      salesHead.TotalTax,
		TotalDiscount: salesHead.TotalDiscount,
		TotalPay:      salesHead.TotalPay,
	}); err != nil {
		return err
	}

	var voidDetails []model.VoidDetail
	for _, v := range salesDetail {
		voidDetails = append(voidDetails, model.VoidDetail{
			Invoice:     v.Invoice,
			UserID:      v.UserID,
			PLU:         v.PLU,
			Name:        v.Name,
			UnitName:    v.UnitName,
			Barcode:     v.Barcode,
			Ppn:         v.Ppn,
			Qty:         v.Qty,
			Price:       v.Price,
			Purchase:    v.Purchase,
			Discount:    v.Discount,
			MemberID:    v.MemberID,
			InventoryID: v.InventoryID,
		})
	}

	if err := salesRepo.InsertVoidDetail(ctx, tx, voidDetails); err != nil {
		return err
	}

	productSalesStats := make([]model.ProductSalesStatsDaily, len(salesDetail))
	for idx, v := range salesDetail {
		qtyVoid := v.Qty * -1

		salesData := model.UpdateStockAfterSalesData{
			ID:  v.InventoryID,
			Qty: qtyVoid,
		}
		if err := inventoryRepo.UpdateStockAfterSales(tx, salesData); err != nil {
			return err
		}

		productSalesStats[idx] = model.ProductSalesStatsDaily{
			DateSold:  v.CreateTime,
			PLU:       v.PLU,
			TotalSold: qtyVoid,
		}
	}

	if err := statRepo.BulkUpsertTotalSold(ctx, tx, productSalesStats); err != nil {
		return err
	}

	if err := salesRepo.DeleteSalesDetailByInvoice(ctx, tx, salesDetail); err != nil {
		return err
	}

	if err := salesRepo.DeleteSalesHeadByInvoice(ctx, tx, salesHead); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
