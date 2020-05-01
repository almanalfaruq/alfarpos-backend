package service

import (
	"encoding/json"

	"github.com/almanalfaruq/alfarpos-backend/model"
)

type OrderDetailService struct {
	order       orderRepositoryIface
	orderDetail orderDetailRepositoryIface
}

func NewOrderDetailService(orderRepo orderRepositoryIface, orderDetailRepo orderDetailRepositoryIface) *OrderDetailService {
	return &OrderDetailService{
		order:       orderRepo,
		orderDetail: orderDetailRepo,
	}
}

func (service *OrderDetailService) GetOrderDetailByOrder(orderDetailData string) ([]model.OrderDetail, error) {
	var order model.Order
	orderDetailDataByte := []byte(orderDetailData)
	err := json.Unmarshal(orderDetailDataByte, &order)
	if err != nil {
		return []model.OrderDetail{}, err
	}
	if order.Invoice != "" {
		order = service.order.FindByInvoice(order.Invoice)
	} else {
		order = service.order.FindById(int(order.ID))
	}
	return service.orderDetail.FindByOrder(order), nil
}

func (service *OrderDetailService) DeleteOrderDetail(id int) (model.OrderDetail, error) {
	return service.orderDetail.Delete(id), nil
}

func (service *OrderDetailService) DeleteOrderDetailByOrderId(id int) (int, error) {
	return service.orderDetail.DeleteByOrderId(id), nil
}
