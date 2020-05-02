package service

import (
	"encoding/json"
	"errors"
	"fmt"

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
		order, err = service.order.FindByInvoice(order.Invoice)
		if err != nil {
			if errors.Is(err, model.ErrNotFound) {
				return []model.OrderDetail{}, fmt.Errorf("Order with invoice: %s is not found", order.Invoice)
			}
			return []model.OrderDetail{}, err
		}
	} else {
		order, err = service.order.FindById(int64(order.ID))
		if err != nil {
			if errors.Is(err, model.ErrNotFound) {
				return []model.OrderDetail{}, fmt.Errorf("Order with id: %d is not found", order.ID)
			}
			return []model.OrderDetail{}, err
		}
	}
	return service.orderDetail.FindByOrder(order)
}

func (service *OrderDetailService) DeleteOrderDetail(id int64) (model.OrderDetail, error) {
	return service.orderDetail.Delete(id)
}

func (service *OrderDetailService) DeleteOrderDetailByOrderId(id int64) (int64, error) {
	return service.orderDetail.DeleteByOrderId(id)
}
