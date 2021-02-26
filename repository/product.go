package repository

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/almanalfaruq/alfarpos-backend/model"
	"gorm.io/gorm/clause"
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
	var products []model.Product
	db := repo.db.GetDb()
	return products, db.Preload(clause.Associations).
		Find(&products).Error
}

func (repo *ProductRepository) FindAllWithLimit(limit, offset int) ([]model.Product, error) {
	var products []model.Product
	db := repo.db.GetDb()
	return products, db.Preload(clause.Associations).
		Find(&products).Limit(limit).Offset(offset).Error
}

func (repo *ProductRepository) FindById(id int64) (model.Product, error) {
	var product model.Product
	db := repo.db.GetDb()
	return product, db.Preload(clause.Associations).
		Where("products.id = ?", id).First(&product).Error
}

func (repo *ProductRepository) FindByIDs(IDs []int64) ([]model.Product, error) {
	var (
		products []model.Product
		strIDs   []string
	)
	for _, id := range IDs {
		strID := strconv.FormatInt(id, 10)
		strIDs = append(strIDs, strID)
	}
	db := repo.db.GetDb()
	return products, db.Preload(clause.Associations).
		Where(fmt.Sprintf("products.id IN (%s)", strings.Join(strIDs, ","))).Find(&products).Error
}

func (repo *ProductRepository) FindByExactCode(code string) (model.Product, error) {
	var product model.Product
	db := repo.db.GetDb()
	return product, db.Preload(clause.Associations).
		Where("products.code = ?", code).First(&product).Error
}

func (repo *ProductRepository) GetMultipleProductByExactCode(code string) ([]model.Product, error) {
	var products []model.Product
	db := repo.db.GetDb()
	return products, db.Where("code = ?", code).Find(&products).Error
}

func (repo *ProductRepository) FindByCode(code string) ([]model.Product, error) {
	var products []model.Product
	db := repo.db.GetDb()
	return products, db.Preload(clause.Associations).
		Where("products.code ILIKE ?", fmt.Sprintf("%%%s%%", code)).Find(&products).Error
}

func (repo *ProductRepository) SearchBy(query string, limit, offset int) ([]model.Product, error) {
	var products []model.Product
	db := repo.db.GetDb()
	// TODO: need to index but for now, just leave it using concat_ws func
	return products, db.Preload(clause.Associations).
		Where("concat_ws(' ', products.code, products.name) ILIKE ?", fmt.Sprintf("%%%s%%", query)).
		Limit(limit).Offset(offset).Find(&products).Error
}

func (repo *ProductRepository) FindByName(name string) ([]model.Product, error) {
	var products []model.Product
	db := repo.db.GetDb()
	return products, db.Preload(clause.Associations).
		Find(&products).Error
}

func (repo *ProductRepository) FindByCategoryName(name string) ([]model.Product, error) {
	var products []model.Product
	db := repo.db.GetDb()
	return products, db.Preload(clause.Associations).
		Where("categories.name ILIKE ?", fmt.Sprintf("%%%s%%", name)).Find(&products).Error
}

func (repo *ProductRepository) FindByUnitName(name string) ([]model.Product, error) {
	var products []model.Product
	db := repo.db.GetDb()
	return products, db.Preload(clause.Associations).
		Where("units.name ILIKE ?", fmt.Sprintf("%%%s%%", name)).Find(&products).Error
}

func (repo *ProductRepository) New(product model.Product) (model.Product, error) {
	db := repo.db.GetDb()
	err := db.Create(&product).Error
	if err != nil {
		return product, err
	}
	return product, db.Preload(clause.Associations).
		Where("products.id = ?", product.ID).First(&product).Error
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
	return product, db.Preload(clause.Associations).
		Where("products.id = ?", product.ID).First(&product).Error
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
