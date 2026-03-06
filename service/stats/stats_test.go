package stats

import (
	"context"
	"testing"
	"time"

	statsentity "github.com/almanalfaruq/alfarpos-backend/model/stats"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Manual mock for statsRepo
type MockStatsRepo struct {
	mock.Mock
}

func (m *MockStatsRepo) InsertShopStats(ctx context.Context, stats statsentity.ShopStats) (statsentity.ShopStats, error) {
	args := m.Called(ctx, stats)
	return args.Get(0).(statsentity.ShopStats), args.Error(1)
}

func (m *MockStatsRepo) GetByDate(ctx context.Context, date string) (statsentity.ShopStats, error) {
	args := m.Called(ctx, date)
	return args.Get(0).(statsentity.ShopStats), args.Error(1)
}

func (m *MockStatsRepo) GetTrendingProducts(ctx context.Context, startDate, endDate time.Time, limit int) ([]statsentity.ProductStats, error) {
	args := m.Called(ctx, startDate, endDate, limit)
	return args.Get(0).([]statsentity.ProductStats), args.Error(1)
}

func (m *MockStatsRepo) GetProductSales(ctx context.Context, productID int64, startDate, endDate time.Time) (statsentity.ProductStats, error) {
	args := m.Called(ctx, productID, startDate, endDate)
	return args.Get(0).(statsentity.ProductStats), args.Error(1)
}

func (m *MockStatsRepo) GetShopStatsByDateRange(ctx context.Context, startDate, endDate time.Time) ([]statsentity.ShopStats, error) {
	args := m.Called(ctx, startDate, endDate)
	return args.Get(0).([]statsentity.ShopStats), args.Error(1)
}

func TestGetTrendingProducts(t *testing.T) {
	mockRepo := new(MockStatsRepo)
	service := New(mockRepo, nil, nil) // orderRepo and moneyRepo are not used in GetTrendingProducts

	ctx := context.Background()

	// Test "week" period
	mockRepo.On("GetTrendingProducts", ctx, mock.AnythingOfType("time.Time"), mock.AnythingOfType("time.Time"), 10).Return([]statsentity.ProductStats{}, nil)

	_, err := service.GetTrendingProducts(ctx, "week")
	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestGetShopStatsByDateRange(t *testing.T) {
	mockRepo := new(MockStatsRepo)
	service := New(mockRepo, nil, nil)

	ctx := context.Background()
	startDateStr := "2023-10-01"
	endDateStr := "2023-10-31"

	mockRepo.On("GetShopStatsByDateRange", ctx, mock.AnythingOfType("time.Time"), mock.AnythingOfType("time.Time")).Return([]statsentity.ShopStats{}, nil)

	_, err := service.GetShopStatsByDateRange(ctx, startDateStr, endDateStr)
	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestGetProductSales(t *testing.T) {
	mockRepo := new(MockStatsRepo)
	service := New(mockRepo, nil, nil)

	ctx := context.Background()
	productID := int64(1)

	// Test "month" period
	mockRepo.On("GetProductSales", ctx, productID, mock.AnythingOfType("time.Time"), mock.AnythingOfType("time.Time")).Return(statsentity.ProductStats{}, nil)

	_, err := service.GetProductSales(ctx, productID, "month")
	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}
