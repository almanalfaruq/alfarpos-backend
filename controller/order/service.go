package order

import orderentity "github.com/almanalfaruq/alfarpos-backend/model/order"

type orderServiceIface interface {
	GetAllOrder() ([]orderentity.Order, error)
	GetOneOrder(id int64) (orderentity.Order, error)
	GetOrderByInvoice(invoice string) (orderentity.Order, error)
	GetOrderByUserId(userId int64) ([]orderentity.Order, error)
	GetOrderUsingFilter(param orderentity.GetOrderUsingFilterParam) ([]orderentity.Order, error)
	NewOrder(orderData orderentity.Order) (orderentity.Order, error)
	UpdateOrder(orderData orderentity.Order) (orderentity.Order, error)
	UpdateOrderStatus(order orderentity.Order) (orderentity.Order, error)
	DeleteOrder(id int64) (orderentity.Order, error)
}
