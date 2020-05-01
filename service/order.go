package service

import (
	"encoding/json"
	"errors"
	"strings"

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

func (service *OrderService) GetAllOrder() ([]model.Order, error) {
	return service.order.FindAll(), nil
}

func (service *OrderService) GetOneOrder(id int) (model.Order, error) {
	Order := service.order.FindById(id)
	if Order.ID == 0 {
		return Order, errors.New("Order not found")
	}
	return Order, nil
}

func (service *OrderService) GetOrderByInvoice(invoice string) (model.Order, error) {
	invoice = strings.ToLower(invoice)
	Order := service.order.FindByInvoice(invoice)
	if Order.ID == 0 {
		return Order, errors.New("Order not found")
	}
	return Order, nil
}

func (service *OrderService) GetOrderByUserId(userId int) ([]model.Order, error) {
	Orders := service.order.FindByUserId(userId)
	if len(Orders) == 0 {
		return Orders, errors.New("Orders not found")
	}
	return Orders, nil
}

func (service *OrderService) NewOrder(OrderData string) (model.Order, error) {
	var Order model.Order
	OrderDataByte := []byte(OrderData)
	err := json.Unmarshal(OrderDataByte, &Order)
	if err != nil {
		return Order, err
	}
	Customer := service.customer.FindById(Order.CustomerID)
	Order.Customer = Customer
	Payment := service.payment.FindById(Order.PaymentID)
	Order.Payment = Payment
	Order = service.order.New(Order)
	for _, OrderDetail := range Order.OrderDetails {
		OrderDetail.OrderID = int(Order.ID)
		OrderDetail.Order = Order
		OrderDetail.Product = service.product.FindById(OrderDetail.ProductID)
		service.orderDetail.New(OrderDetail)
		Product := OrderDetail.Product
		ProductQty := *Product.Quantity
		stockQty := ProductQty - int64(OrderDetail.Quantity)
		Product.Quantity = &stockQty
		service.product.Update(Product)
	}
	return Order, nil
}

func (service *OrderService) UpdateOrder(OrderData string) (model.Order, error) {
	var Order model.Order
	OrderDataByte := []byte(OrderData)
	err := json.Unmarshal(OrderDataByte, &Order)
	if err != nil {
		return Order, err
	}
	Order = service.order.Update(Order)
	return Order, nil
}

func (service *OrderService) DeleteOrder(id int) (model.Order, error) {
	return service.order.Delete(id)
}
