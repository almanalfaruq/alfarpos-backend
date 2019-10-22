package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"mime/multipart"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"

	"../model"
	"../repository"
)

type ProductService struct {
	product  repository.IProductRepository
	category repository.ICategoryRepository
	unit     repository.IUnitRepository
	stock    repository.IStockRepository
}

type IProductService interface {
	GetAllProduct() []model.Product
	GetOneProduct(id int) (model.Product, error)
	GetProductsByName(name string) ([]model.Product, error)
	GetProductsByCategoryName(categoryName string) ([]model.Product, error)
	GetProductsByUnitName(unitName string) ([]model.Product, error)
	NewProduct(productData string) (model.Product, error)
	NewProductUsingExcel(sheetName string, excelFile multipart.File) error
	UpdateProduct(productData string) (model.Product, error)
	DeleteProduct(id int) int
}

func (service *ProductService) GetAllProduct() []model.Product {
	return service.product.FindAll()
}

func (service *ProductService) GetOneProduct(id int) (model.Product, error) {
	return service.product.FindById(id), nil
}

func (service *ProductService) GetProductsByName(productData string) ([]model.Product, error) {
	var product model.Product
	productDataByte := []byte(productData)
	err := json.Unmarshal(productDataByte, &product)
	if err != nil {
		return []model.Product{}, err
	}
	return service.product.FindByName(product.Name), nil
}

func (service *ProductService) GetProductsByCategoryName(productData string) ([]model.Product, error) {
	var product model.Product
	productDataByte := []byte(productData)
	err := json.Unmarshal(productDataByte, &product)
	if err != nil {
		return []model.Product{}, err
	}
	return service.product.FindByCategoryName(product.Category.Name), nil
}

func (service *ProductService) GetProductsByUnitName(productData string) ([]model.Product, error) {
	var product model.Product
	productDataByte := []byte(productData)
	err := json.Unmarshal(productDataByte, &product)
	if err != nil {
		return []model.Product{}, err
	}
	return service.product.FindByUnitName(product.Unit.Name), nil
}

func (service *ProductService) NewProduct(productData string) (model.Product, error) {
	var product model.Product
	productDataByte := []byte(productData)
	err := json.Unmarshal(productDataByte, &product)
	if err != nil {
		return product, err
	}
	product = service.product.New(product)
	stock := model.Stock{
		ProductID: int(product.ID),
		Product:   product,
		Quantity:  0,
	}
	service.stock.New(stock)
	return product, nil
}

func (service *ProductService) NewProductUsingExcel(sheetName string, excelFile multipart.File) error {
	excel, err := excelize.OpenReader(excelFile)
	if err != nil {
		return err
	}
	if sheetName == "" {
		sheetName = "Sheet1"
	}
	rows, err := excel.GetRows(sheetName)
	if err != nil {
		return err
	}
	productCounter := 0
	for indexRow, row := range rows {
		if indexRow == 0 {
			continue
		}
		code := row[0]
		name := row[1]
		sellPrice, _ := strconv.ParseInt(row[2], 10, 32)
		quantity, _ := strconv.ParseInt(row[3], 10, 32)
		categoryName := row[4]
		buyPrice, _ := strconv.ParseInt(row[5], 10, 32)
		unitName := row[6]
		product := model.Product{
			Code:      code,
			Name:      name,
			SellPrice: int(sellPrice),
			Quantity:  int(quantity),
			Category: model.Category{
				Name: categoryName,
			},
			BuyPrice: int(buyPrice),
			Unit: model.Unit{
				Name: unitName,
			},
		}
		category := service.category.FindByName(product.Category.Name)
		if category.ID == 0 {
			category = model.Category{Name: product.Category.Name}
			category = service.category.New(category)
		}
		product.Category.ID = category.ID
		unit := service.unit.FindByName(product.Category.Name)
		if unit.ID == 0 {
			unit = model.Unit{Name: product.Category.Name}
			unit = service.unit.New(unit)
		}
		product.Unit.ID = unit.ID
		product = service.product.New(product)
		stock := model.Stock{
			ProductID: int(product.ID),
			Product:   product,
			Quantity:  0,
		}
		service.stock.New(stock)
		productCounter++
	}
	if productCounter != len(rows)-1 {
		errorText := fmt.Sprintf("There are %v products, but only %v products were created", len(rows)-1, productCounter)
		return errors.New(errorText)
	}
	return nil
}

func (service *ProductService) UpdateProduct(productData string) (model.Product, error) {
	var product model.Product
	productDataByte := []byte(productData)
	err := json.Unmarshal(productDataByte, &product)
	if err != nil {
		return product, err
	}
	product = service.product.Update(product)
	return product, nil
}

func (service *ProductService) DeleteProduct(id int) (model.Product, error) {
	return service.product.Delete(id), nil
}
