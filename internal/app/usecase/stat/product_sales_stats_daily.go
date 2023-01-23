package stat

import (
	"github.com/fiber-go-sis-app/internal/app/model"
	"github.com/gofiber/fiber/v2"
	"sync"

	statRepo "github.com/fiber-go-sis-app/internal/app/repo/stat"
)

func GetTop3ProductSalesDaily(ctx *fiber.Ctx) ([]model.ProductSalesStatsDaily, error) {
	results, err := statRepo.GetProductSalesStatsDaily(ctx, 3)
	if err != nil {
		return []model.ProductSalesStatsDaily{}, err
	}

	return results, nil
}

func GetStatisticDashboard(ctx *fiber.Ctx) (model.StatisticDashboard, error) {
	var err error
	var wg sync.WaitGroup
	var totalProduct int64
	var totalPemasukan float64
	var totalTransaksi float64
	var totalPendapatan float64

	wg.Add(4)
	go func() {
		defer wg.Done()
		totalProduct, err = statRepo.GetTotalProductSoldToday(ctx)
	}()
	go func() {
		defer wg.Done()
		totalPemasukan, err = statRepo.GetTotalPemasukanToday(ctx)
	}()
	go func() {
		defer wg.Done()
		totalTransaksi, err = statRepo.GetTotalTransaksiToday(ctx)
	}()
	go func() {
		defer wg.Done()
		totalPendapatan, err = statRepo.GetTotalPendapatanToday(ctx)
	}()
	wg.Wait()
	if err != nil {
		return model.StatisticDashboard{}, err
	}

	return model.StatisticDashboard{
		TotalProductTerjualHariIni: totalProduct,
		TotalPemasukanHariIni:      totalPemasukan,
		TotalPendapatanHariIni:     totalPendapatan,
		TotalTransaksiHariIni:      totalTransaksi,
	}, nil
}
