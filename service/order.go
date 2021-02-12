package service

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/almanalfaruq/alfarpos-backend/model"
	orderentity "github.com/almanalfaruq/alfarpos-backend/model/order"
	"github.com/kataras/golog"
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

func (s *OrderService) NewOrder(order orderentity.Order) (orderentity.Order, error) {
	customer, err := s.customer.FindById(order.CustomerID)
	if err != nil {
		return orderentity.Order{}, err
	}
	order.Customer = customer
	payment, err := s.payment.FindById(order.PaymentID)
	if err != nil {
		return orderentity.Order{}, err
	}
	order.Payment = payment
	now := time.Now()
	order.Invoice = fmt.Sprintf("INV/%s/%d", now.Format("20060201"), now.Unix())
	order, err = s.order.New(order)
	if err != nil {
		return orderentity.Order{}, err
	}
	for _, orderDetail := range order.OrderDetails {
		product, err := s.product.FindById(orderDetail.ProductOrder.ProductID)
		if err != nil {
			golog.Errorf("Cannot find product: %v", err)
		}
		productQty := product.Quantity.Int64
		stockQty := productQty - int64(orderDetail.Quantity)
		product.Quantity = sql.NullInt64{
			Int64: stockQty,
			Valid: true,
		}
		// update product stock
		_, err = s.product.Update(product)
		if err != nil {
			golog.Errorf("Update Product Stock error: %v", err)
		}

	}
	return order, nil
}

func (s *OrderService) UpdateOrder(order orderentity.Order) (orderentity.Order, error) {
	return s.order.Update(order)
}

func (s *OrderService) DeleteOrder(id int64) (orderentity.Order, error) {
	return s.order.Delete(id)
}
