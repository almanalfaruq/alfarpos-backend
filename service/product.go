package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"mime/multipart"
	"strconv"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"

	"../model"
	"../repository"
)

type ProductService struct {
	Product  repository.IProductRepository
	Category repository.ICategoryRepository
	Unit     repository.IUnitRepository
	Stock    repository.IStockRepository
}

type IProductService interface {
	GetAllProduct() ([]model.Product, error)
	GetOneProduct(id int) (model.Product, error)
	GetProductsByName(productName string) ([]model.Product, error)
	GetProductsByCategoryName(categoryName string) ([]model.Product, error)
	GetProductsByUnitName(unitName string) ([]model.Product, error)
	NewProduct(productData string) (model.Product, error)
	NewProductUsingExcel(sheetName string, excelFile multipart.File) error
	UpdateProduct(productData string) (model.Product, error)
	DeleteProduct(id int) model.Product
}

func (service *ProductService) GetAllProduct() ([]model.Product, error) {
	return service.Product.FindAll(), nil
}

func (service *ProductService) GetOneProduct(id int) (model.Product, error) {
	return service.Product.FindById(id), nil
}

func (service *ProductService) GetProductsByName(productName string) ([]model.Product, error) {
	productName = strings.ToLower(productName)
	return service.Product.FindByName(productName), nil
}

func (service *ProductService) GetProductsByCategoryName(categoryName string) ([]model.Product, error) {
	return service.Product.FindByCategoryName(categoryName), nil
}

func (service *ProductService) GetProductsByUnitName(unitName string) ([]model.Product, error) {
	return service.Product.FindByUnitName(unitName), nil
}

func (service *ProductService) NewProduct(productData string) (model.Product, error) {
	var product model.Product
	productDataByte := []byte(productData)
	err := json.Unmarshal(productDataByte, &product)
	if err != nil {
		return product, err
	}
	product = service.Product.New(product)
	stock := model.Stock{
		ProductID: int(product.ID),
		Product:   product,
		Quantity:  0,
	}
	service.Stock.New(stock)
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
		sellPrice, _ := strconv.ParseInt(row[2], 10, 64)
		quantity, _ := strconv.ParseInt(row[3], 10, 64)
		categoryName := row[4]
		buyPrice, _ := strconv.ParseInt(row[5], 10, 64)
		unitName := row[6]
		product := model.Product{
			Code:      code,
			Name:      name,
			SellPrice: sellPrice,
			Quantity:  quantity,
			Category: model.Category{
				Name: categoryName,
			},
			BuyPrice: buyPrice,
			Unit: model.Unit{
				Name: unitName,
			},
		}
		category := service.Category.FindByName(product.Category.Name)
		if category.ID == 0 {
			category = model.Category{Name: product.Category.Name}
			category = service.Category.New(category)
		}
		product.Category.ID = category.ID
		unit := service.Unit.FindByName(product.Unit.Name)
		if unit.ID == 0 {
			unit = model.Unit{Name: product.Unit.Name}
			unit = service.Unit.New(unit)
		}
		product.Unit.ID = unit.ID
		product = service.Product.New(product)
		stock := model.Stock{
			ProductID: int(product.ID),
			Product:   product,
			Quantity:  0,
		}
		service.Stock.New(stock)
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
	product = service.Product.Update(product)
	return product, nil
}

func (service *ProductService) DeleteProduct(id int) model.Product {
	return service.Product.Delete(id)
}
