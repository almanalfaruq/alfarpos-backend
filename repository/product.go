package repository

import (
	"../model"
	"../util"
)

type ProductRepository struct {
	util.IDatabaseConnection
}

type IProductRepository interface {
	FindAll() []model.Product
	FindById(id int) model.Product
	FindByName(name string) []model.Product
	FindByCategoryName(name string) []model.Product
	FindByUnitName(name string) []model.Product
	New(product model.Product) model.Product
	Update(product model.Product) model.Product
	Delete(id int) model.Product
	DeleteAll() int
}

func (repo *ProductRepository) FindAll() []model.Product {
	var categories []model.Product
	db := repo.GetDb()
	db.Find(&categories)
	return categories
}

func (repo *ProductRepository) FindById(id int) model.Product {
	var product model.Product
	db := repo.GetDb()
	db.Where("id = ?", id).First(&product)
	return product
}

func (repo *ProductRepository) FindByName(name string) []model.Product {
	var products []model.Product
	db := repo.GetDb()
	db.Where("name = ?", name).Find(&products)
	return products
}

func (repo *ProductRepository) FindByCategoryName(name string) []model.Product {
	var products []model.Product
	db := repo.GetDb()
	db.Joins("JOIN categories ON categories.ID = products.category_id").Where("categories.name = ?", name).Find(&products)
	return products
}

func (repo *ProductRepository) FindByUnitName(name string) []model.Product {
	var products []model.Product
	db := repo.GetDb()
	db.Joins("JOIN units ON units.ID = products.unit_id").Where("units.name = ?", name).Find(&products)
	return products
}

func (repo *ProductRepository) New(product model.Product) model.Product {
	db := repo.GetDb()
	isNotExist := db.NewRecord(product)
	if isNotExist {
		db.Create(&product)
	}
	return product
}

func (repo *ProductRepository) Update(product model.Product) model.Product {
	var oldProduct model.Product
	db := repo.GetDb()
	db.Where("id = ?", product.ID).First(&oldProduct)
	oldProduct = product
	db.Save(&oldProduct)
	return product
}

func (repo *ProductRepository) Delete(id int) model.Product {
	var product model.Product
	db := repo.GetDb()
	db.Where("id = ?", id).First(&product)
	db.Delete(&product)
	return product
}

func (repo *ProductRepository) DeleteAll() int {
	var product model.Product
	var productCount int
	db := repo.GetDb()
	db.Model(&product).Count(&productCount)
	db.Delete(&product)
	return productCount
}
