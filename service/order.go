package service

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/almanalfaruq/alfarpos-backend/model"
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

func (s *OrderService) GetAllOrder() ([]model.Order, error) {
	return s.order.FindAll(), nil
}

func (s *OrderService) GetOneOrder(id int64) (model.Order, error) {
	order, err := s.order.FindById(id)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return model.Order{}, fmt.Errorf("Order with id: %d is not found", id)
		}
		return model.Order{}, err
	}
	return order, nil
}

func (s *OrderService) GetOrderByInvoice(invoice string) (model.Order, error) {
	invoice = strings.ToLower(invoice)
	order, err := s.order.FindByInvoice(invoice)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return model.Order{}, fmt.Errorf("Order with invoice: %s is not found", invoice)
		}
		return model.Order{}, err
	}
	return order, nil
}

func (s *OrderService) GetOrderByUserId(userId int64) ([]model.Order, error) {
	orders, err := s.order.FindByUserId(userId)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return []model.Order{}, fmt.Errorf("Order with user id: %d is not found", userId)
		}
		return []model.Order{}, err
	}
	if len(orders) == 0 {
		return orders, errors.New("Orders not found")
	}
	return orders, nil
}

func (s *OrderService) NewOrder(order model.Order) (model.Order, error) {
	customer := s.customer.FindById(order.CustomerID)
	order.Customer = customer
	payment := s.payment.FindById(order.PaymentID)
	order.Payment = payment
	now := time.Now()
	order.Invoice = fmt.Sprintf("INV/%s/%d", now.Format("20060201"), now.Unix())
	order, err := s.order.New(order)
	if err != nil {
		return model.Order{}, err
	}
	for _, orderDetail := range order.OrderDetails {
		orderDetail.OrderID = int64(order.ID)
		orderDetail.Order = order
		orderDetail.Product, err = s.product.FindById(orderDetail.ProductID)
		if err != nil {
			return model.Order{}, err
		}
		_, err := s.orderDetail.New(orderDetail)
		if err != nil {
			return model.Order{}, err
		}
		product := orderDetail.Product
		productQty := product.Quantity.Int32
		stockQty := productQty - orderDetail.Quantity
		product.Quantity = sql.NullInt32{
			Int32: stockQty,
			Valid: true,
		}
		// update product stock
		s.product.Update(product)
	}
	return order, nil
}

func (s *OrderService) UpdateOrder(order model.Order) (model.Order, error) {
	return s.order.Update(order)
}

func (s *OrderService) DeleteOrder(id int64) (model.Order, error) {
	return s.order.Delete(id)
}
