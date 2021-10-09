package product

import (
	"context"
	"io"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	productentity "github.com/almanalfaruq/alfarpos-backend/model/product"
)

type productServiceIface interface {
	GetAllProduct(limit, page int) (products []productentity.Product, hasNext bool, err error)
	GetOneProduct(id int64) (productentity.Product, error)
	GetProductsByIDs(IDs []int64) ([]productentity.Product, error)
	GetOneProductByCode(code string) (productentity.Product, error)
	GetProductsBySearchQuery(query string, limit, page int) (products []productentity.Product, hasNext bool, err error)
	GetProductsByCode(productCode string) ([]productentity.Product, error)
	GetProductsByName(productName string) ([]productentity.Product, error)
	GetProductsByCategoryName(categoryName string) ([]productentity.Product, error)
	GetProductsByUnitName(unitName string) ([]productentity.Product, error)
	ExportAllProductsToExcel() (*excelize.File, error)
	NewProduct(productData string) (productentity.Product, error)
	NewProductUsingExcel(sheetName string, excelFile io.Reader) (int, error)
	UpdateProduct(productData string) (productentity.Product, error)
	UpsertWithExcel(ctx context.Context, sheetName string, excelFile io.Reader) error
	DeleteProduct(id int64) (productentity.Product, error)
}
