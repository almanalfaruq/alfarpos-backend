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

	// get from each table order and transaction if shop_stats return no data
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
		transactions, err := s.moneyRepo.GetMoneyTransactionByFilter([]int32{}, date, date)
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
				netProfit += int64(odetail.Quantity) * odetail.Product.BuyPrice.Int64
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

	result.Date = timeParam

	// insert if not today
	if !isToday {
		_, err = s.statsRepo.InsertShopStats(ctx, result)
		if err != nil {
			logger.Log.Errorf("[GetShopStats] InsertShopStats error: %v", err)
		}
	}

	return result, nil
}
