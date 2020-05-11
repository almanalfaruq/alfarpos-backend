package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/kataras/golog"

	"github.com/360EntSecGroup-Skylar/excelize/v2"

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

func (service *ProductService) GetOneProduct(id int64) (model.Product, error) {
	product := service.product.FindById(id)
	if product.ID == 0 {
		return product, errors.New("Product not found")
	}
	return product, nil
}

func (service *ProductService) GetOneProductByCode(code string) (model.Product, error) {
	code = strings.ToLower(code)
	product := service.product.FindByExactCode(code)
	if product.ID == 0 {
		return model.Product{}, errors.New("Product not found")
	}
	return product, nil
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
		ProductID: int64(product.ID),
		Product:   product,
		Quantity:  0,
	}
	service.stock.New(stock)
	return product, nil
}

func (s *ProductService) ExportAllProductsToExcel() (*excelize.File, error) {
	products := s.product.FindAll()
	sheetName := "Products"
	xlsx := excelize.NewFile()

	xlsx.SetSheetName(xlsx.GetSheetName(1), sheetName)
	xlsx.SetCellValue(sheetName, "A1", "Barcode")
	xlsx.SetCellValue(sheetName, "B1", "Nama Barang")
	xlsx.SetCellValue(sheetName, "C1", "Harga Jual")
	xlsx.SetCellValue(sheetName, "D1", "Stok")
	xlsx.SetCellValue(sheetName, "E1", "Jenis Barang")
	xlsx.SetCellValue(sheetName, "F1", "Harga Beli")
	xlsx.SetCellValue(sheetName, "G1", "Satuan")

	for i, product := range products {
		xlsx.SetCellValue(sheetName, fmt.Sprintf("A%d", i+2), product.Code)
		xlsx.SetCellValue(sheetName, fmt.Sprintf("B%d", i+2), product.Name)
		xlsx.SetCellValue(sheetName, fmt.Sprintf("C%d", i+2), *product.SellPrice)
		xlsx.SetCellValue(sheetName, fmt.Sprintf("D%d", i+2), *product.Quantity)
		xlsx.SetCellValue(sheetName, fmt.Sprintf("E%d", i+2), product.Category.Name)
		xlsx.SetCellValue(sheetName, fmt.Sprintf("F%d", i+2), *product.BuyPrice)
		xlsx.SetCellValue(sheetName, fmt.Sprintf("G%d", i+2), product.Unit.Name)
	}

	err := xlsx.SaveAs("./exported-product.xlsx")
	if err != nil {
		return nil, err
	}
	return xlsx, nil
}

func (s *ProductService) NewProductUsingExcel(sheetName string, excelFile io.Reader) (int, error) {
	golog.Info("Starting excel import...")
	excel, err := excelize.OpenReader(excelFile)
	if err != nil {
		return 0, err
	}
	if sheetName == "" {
		sheetName = "Sheet1"
	}
	rows, err := excel.GetRows(sheetName)
	if err != nil {
		return 0, err
	}
	if len(rows) < 1 {
		return 0, fmt.Errorf("Rows length < 1")
	}
	products, errIndex := s.parseExcelRowsToProduct(rows)
	var (
		productCounter = 0
	)
	go func() {
		for _, product := range products {
			golog.Infof("Product Name: %s\nQuantity: %d\nSell Price: %d\nBuy Price: %d\n\n", product.Name, *product.Quantity, *product.SellPrice, *product.BuyPrice)
			categories := s.category.FindByName(product.Category.Name)
			var category model.Category
			if len(categories) == 0 {
				category = model.Category{Name: product.Category.Name}
				category, err = s.category.New(category)
				if err != nil {
					errIndex = append(errIndex, product.Name)
					continue
				}
			} else {
				category = categories[0]
			}
			product.CategoryID = int64(category.ID)
			product.Category.ID = category.ID
			units := s.unit.FindByName(product.Unit.Name)
			var unit model.Unit
			if len(units) == 0 {
				unit = model.Unit{Name: product.Unit.Name}
				unit = s.unit.New(unit)
			} else {
				unit = units[0]
			}
			product.UnitID = int64(unit.ID)
			product.Unit.ID = unit.ID
			oldProduct := s.product.FindByExactCode(product.Code)
			if product.ID == 0 {
				product = s.product.New(product)
				golog.Infof("%#v created!", product)
			} else {
				product.ID = oldProduct.ID
				product = s.product.Update(product)
				golog.Infof("%#v updated!", product)
			}
			stock := model.Stock{
				ProductID: int64(product.ID),
				Product:   product,
				Quantity:  0,
			}
			s.stock.New(stock)
			productCounter++
		}
		if productCounter != len(rows)-1 {
			warnText := fmt.Sprintf("There are %v rows, but only %v products were created.", len(rows)-1, productCounter)
			golog.Warn(warnText)
			warnText = fmt.Sprintf("These are the name of products of unimported rows: %s", strings.Join(errIndex, ", "))
			golog.Warn(warnText)
		}
		golog.Infof("%v products imported!", productCounter)
	}()
	return len(products), nil
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

func (service *ProductService) DeleteProduct(id int64) (model.Product, error) {
	return service.product.Delete(id)
}

func (s *ProductService) parseExcelRowsToProduct(rows [][]string) ([]model.Product, []string) {
	var (
		products []model.Product
		errIndex []string
	)
	// skip index 0 - Header
	for _, row := range rows[1:] {
		code := row[0]
		if code == "" {
			code = uuid.New().String()
		}
		name := row[1]
		sellPrice, err := strconv.ParseInt(row[2], 10, 64)
		if err != nil {
			errIndex = append(errIndex, name)
			continue
		}
		quantity, err := strconv.ParseInt(row[3], 10, 64)
		if err != nil {
			errIndex = append(errIndex, name)
			continue
		}
		categoryName := row[4]
		buyPrice, err := strconv.ParseInt(row[5], 10, 64)
		if err != nil {
			errIndex = append(errIndex, name)
			continue
		}
		unitName := row[6]
		if code == "" || name == "" {
			errIndex = append(errIndex, name)
			continue
		}

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

		products = append(products, product)
	}
	return products, errIndex
}
