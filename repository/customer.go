package repository

import (
	"github.com/almanalfaruq/alfarpos-backend/model"
)

type CustomerRepository struct {
	db dbIface
}

func NewCustomerRepo(db dbIface) *CustomerRepository {
	return &CustomerRepository{
		db: db,
	}
}

func (repo *CustomerRepository) FindAll() ([]model.Customer, error) {
	var customers []model.Customer
	db := repo.db.GetDb()
	return customers, db.Find(&customers).Error
}

func (repo *CustomerRepository) FindById(id int64) (model.Customer, error) {
	var customer model.Customer
	db := repo.db.GetDb()
	return customer, db.Where("id = ?", id).First(&customer).Error
}

func (repo *CustomerRepository) New(customer model.Customer) (model.Customer, error) {
	db := repo.db.GetDb()
	return customer, db.Create(&customer).Error
}

func (repo *CustomerRepository) Update(customer model.Customer) (model.Customer, error) {
	var oldCustomer model.Customer
	db := repo.db.GetDb()
	db.Where("id = ?", customer.ID).First(&oldCustomer)
	oldCustomer = customer
	return customer, db.Save(&oldCustomer).Error
}

func (repo *CustomerRepository) Delete(id int64) (model.Customer, error) {
	var customer model.Customer
	db := repo.db.GetDb()
	err := db.Where("id = ?", id).First(&customer).Error
	if err != nil {
		return customer, err
	}
	return customer, db.Delete(&customer).Error
}
