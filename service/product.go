package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"strconv"
	"strings"

	"github.com/kataras/golog"

	"github.com/360EntSecGroup-Skylar/excelize"

	"github.com/almanalfaruq/alfarpos-backend/model"
	"github.com/almanalfaruq/alfarpos-backend/repository"
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
	GetOneProductByCode(code string) (model.Product, error)
	GetProductsByCode(productCode string) ([]model.Product, error)
	GetProductsByName(productName string) ([]model.Product, error)
	GetProductsByCategoryName(categoryName string) ([]model.Product, error)
	GetProductsByUnitName(unitName string) ([]model.Product, error)
	NewProduct(productData string) (model.Product, error)
	NewProductUsingExcel(sheetName string, excelFile multipart.File) error
	UpdateProduct(productData string) (model.Product, error)
	DeleteProduct(id int) (model.Product, error)
}

func (service *ProductService) GetAllProduct() ([]model.Product, error) {
	return service.Product.FindAll(), nil
}

func (service *ProductService) GetOneProduct(id int) (model.Product, error) {
	product := service.Product.FindById(id)
	if product.ID == 0 {
		return product, errors.New("Product not found")
	}
	return product, nil
}

func (service *ProductService) GetOneProductByCode(code string) (model.Product, error) {
	var product model.Product
	products := service.Product.FindByCode(code)
	if len(products) == 0 {
		return product, errors.New("Product not found")
	}
	product = products[0]
	return product, nil
}

func (service *ProductService) GetProductsByCode(productCode string) ([]model.Product, error) {
	products := service.Product.FindByCode(productCode)
	if len(products) == 0 {
		return products, errors.New("Products not found")
	}
	return products, nil
}

func (service *ProductService) GetProductsByName(productName string) ([]model.Product, error) {
	productName = strings.ToLower(productName)
	products := service.Product.FindByName(productName)
	if len(products) == 0 {
		return products, errors.New("Products not found")
	}
	return products, nil
}

func (service *ProductService) GetProductsByCategoryName(categoryName string) ([]model.Product, error) {
	categoryName = strings.ToLower(categoryName)
	products := service.Product.FindByCategoryName(categoryName)
	if len(products) == 0 {
		return products, errors.New("Products not found")
	}
	return products, nil
}

func (service *ProductService) GetProductsByUnitName(unitName string) ([]model.Product, error) {
	unitName = strings.ToLower(unitName)
	products := service.Product.FindByUnitName(unitName)
	if len(products) == 0 {
		return products, errors.New("Products not found")
	}
	return products, nil
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

func (service *ProductService) NewProductUsingExcel(sheetName string, excelFile io.Reader) error {
	golog.Info("Starting excel import...")
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
		if code == "" || name == "" {
			continue
		}
		golog.Infof("Product Name: %s\nQuantity: %d\nSell Price: %d\nBuy Price: %d\n\n", name, quantity, sellPrice, buyPrice)
		product := model.Product{
			Code:      code,
			Name:      name,
			SellPrice: &sellPrice,
			Quantity:  &quantity,
			Category: model.Category{
				Name: categoryName,
			},
			BuyPrice: &buyPrice,
			Unit: model.Unit{
				Name: unitName,
			},
		}
		categories := service.Category.FindByName(product.Category.Name)
		var category model.Category
		if len(categories) == 0 {
			category = model.Category{Name: product.Category.Name}
			category = service.Category.New(category)
		} else {
			category = categories[0]
		}
		product.Category.ID = category.ID
		units := service.Unit.FindByName(product.Unit.Name)
		var unit model.Unit
		if len(units) == 0 {
			unit = model.Unit{Name: product.Unit.Name}
			unit = service.Unit.New(unit)
		} else {
			unit = units[0]
		}
		product.Unit.ID = unit.ID
		products := service.Product.FindByCode(product.Code)
		if len(products) == 0 {
			product = service.Product.New(product)
			golog.Infof("%#v created!", product)
		} else {
			product.ID = products[0].ID
			product = service.Product.Update(product)
			golog.Infof("%#v updated!", product)
		}
		stock := model.Stock{
			ProductID: int(product.ID),
			Product:   product,
			Quantity:  0,
		}
		service.Stock.New(stock)
		productCounter++
	}
	if productCounter != len(rows)-1 {
		warnText := fmt.Sprintf("There are %v rows, but only %v products were created", len(rows)-1, productCounter)
		golog.Warn(warnText)
	}
	golog.Infof("%v products imported!", productCounter)
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

func (service *ProductService) DeleteProduct(id int) (model.Product, error) {
	return service.Product.Delete(id)
}
