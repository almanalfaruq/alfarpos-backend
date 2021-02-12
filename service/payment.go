package service

import (
	"encoding/json"
	"strings"

	"github.com/almanalfaruq/alfarpos-backend/model"
)

type PaymentService struct {
	payment paymentRepositoryIface
}

func NewPaymentService(paymentRepo paymentRepositoryIface) *PaymentService {
	return &PaymentService{
		payment: paymentRepo,
	}
}

func (service *PaymentService) GetAllPayment() ([]model.Payment, error) {
	return service.payment.FindAll()
}

func (service *PaymentService) GetOnePayment(id int64) (model.Payment, error) {
	return service.payment.FindById(id)
}

func (service *PaymentService) GetPaymentsByName(name string) ([]model.Payment, error) {
	name = strings.ToLower(name)
	return service.payment.FindByName(name)
}

func (service *PaymentService) NewPayment(paymentData string) (model.Payment, error) {
	var payment model.Payment
	paymentDataByte := []byte(paymentData)
	err := json.Unmarshal(paymentDataByte, &payment)
	if err != nil {
		return payment, err
	}
	return service.payment.New(payment)
}

func (service *PaymentService) UpdatePayment(paymentData string) (model.Payment, error) {
	var payment model.Payment
	paymentDataByte := []byte(paymentData)
	err := json.Unmarshal(paymentDataByte, &payment)
	if err != nil {
		return payment, err
	}
	return service.payment.Update(payment)
}

func (service *PaymentService) DeletePayment(id int64) (model.Payment, error) {
	return service.payment.Delete(id)
}
