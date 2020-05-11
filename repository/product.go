package repository

import (
	"fmt"

	"github.com/almanalfaruq/alfarpos-backend/model"
)

type ProductRepository struct {
	db dbIface
}

func NewProductRepo(db dbIface) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (repo *ProductRepository) FindAll() []model.Product {
	var categories []model.Product
	db := repo.db.GetDb()
	db.Set("gorm:auto_preload", true).Find(&categories)
	return categories
}

func (repo *ProductRepository) FindById(id int64) model.Product {
	var product model.Product
	db := repo.db.GetDb()
	db.Set("gorm:auto_preload", true).Where("id = ?", id).First(&product)
	return product
}

func (repo *ProductRepository) FindByExactCode(code string) model.Product {
	var product model.Product
	db := repo.db.GetDb()
	db.Set("gorm:auto_preload", true).Where("LOWER(code) = ?", code).First(&product)
	return product
}

func (repo *ProductRepository) FindByCode(code string) []model.Product {
	var products []model.Product
	db := repo.db.GetDb()
	db.Set("gorm:auto_preload", true).Where("LOWER(code) LIKE ?", fmt.Sprintf("%%%s%%", code)).Find(&products)
	return products
}

func (repo *ProductRepository) FindByName(name string) []model.Product {
	var products []model.Product
	db := repo.db.GetDb()
	db.Set("gorm:auto_preload", true).Where("LOWER(name) LIKE ?", fmt.Sprintf("%%%s%%", name)).Find(&products)
	return products
}

func (repo *ProductRepository) FindByCategoryName(name string) []model.Product {
	var products []model.Product
	db := repo.db.GetDb()
	db.Set("gorm:auto_preload", true).Joins("JOIN categories ON categories.id = products.category_id").Where("LOWER(categories.name) LIKE ?", fmt.Sprintf("%%%s%%", name)).Find(&products)
	return products
}

func (repo *ProductRepository) FindByUnitName(name string) []model.Product {
	var products []model.Product
	db := repo.db.GetDb()
	db.Set("gorm:auto_preload", true).Joins("JOIN units ON units.id = products.unit_id").Where("LOWER(units.name) LIKE ?", fmt.Sprintf("%%%s%%", name)).Find(&products)
	return products
}

func (repo *ProductRepository) New(product model.Product) model.Product {
	db := repo.db.GetDb()
	isNotExist := db.NewRecord(product)
	if isNotExist {
		db.Create(&product)
	}
	db.Set("gorm:auto_preload", true).Where("id = ?", product.ID).First(&product)
	return product
}

func (repo *ProductRepository) Update(product model.Product) model.Product {
	var oldProduct model.Product
	db := repo.db.GetDb()
	db.Where("id = ?", product.ID).First(&oldProduct)
	oldProduct = product
	db.Save(&oldProduct)
	db.Set("gorm:auto_preload", true).Where("id = ?", product.ID).First(&product)
	return product
}

func (repo *ProductRepository) Delete(id int64) (model.Product, error) {
	var product model.Product
	db := repo.db.GetDb()
	db.Where("id = ?", id).First(&product)
	err := db.Delete(&product).Error
	return product, err
}

func (repo *ProductRepository) DeleteAll() int64 {
	var product model.Product
	var productCount int64
	db := repo.db.GetDb()
	db.Model(&product).Count(&productCount)
	db.Unscoped().Delete(&product)
	return productCount
}
