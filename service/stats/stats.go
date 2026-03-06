package stats

import (
	"context"
	"fmt"
	"math"
	"sync"
	"time"

	orderentity "github.com/almanalfaruq/alfarpos-backend/model/order"
	statsentity "github.com/almanalfaruq/alfarpos-backend/model/stats"
	transactionentity "github.com/almanalfaruq/alfarpos-backend/model/transaction"
	"github.com/almanalfaruq/alfarpos-backend/util/logger"
)

type StatsService struct {
	statsRepo statsRepo
	orderRepo orderRepositoryIface
	moneyRepo moneyRepo
}

func New(statsRepo statsRepo, orderRepo orderRepositoryIface, moneyRepo moneyRepo) *StatsService {
	return &StatsService{
		statsRepo: statsRepo,
		orderRepo: orderRepo,
		moneyRepo: moneyRepo,
	}
}

func (s *StatsService) GetShopStats(ctx context.Context, date string) (statsentity.ShopStats, error) {
	now := time.Now()
	timeParam, err := time.Parse("2006-01-02", date)
	if err != nil {
		return statsentity.ShopStats{}, err
	}
	var isToday bool
	if now.Day() == timeParam.Day() && now.Month() == timeParam.Month() && now.Year() == timeParam.Year() {
		isToday = true
	}
	// get from shop_stats table if not today
	if !isToday {
		result, err := s.statsRepo.GetByDate(ctx, date)
		if err == nil {
			return result, nil
		}
		logger.Log.Errorf("[GetShopStats] GetByDate error: %v", err)
	}

	result, err := s.calculateStatsForDate(ctx, date)
	if err != nil {
		return statsentity.ShopStats{}, err
	}

	// insert if not today
	if !isToday {
		_, err = s.statsRepo.InsertShopStats(ctx, result)
		if err != nil {
			logger.Log.Errorf("[GetShopStats] InsertShopStats error: %v", err)
		}
	}

	return result, nil
}

func (s *StatsService) calculateStatsForDate(ctx context.Context, date string) (statsentity.ShopStats, error) {
	var wg sync.WaitGroup
	wg.Add(2)

	var (
		allOrders []orderentity.Order
		errOrder  error
	)
	go func() {
		defer wg.Done()
		orders, err := s.orderRepo.FindByFilter([]int32{orderentity.StatusFinish}, "", date, date, "", 0, 0)
		if err != nil {
			errOrder = err
			return
		}
		allOrders = orders
	}()

	var (
		allTransactions []transactionentity.Money
		errTransaction  error
	)
	go func() {
		defer wg.Done()
		transactions, err := s.moneyRepo.GetMoneyTransactionByFilter([]int32{}, date, date, "")
		if err != nil {
			errTransaction = err
			return
		}
		allTransactions = transactions
	}()

	wg.Wait()

	switch {
	case errOrder != nil:
		return statsentity.ShopStats{}, fmt.Errorf("Error getting orders: %v", errOrder)
	case errTransaction != nil:
		return statsentity.ShopStats{}, fmt.Errorf("Error getting money transactions: %v", errTransaction)
	}

	wg = sync.WaitGroup{}
	wg.Add(2)

	var result statsentity.ShopStats
	go func() {
		defer wg.Done()
		var grossProfit, netProfit int64
		for _, order := range allOrders {
			grossProfit += order.Total
			for _, odetail := range order.OrderDetails {
				netProfit += odetail.SubTotal - (int64(odetail.Quantity) * odetail.Product.BuyPrice.Int64)
			}
		}
		result.GrossProfit = grossProfit
		result.NetProfit = netProfit
		result.OrderCount = int32(len(allOrders))
		result.SellAverage = float64(grossProfit) / float64(len(allOrders))
		if math.IsNaN(result.SellAverage) {
			result.SellAverage = 0
		}
	}()

	go func() {
		defer wg.Done()
		var moneyIn, moneyOut float64
		for _, transaction := range allTransactions {
			if transaction.Type == transactionentity.TypeIn {
				moneyIn += transaction.Amount
			} else {
				moneyOut += transaction.Amount
			}
		}
		result.MoneyIn = moneyIn
		result.MoneyOut = moneyOut
	}()

	wg.Wait()

	timeParam, _ := time.Parse("2006-01-02", date)
	result.Date = timeParam

	return result, nil
}

func (s *StatsService) GetTrendingProducts(ctx context.Context, period string) ([]statsentity.ProductStats, error) {
	now := time.Now()
	var startDate, endDate time.Time
	endDate = now

	switch period {
	case "week":
		startDate = now.AddDate(0, 0, -7)
	case "month":
		startDate = now.AddDate(0, -1, 0)
	case "year":
		startDate = now.AddDate(-1, 0, 0)
	default:
		// Default to 1 month
		startDate = now.AddDate(0, -1, 0)
	}

	return s.statsRepo.GetTrendingProducts(ctx, startDate, endDate, 10)
}

func (s *StatsService) GetProductSales(ctx context.Context, productID int64, period string) (statsentity.ProductStats, error) {
	now := time.Now()
	var startDate, endDate time.Time
	endDate = now

	switch period {
	case "week":
		startDate = now.AddDate(0, 0, -7)
	case "month":
		startDate = now.AddDate(0, -1, 0)
	case "year":
		startDate = now.AddDate(-1, 0, 0)
	default:
		// Default to 1 month
		startDate = now.AddDate(0, -1, 0)
	}

	return s.statsRepo.GetProductSales(ctx, productID, startDate, endDate)
}

func (s *StatsService) GetShopStatsByDateRange(ctx context.Context, startDateStr, endDateStr string) ([]statsentity.ShopStats, error) {
	startDate, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		return nil, fmt.Errorf("invalid start date format: %v", err)
	}
	endDate, err := time.Parse("2006-01-02", endDateStr)
	if err != nil {
		return nil, fmt.Errorf("invalid end date format: %v", err)
	}

	// Set endDate to end of the day
	endDate = endDate.Add(23*time.Hour + 59*time.Minute + 59*time.Second)

	results, err := s.statsRepo.GetShopStatsByDateRange(ctx, startDate, endDate)
	if err != nil {
		return nil, err
	}

	// Check if today is in range
	now := time.Now()
	todayStr := now.Format("2006-01-02")
	todayDate, _ := time.Parse("2006-01-02", todayStr)

	// Check if range covers today: startDate <= today <= endDate (ignoring time for comparison)
	// We compare dates by truncating time or using the parsed todayDate (which is midnight)
	// startDate is midnight. endDate is 23:59:59.
	if (startDate.Before(todayDate) || startDate.Equal(todayDate)) && (endDate.After(todayDate) || endDate.Equal(todayDate)) {
		// Check if results already contain today
		hasToday := false
		for _, r := range results {
			if r.Date.Format("2006-01-02") == todayStr {
				hasToday = true
				break
			}
		}

		if !hasToday {
			todayStats, err := s.calculateStatsForDate(ctx, todayStr)
			if err == nil {
				results = append(results, todayStats)
			} else {
				logger.Log.Errorf("[GetShopStatsByDateRange] calculateStatsForDate error: %v", err)
			}
		}
	}

	return results, nil
}
