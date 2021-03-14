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

func (repo *PaymentRepository) FindAll() ([]model.Payment, error) {
	var categories []model.Payment
	db := repo.db.GetDb()
	return categories, db.Find(&categories).Error
}

func (repo *PaymentRepository) FindById(id int64) (model.Payment, error) {
	var payment model.Payment
	db := repo.db.GetDb()
	return payment, db.Where("id = ?", id).First(&payment).Error
}

func (repo *PaymentRepository) FindByName(name string) ([]model.Payment, error) {
	var payments []model.Payment
	db := repo.db.GetDb()
	return payments, db.Where("LOWER(name) LIKE ?", fmt.Sprintf("%%%s%%", name)).Find(&payments).Error
}

func (repo *PaymentRepository) New(payment model.Payment) (model.Payment, error) {
	db := repo.db.GetDb()
	return payment, db.Create(&payment).Error
}

func (repo *PaymentRepository) Update(payment model.Payment) (model.Payment, error) {
	var oldPayment model.Payment
	db := repo.db.GetDb()
	err := db.Where("id = ?", payment.ID).First(&oldPayment).Error
	if err != nil {
		return payment, err
	}
	oldPayment = payment
	return payment, db.Save(&oldPayment).Error
}

func (repo *PaymentRepository) Delete(id int64) (model.Payment, error) {
	var payment model.Payment
	db := repo.db.GetDb()
	err := db.Where("id = ?", id).First(&payment).Error
	if err != nil {
		return payment, err
	}
	return payment, db.Delete(&payment).Error
}
