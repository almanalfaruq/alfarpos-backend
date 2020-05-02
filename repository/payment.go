package repository

import (
	"fmt"

	"github.com/almanalfaruq/alfarpos-backend/model"
)

type PaymentRepository struct {
	db dbIface
}

func NewPaymentRepo(db dbIface) *PaymentRepository {
	return &PaymentRepository{
		db: db,
	}
}

func (repo *PaymentRepository) FindAll() []model.Payment {
	var categories []model.Payment
	db := repo.db.GetDb()
	db.Find(&categories)
	return categories
}

func (repo *PaymentRepository) FindById(id int64) model.Payment {
	var payment model.Payment
	db := repo.db.GetDb()
	db.Where("id = ?", id).First(&payment)
	return payment
}

func (repo *PaymentRepository) FindByName(name string) []model.Payment {
	var payments []model.Payment
	db := repo.db.GetDb()
	db.Where("LOWER(name) LIKE ?", fmt.Sprintf("%%%s%%", name)).Find(&payments)
	return payments
}

func (repo *PaymentRepository) New(payment model.Payment) model.Payment {
	db := repo.db.GetDb()
	isNotExist := db.NewRecord(payment)
	if isNotExist {
		db.Create(&payment)
	}
	return payment
}

func (repo *PaymentRepository) Update(payment model.Payment) model.Payment {
	var oldPayment model.Payment
	db := repo.db.GetDb()
	db.Where("id = ?", payment.ID).First(&oldPayment)
	oldPayment = payment
	db.Save(&oldPayment)
	return payment
}

func (repo *PaymentRepository) Delete(id int64) (model.Payment, error) {
	var payment model.Payment
	db := repo.db.GetDb()
	db.Where("id = ?", id).First(&payment)
	err := db.Delete(&payment).Error
	return payment, err
}
