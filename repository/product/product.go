package product

import (
	"fmt"
	"strconv"
	"strings"

	productentity "github.com/almanalfaruq/alfarpos-backend/model/product"
	"github.com/almanalfaruq/alfarpos-backend/util"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ProductRepository struct {
	db util.DBIface
}

func NewProductRepo(db util.DBIface) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (repo *ProductRepository) FindAll() ([]productentity.Product, error) {
	var products []productentity.Product
	db := repo.db.GetDb()
	return products, db.Preload(clause.Associations).
		Find(&products).Error
}

func (repo *ProductRepository) FindAllWithLimit(limit, offset int) ([]productentity.Product, error) {
	var products []productentity.Product
	db := repo.db.GetDb()
	return products, db.Preload(clause.Associations).Order("products.name").
		Limit(limit).Offset(offset).Find(&products).Error
}

func (repo *ProductRepository) FindById(id int64) (productentity.Product, error) {
	var product productentity.Product
	db := repo.db.GetDb()
	return product, db.Preload(clause.Associations).
		Where("products.id = ?", id).First(&product).Error
}

func (repo *ProductRepository) FindByIDs(IDs []int64) ([]productentity.Product, error) {
	var (
		products []productentity.Product
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

func (repo *ProductRepository) FindByExactCode(code string) (productentity.Product, error) {
	var product productentity.Product
	db := repo.db.GetDb()
	return product, db.Preload(clause.Associations).
		Where("products.code = ?", code).First(&product).Error
}

func (repo *ProductRepository) GetMultipleProductByExactCode(code string) (productentity.Products, error) {
	var products []productentity.Product
	db := repo.db.GetDb()
	return products, db.Preload(clause.Associations).Where("code = ?", code).Find(&products).Error
}

func (repo *ProductRepository) FindByCode(code string) ([]productentity.Product, error) {
	var products []productentity.Product
	db := repo.db.GetDb()
	return products, db.Preload(clause.Associations).
		Where("products.code ILIKE ?", fmt.Sprintf("%%%s%%", code)).Find(&products).Error
}

func (repo *ProductRepository) SearchBy(query string, limit, offset int) ([]productentity.Product, error) {
	var products []productentity.Product
	db := repo.db.GetDb()
	// TODO: need to index but for now, just leave it using concat_ws func
	return products, db.Preload(clause.Associations).
		Where("concat_ws(' ', products.code, products.name) ILIKE ?", fmt.Sprintf("%%%s%%", query)).
		Order("products.name").Limit(limit).Offset(offset).Find(&products).Error
}

func (repo *ProductRepository) FindByName(name string) ([]productentity.Product, error) {
	var products []productentity.Product
	db := repo.db.GetDb()
	return products, db.Preload(clause.Associations).
		Find(&products).Error
}

func (repo *ProductRepository) FindByCategoryName(name string) ([]productentity.Product, error) {
	var products []productentity.Product
	db := repo.db.GetDb()
	return products, db.Preload(clause.Associations).
		Where("categories.name ILIKE ?", fmt.Sprintf("%%%s%%", name)).Find(&products).Error
}

func (repo *ProductRepository) FindByUnitName(name string) ([]productentity.Product, error) {
	var products []productentity.Product
	db := repo.db.GetDb()
	return products, db.Preload(clause.Associations).
		Where("units.name ILIKE ?", fmt.Sprintf("%%%s%%", name)).Find(&products).Error
}

func (repo *ProductRepository) New(product productentity.Product) (productentity.Product, error) {
	db := repo.db.GetDb()
	err := db.Create(&product).Error
	if err != nil {
		return product, err
	}
	return product, db.Preload(clause.Associations).
		Where("products.id = ?", product.ID).First(&product).Error
}

func (repo *ProductRepository) Update(product productentity.Product) (productentity.Product, error) {
	var oldProduct productentity.Product
	db := repo.db.GetDb()
	existingProductPrices, err := repo.FindProductPricesByProductID(product.ID)
	if err != nil && err != gorm.ErrRecordNotFound {
		return productentity.Product{}, err
	}
	var updatedProductPrice productentity.ProductPrices
	for _, pp := range product.ProductPrices {
		if pp.ID == 0 {
			continue
		}
		updatedProductPrice = append(updatedProductPrice, pp)
	}
	diffExistingPP := productentity.DifferenceProductPrice(updatedProductPrice, existingProductPrices)
	for _, pp := range diffExistingPP {
		err = repo.DeleteProductPrice(pp.ID)
		if err != nil {
			continue
		}
	}
	for _, pp := range updatedProductPrice {
		err = repo.UpdateProductPrice(pp)
		if err != nil {
			continue
		}
	}
	err = db.Where("id = ?", product.ID).First(&oldProduct).Error
	if err != nil {
		return product, err
	}
	oldProduct = product
	err = db.Model(&oldProduct).Updates(product).Error
	if err != nil {
		return product, err
	}
	return product, db.Preload(clause.Associations).
		Where("products.id = ?", product.ID).First(&product).Error
}

func (repo *ProductRepository) Delete(id int64) (productentity.Product, error) {
	var product productentity.Product
	db := repo.db.GetDb()
	db.Where("id = ?", id).First(&product)
	err := db.Delete(&product).Error
	return product, err
}

func (repo *ProductRepository) DeleteAll() (int64, error) {
	var product productentity.Product
	var productCount int64
	db := repo.db.GetDb()
	err := db.Model(&product).Count(&productCount).Error
	if err != nil {
		return 0, err
	}
	return productCount, db.Unscoped().Delete(&product).Error
}

func (repo *ProductRepository) FindProductPricesByProductID(productID int64) (productentity.ProductPrices, error) {
	var productPrices productentity.ProductPrices
	db := repo.db.GetDb()
	return productPrices, db.Where("product_id = ?", productID).Find(&productPrices).Error
}

func (repo *ProductRepository) UpdateProductPrice(productPrice productentity.ProductPrice) error {
	var oldProductPrice productentity.ProductPrice
	db := repo.db.GetDb()
	err := db.Where("id = ?", productPrice.ID).First(&oldProductPrice).Error
	if err != nil {
		return err
	}
	oldProductPrice = productPrice
	return db.Save(&oldProductPrice).Error
}

func (repo *ProductRepository) DeleteProductPrice(id int64) error {
	var productPrice productentity.ProductPrice
	db := repo.db.GetDb()
	db.Where("id = ?", id).First(&productPrice)
	return db.Delete(&productPrice).Error
}
