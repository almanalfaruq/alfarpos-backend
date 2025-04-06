package order

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/almanalfaruq/alfarpos-backend/model"
	orderentity "github.com/almanalfaruq/alfarpos-backend/model/order"
	"github.com/almanalfaruq/alfarpos-backend/util/logger"
	"github.com/almanalfaruq/alfarpos-backend/util/memcache"
)

type OrderService struct {
	order       orderRepositoryIface
	orderDetail orderDetailRepositoryIface
	payment     paymentRepositoryIface
	customer    customerRepositoryIface
	product     productRepositoryIface
}

func NewOrderService(orderRepo orderRepositoryIface, orderDetailRepo orderDetailRepositoryIface, paymentRepo paymentRepositoryIface,
	customerRepo customerRepositoryIface, productRepo productRepositoryIface) *OrderService {
	return &OrderService{
		order:       orderRepo,
		orderDetail: orderDetailRepo,
		payment:     paymentRepo,
		customer:    customerRepo,
		product:     productRepo,
	}
}

func (s *OrderService) GetAllOrder() ([]orderentity.Order, error) {
	return s.order.FindAll(), nil
}

func (s *OrderService) GetOneOrder(id int64) (orderentity.Order, error) {
	order, err := s.order.FindById(id)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return orderentity.Order{}, fmt.Errorf("Order with id: %d is not found", id)
		}
		return orderentity.Order{}, err
	}
	return order, nil
}

func (s *OrderService) GetOrderByInvoice(invoice string) (orderentity.Order, error) {
	invoice = strings.ToLower(invoice)
	order, err := s.order.FindByInvoice(invoice)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return orderentity.Order{}, fmt.Errorf("Order with invoice: %s is not found", invoice)
		}
		return orderentity.Order{}, err
	}
	return order, nil
}

func (s *OrderService) GetOrderByUserId(userId int64) ([]orderentity.Order, error) {
	orders, err := s.order.FindByUserId(userId)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return []orderentity.Order{}, fmt.Errorf("Order with user id: %d is not found", userId)
		}
		return []orderentity.Order{}, err
	}
	if len(orders) == 0 {
		return orders, errors.New("Orders not found")
	}
	return orders, nil
}

func (s *OrderService) GetOrderUsingFilter(param orderentity.GetOrderUsingFilterParam) (orderentity.GetOrderUsingFilterResponse, error) {
	var limitPlusOne int32
	if param.Limit != 0 {
		limitPlusOne = param.Limit + 1
	}
	var offset int32
	if param.Page > 0 {
		offset = (param.Page - 1) * param.Limit
	}
	orders, err := s.order.FindByFilter(param.Statuses, param.Invoice, param.StartDate, param.EndDate, param.Sort, limitPlusOne, offset)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return orderentity.GetOrderUsingFilterResponse{}, fmt.Errorf("Order with filter: %+v is not found", param)
		}
		return orderentity.GetOrderUsingFilterResponse{}, err
	}
	if len(orders) == 0 {
		return orderentity.GetOrderUsingFilterResponse{}, errors.New("Orders not found")
	}
	var hasNext bool
	if param.Limit != 0 && param.Page >= 1 && len(orders) > int(param.Limit) {
		orders = orders[:param.Limit]
		hasNext = true
	}
	return orderentity.GetOrderUsingFilterResponse{
		Orders:  orders,
		HasNext: hasNext,
	}, nil
}

func (s *OrderService) NewOrder(order orderentity.Order) (orderentity.Order, error) {
	customer, err := s.customer.FindById(order.CustomerID)
	if err != nil {
		return orderentity.Order{}, err
	}
	order.Customer = customer
	if order.PaymentID == 0 {
		order.PaymentID = 1
	}
	payment, err := s.payment.FindById(order.PaymentID)
	if err != nil {
		return orderentity.Order{}, err
	}
	order.Payment = payment
	now := time.Now()
	order.Invoice = fmt.Sprintf("INV/%s/%d", now.Format("20060201"), now.Unix())
	if s.validateDuplicateOrder(order) {
		return order, errors.New("order duplicated")
	}
	order, err = s.order.New(order)
	if err != nil {
		return orderentity.Order{}, err
	}
	for _, orderDetail := range order.OrderDetails {
		product, err := s.product.FindById(orderDetail.ProductOrder.ProductID)
		if err != nil {
			logger.Log.Errorf("Cannot find product: %v", err)
		}
		// Only update stock when the order is finished
		if order.Status == orderentity.StatusFinish {
			productQty := product.Quantity.Int64
			stockQty := productQty - int64(orderDetail.Quantity)
			product.Quantity = sql.NullInt64{
				Int64: stockQty,
				Valid: true,
			}
			// update product stock
			_, err = s.product.Update(product)
			if err != nil {
				logger.Log.Errorf("Update Product Stock error: %v", err)
			}
		}
	}

	err = memcache.GetInstance().Set(order.GenerateCacheKey(), true, 3)
	if err != nil {
		logger.Log.Errorf("Set order cache error: %v", err)
	}

	return order, nil
}

func (s *OrderService) validateDuplicateOrder(order orderentity.Order) bool {
	key := order.GenerateCacheKey()
	cacheResult, err := memcache.GetInstance().Get(key)
	if err != nil {
		if !errors.Is(err, memcache.ErrCacheMiss) {
			logger.Log.Errorf("Get order cache error: %v", err)
		}
		return false
	}

	result, _ := cacheResult.(bool)
	return result
}

func (s *OrderService) UpdateOrder(order orderentity.Order) (orderentity.Order, error) {
	return s.order.Update(order)
}

func (s *OrderService) UpdateOrderStatus(order orderentity.Order) (orderentity.Order, error) {
	return s.order.UpdateStatus(int64(order.ID), order.Status)
}

func (s *OrderService) DeleteOrder(id int64) (orderentity.Order, error) {
	return s.order.Delete(id)
}
