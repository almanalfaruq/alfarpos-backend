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

func (repo *ProductRepository) FindAll() ([]model.Product, error) {
	var categories []model.Product
	db := repo.db.GetDb()
	return categories, db.Set("gorm:auto_preload", true).Find(&categories).Error
}

func (repo *ProductRepository) FindById(id int64) (model.Product, error) {
	var product model.Product
	db := repo.db.GetDb()
	return product, db.Set("gorm:auto_preload", true).Where("id = ?", id).First(&product).Error
}

func (repo *ProductRepository) FindByExactCode(code string) (model.Product, error) {
	var product model.Product
	db := repo.db.GetDb()
	return product, db.Set("gorm:auto_preload", true).Where("LOWER(code) = ?", code).First(&product).Error
}

func (repo *ProductRepository) FindByCode(code string) ([]model.Product, error) {
	var products []model.Product
	db := repo.db.GetDb()
	return products, db.Set("gorm:auto_preload", true).Where("LOWER(code) LIKE ?", fmt.Sprintf("%%%s%%", code)).Find(&products).Error
}

func (repo *ProductRepository) FindByName(name string) ([]model.Product, error) {
	var products []model.Product
	db := repo.db.GetDb()
	return products, db.Set("gorm:auto_preload", true).Where("LOWER(name) LIKE ?", fmt.Sprintf("%%%s%%", name)).Find(&products).Error
}

func (repo *ProductRepository) FindByCategoryName(name string) ([]model.Product, error) {
	var products []model.Product
	db := repo.db.GetDb()
	return products, db.Set("gorm:auto_preload", true).Joins("JOIN categories ON categories.id = products.category_id").
		Where("LOWER(categories.name) LIKE ?", fmt.Sprintf("%%%s%%", name)).Find(&products).Error
}

func (repo *ProductRepository) FindByUnitName(name string) ([]model.Product, error) {
	var products []model.Product
	db := repo.db.GetDb()
	return products, db.Set("gorm:auto_preload", true).Joins("JOIN units ON units.id = products.unit_id").
		Where("LOWER(units.name) LIKE ?", fmt.Sprintf("%%%s%%", name)).Find(&products).Error
}

func (repo *ProductRepository) New(product model.Product) (model.Product, error) {
	db := repo.db.GetDb()
	isNotExist := db.NewRecord(product)
	if isNotExist {
		err := db.Create(&product).Error
		if err != nil {
			return product, err
		}
	}
	return product, db.Set("gorm:auto_preload", true).Where("id = ?", product.ID).First(&product).Error
}

func (repo *ProductRepository) Update(product model.Product) (model.Product, error) {
	var oldProduct model.Product
	db := repo.db.GetDb()
	err := db.Where("id = ?", product.ID).First(&oldProduct).Error
	if err != nil {
		return product, err
	}
	oldProduct = product
	err = db.Save(&oldProduct).Error
	if err != nil {
		return product, err
	}
	return product, db.Set("gorm:auto_preload", true).Where("id = ?", product.ID).First(&product).Error
}

func (repo *ProductRepository) Delete(id int64) (model.Product, error) {
	var product model.Product
	db := repo.db.GetDb()
	db.Where("id = ?", id).First(&product)
	err := db.Delete(&product).Error
	return product, err
}

func (repo *ProductRepository) DeleteAll() (int64, error) {
	var product model.Product
	var productCount int64
	db := repo.db.GetDb()
	err := db.Model(&product).Count(&productCount).Error
	if err != nil {
		return 0, err
	}
	return productCount, db.Unscoped().Delete(&product).Error
}
