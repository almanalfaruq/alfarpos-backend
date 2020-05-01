package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/kataras/golog"

	"github.com/360EntSecGroup-Skylar/excelize"

	"github.com/almanalfaruq/alfarpos-backend/model"
)

type ProductService struct {
	product  productRepositoryIface
	category categoryRepositoryIface
	unit     unitRepositoryIface
	stock    stockRepositoryIface
}

func NewProductService(productRepo productRepositoryIface, categoryRepo categoryRepositoryIface, unitRepo unitRepositoryIface,
	stockRepo stockRepositoryIface) *ProductService {
	return &ProductService{
		product:  productRepo,
		category: categoryRepo,
		unit:     unitRepo,
		stock:    stockRepo,
	}
}

func (service *ProductService) GetAllProduct() ([]model.Product, error) {
	return service.product.FindAll(), nil
}

func (service *ProductService) GetOneProduct(id int) (model.Product, error) {
	product := service.product.FindById(id)
	if product.ID == 0 {
		return product, errors.New("Product not found")
	}
	return product, nil
}

func (service *ProductService) GetOneProductByCode(code string) (model.Product, error) {
	products := service.product.FindByCode(code)
	if products[0].ID == 0 {
		return products[0], errors.New("Product not found")
	}
	return products[0], nil
}

func (service *ProductService) GetProductsByCode(productCode string) ([]model.Product, error) {
	products := service.product.FindByCode(productCode)
	if len(products) == 0 {
		return products, errors.New("Products not found")
	}
	return products, nil
}

func (service *ProductService) GetProductsByName(productName string) ([]model.Product, error) {
	productName = strings.ToLower(productName)
	products := service.product.FindByName(productName)
	if len(products) == 0 {
		return products, errors.New("Products not found")
	}
	return products, nil
}

func (service *ProductService) GetProductsByCategoryName(categoryName string) ([]model.Product, error) {
	categoryName = strings.ToLower(categoryName)
	products := service.product.FindByCategoryName(categoryName)
	if len(products) == 0 {
		return products, errors.New("Products not found")
	}
	return products, nil
}

func (service *ProductService) GetProductsByUnitName(unitName string) ([]model.Product, error) {
	unitName = strings.ToLower(unitName)
	products := service.product.FindByUnitName(unitName)
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
	product = service.product.New(product)
	stock := model.Stock{
		ProductID: int(product.ID),
		Product:   product,
		Quantity:  0,
	}
	service.stock.New(stock)
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
		categories := service.category.FindByName(product.Category.Name)
		var category model.Category
		if len(categories) == 0 {
			category = model.Category{Name: product.Category.Name}
			category, err = service.category.New(category)
			if err != nil {
				continue
			}
		} else {
			category = categories[0]
		}
		product.CategoryID = int(category.ID)
		product.Category.ID = category.ID
		units := service.unit.FindByName(product.Unit.Name)
		var unit model.Unit
		if len(units) == 0 {
			unit = model.Unit{Name: product.Unit.Name}
			unit = service.unit.New(unit)
		} else {
			unit = units[0]
		}
		product.UnitID = int(unit.ID)
		product.Unit.ID = unit.ID
		oldProduct := service.product.FindByCode(product.Code)[0]
		if oldProduct.ID == 0 {
			product = service.product.New(product)
			golog.Infof("%#v created!", product)
		} else {
			product.ID = oldProduct.ID
			product = service.product.Update(product)
			golog.Infof("%#v updated!", product)
		}
		stock := model.Stock{
			ProductID: int(product.ID),
			Product:   product,
			Quantity:  0,
		}
		service.stock.New(stock)
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
	product = service.product.Update(product)
	return product, nil
}

func (service *ProductService) DeleteProduct(id int) (model.Product, error) {
	return service.product.Delete(id)
}
