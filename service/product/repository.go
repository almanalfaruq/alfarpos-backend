package product

import (
	"github.com/almanalfaruq/alfarpos-backend/model"
	productentity "github.com/almanalfaruq/alfarpos-backend/model/product"
	stockentity "github.com/almanalfaruq/alfarpos-backend/model/stock"
)

//go:generate mockgen -source=repository.go -package=product -destination=repository_mock_test.go
type categoryRepositoryIface interface {
	FindAll() ([]model.Category, error)
	FindById(id int64) (model.Category, error)
	FindByName(name string) ([]model.Category, error)
	New(ct model.Category) (model.Category, error)
	Update(ct model.Category) (model.Category, error)
	Delete(id int64) (model.Category, error)
}

type productRepositoryIface interface {
	FindAll() ([]productentity.Product, error)
	FindAllWithLimit(limit, offset int) ([]productentity.Product, error)
	FindById(id int64) (productentity.Product, error)
	FindByIDs(IDs []int64) ([]productentity.Product, error)
	FindByExactCode(code string) (productentity.Product, error)
	GetMultipleProductByExactCode(code string) (productentity.Products, error)
	SearchBy(query string, limit, offset int) ([]productentity.Product, error)
	FindByCode(code string) ([]productentity.Product, error)
	FindByName(name string) ([]productentity.Product, error)
	FindByCategoryName(name string) ([]productentity.Product, error)
	FindByUnitName(name string) ([]productentity.Product, error)
	New(p productentity.Product) (productentity.Product, error)
	Update(p productentity.Product) (productentity.Product, error)
	Delete(id int64) (productentity.Product, error)
	DeleteAll() (int64, error)
}

type stockRepositoryIface interface {
	FindAll() ([]stockentity.Stock, error)
	FindByProduct(product productentity.Product) (stockentity.Stock, error)
	New(st stockentity.Stock) (stockentity.Stock, error)
	Update(st stockentity.Stock) (stockentity.Stock, error)
	Delete(id int64) (stockentity.Stock, error)
	DeleteAll() (int64, error)
}

type unitRepositoryIface interface {
	FindAll() ([]model.Unit, error)
	FindById(id int64) (model.Unit, error)
	FindByName(name string) ([]model.Unit, error)
	New(unit model.Unit) (model.Unit, error)
	Update(unit model.Unit) (model.Unit, error)
	Delete(id int64) (model.Unit, error)
}
