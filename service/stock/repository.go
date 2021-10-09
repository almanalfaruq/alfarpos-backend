package stock

import (
	productentity "github.com/almanalfaruq/alfarpos-backend/model/product"
	stockentity "github.com/almanalfaruq/alfarpos-backend/model/stock"
)

//go:generate mockgen -source=repository.go -package=stock -destination=repository_mock_test.go
type stockRepositoryIface interface {
	FindAll() ([]stockentity.Stock, error)
	FindByProduct(prd productentity.Product) (stockentity.Stock, error)
	New(stck stockentity.Stock) (stockentity.Stock, error)
	Update(stck stockentity.Stock) (stockentity.Stock, error)
	Delete(id int64) (stockentity.Stock, error)
	DeleteAll() (int64, error)
}

type productRepositoryIface interface {
	FindAll() ([]productentity.Product, error)
	FindAllWithLimit(limit, offset int) ([]productentity.Product, error)
	FindById(id int64) (productentity.Product, error)
	FindByIDs(IDs []int64) ([]productentity.Product, error)
	FindByExactCode(code string) (productentity.Product, error)
	GetMultipleProductByExactCode(code string) ([]productentity.Product, error)
	SearchBy(query string, limit, offset int) ([]productentity.Product, error)
	FindByCode(code string) ([]productentity.Product, error)
	FindByName(name string) ([]productentity.Product, error)
	FindByCategoryName(name string) ([]productentity.Product, error)
	FindByUnitName(name string) ([]productentity.Product, error)
	New(prd productentity.Product) (productentity.Product, error)
	Update(prd productentity.Product) (productentity.Product, error)
	Delete(id int64) (productentity.Product, error)
	DeleteAll() (int64, error)
}
