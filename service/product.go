package service

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"

	"github.com/kataras/golog"
	"github.com/lib/pq"

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

func (service *ProductService) GetAllProduct(limit, page int) (products []model.Product, hasNext bool, err error) {
	if limit == 0 {
		products, err = service.product.FindAll()
		if err != nil {
			return nil, false, err
		}
		return products, false, nil
	}
	offset := (page - 1) * limit
	limitPlusOne := limit + 1
	products, err = service.product.FindAllWithLimit(limit, offset)
	if err != nil {
		return nil, false, err
	}

	if len(products) == limitPlusOne {
		products = products[:limit]
		hasNext = true
	}
	return products, hasNext, nil
}

func (service *ProductService) GetOneProduct(id int64) (model.Product, error) {
	product, err := service.product.FindById(id)
	if err != nil {
		return model.Product{}, err
	}
	if len(product.ProductPrices) > 0 {
		sort.Slice(product.ProductPrices, func(i, j int) bool {
			return product.ProductPrices[i].QuantityMultiplier > product.ProductPrices[j].QuantityMultiplier
		})
	}
	return product, nil
}

func (service *ProductService) GetProductsByIDs(IDs []int64) ([]model.Product, error) {
	products, err := service.product.FindByIDs(IDs)
	if err != nil {
		return []model.Product{}, err
	}
	for _, product := range products {
		if len(product.ProductPrices) > 0 {
			sort.Slice(product.ProductPrices, func(i, j int) bool {
				return product.ProductPrices[i].QuantityMultiplier > product.ProductPrices[j].QuantityMultiplier
			})
		}
	}
	return products, nil
}

func (service *ProductService) GetOneProductByCode(code string) (model.Product, error) {
	code = strings.ToLower(code)
	product, err := service.product.FindByExactCode(code)
	if err != nil {
		return model.Product{}, err
	}
	if len(product.ProductPrices) > 0 {
		sort.Slice(product.ProductPrices, func(i, j int) bool {
			return product.ProductPrices[i].QuantityMultiplier > product.ProductPrices[j].QuantityMultiplier
		})
	}
	return product, nil
}

func (service *ProductService) GetProductsByCode(productCode string) ([]model.Product, error) {
	products, err := service.product.FindByCode(productCode)
	if err != nil {
		return nil, err
	}
	for _, product := range products {
		if len(product.ProductPrices) > 0 {
			sort.Slice(product.ProductPrices, func(i, j int) bool {
				return product.ProductPrices[i].QuantityMultiplier > product.ProductPrices[j].QuantityMultiplier
			})
		}
	}
	return products, nil
}

func (service *ProductService) GetProductsBySearchQuery(query string, limit, page int) (products []model.Product, hasNext bool, err error) {
	if limit == 0 {
		return []model.Product{}, false, ErrEmptyParam
	}
	offset := (page - 1) * limit
	limitPlusOne := limit + 1
	products, err = service.product.SearchBy(query, limitPlusOne, offset)
	if err != nil {
		return nil, false, err
	}
	for _, product := range products {
		if len(product.ProductPrices) > 0 {
			sort.Slice(product.ProductPrices, func(i, j int) bool {
				return product.ProductPrices[i].QuantityMultiplier > product.ProductPrices[j].QuantityMultiplier
			})
		}
	}
	if len(products) == limitPlusOne {
		products = products[:limit]
		hasNext = true
	}
	return products, hasNext, nil
}

func (service *ProductService) GetProductsByName(productName string) ([]model.Product, error) {
	productName = strings.ToLower(productName)
	products, err := service.product.FindByName(productName)
	if err != nil {
		return nil, err
	}
	for _, product := range products {
		if len(product.ProductPrices) > 0 {
			sort.Slice(product.ProductPrices, func(i, j int) bool {
				return product.ProductPrices[i].QuantityMultiplier > product.ProductPrices[j].QuantityMultiplier
			})
		}
	}
	return products, nil
}

func (service *ProductService) GetProductsByCategoryName(categoryName string) ([]model.Product, error) {
	categoryName = strings.ToLower(categoryName)
	products, err := service.product.FindByCategoryName(categoryName)
	if err != nil {
		return nil, err
	}
	for _, product := range products {
		if len(product.ProductPrices) > 0 {
			sort.Slice(product.ProductPrices, func(i, j int) bool {
				return product.ProductPrices[i].QuantityMultiplier > product.ProductPrices[j].QuantityMultiplier
			})
		}
	}
	return products, nil
}

func (service *ProductService) GetProductsByUnitName(unitName string) ([]model.Product, error) {
	unitName = strings.ToLower(unitName)
	products, err := service.product.FindByUnitName(unitName)
	if err != nil {
		return nil, err
	}
	for _, product := range products {
		if len(product.ProductPrices) > 0 {
			sort.Slice(product.ProductPrices, func(i, j int) bool {
				return product.ProductPrices[i].QuantityMultiplier > product.ProductPrices[j].QuantityMultiplier
			})
		}
	}
	return products, nil
}

func (service *ProductService) NewProduct(productData string) (model.Product, error) {
	var product model.Product
	productDataByte := []byte(productData)
	err := json.Unmarshal(productDataByte, &product)
	if err != nil {
		return model.Product{}, err
	}
	product, err = service.product.New(product)
	if err != nil {
		return model.Product{}, err
	}
	// stock := model.Stock{
	// 	ProductID: int64(product.ID),
	// 	Product:   product,
	// 	Quantity:  0,
	// }
	// _, err = service.stock.New(stock)
	// if err != nil {
	// 	golog.Errorf("Error new product stock: %v", err)
	// }
	return product, nil
}

var sheetColumnName = map[string]string{
	"A1": "Barcode",
	"B1": "Nama Barang",
	"C1": "Harga Jual",
	"D1": "Stok",
	"E1": "Jenis Barang",
	"F1": "Harga Beli",
	"G1": "Satuan",
}

func (s *ProductService) ExportAllProductsToExcel() (*excelize.File, error) {
	products, err := s.product.FindAll()
	if err != nil {
		return nil, err
	}
	sheetName := "Products"
	xlsx := excelize.NewFile()

	xlsx.SetSheetName(xlsx.GetSheetName(1), sheetName)
	for key, val := range sheetColumnName {
		err := xlsx.SetCellValue(sheetName, key, val)
		if err != nil {
			return nil, err
		}
	}

	for i, product := range products {
		xlsx.SetCellValue(sheetName, fmt.Sprintf("A%d", i+2), product.Code)
		xlsx.SetCellValue(sheetName, fmt.Sprintf("B%d", i+2), product.Name)
		xlsx.SetCellValue(sheetName, fmt.Sprintf("C%d", i+2), product.SellPrice.Int64)
		xlsx.SetCellValue(sheetName, fmt.Sprintf("D%d", i+2), product.Quantity.Int64)
		xlsx.SetCellValue(sheetName, fmt.Sprintf("E%d", i+2), product.Category.Name)
		xlsx.SetCellValue(sheetName, fmt.Sprintf("F%d", i+2), product.BuyPrice.Int64)
		xlsx.SetCellValue(sheetName, fmt.Sprintf("G%d", i+2), product.Unit.Name)
	}

	err = xlsx.SaveAs("./exported-product.xlsx")
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
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		_, errIndex = s.importProducts(ctx, products)
		if len(errIndex) == len(products) {
			golog.Error("The products are not imported")
			cancel()
		}
		allProducts, err := s.product.FindAll()
		if err != nil {
			cancel()
		}

		golog.Info("Starting to compile the products")
		for _, product := range allProducts {
			var relatedIDs []int64
			relatedProducts, err := s.product.GetMultipleProductByExactCode(product.Code.String)
			if err != nil {
				continue
			}
			for _, rp := range relatedProducts {
				if rp.ID != product.ID {
					relatedIDs = append(relatedIDs, rp.ID)
				}
			}
			product.RelatedProducts = pq.Int64Array(relatedIDs)
			_, err = s.product.Update(product)
			if err != nil {
				golog.Errorf("Cannot update product; id: %d, err: %v", product.ID, err)
			}
		}
		golog.Info("Compiling product finished")
	}()
	return len(products), nil
}

func (s *ProductService) importProducts(ctx context.Context, products []model.Product) (int, []string) {
	errIndex := []string{}
	productCounter := 0
	for _, product := range products {
		golog.Infof("Product Name: %s\nQuantity: %d\nSell Price: %d\nBuy Price: %d\n\n", product.Name, product.Quantity.Int64, product.SellPrice.Int64, product.BuyPrice.Int64)
		categories, err := s.category.FindByName(strings.ToLower(product.Category.Name))
		var category model.Category
		if err != nil || len(categories) == 0 {
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
		units, err := s.unit.FindByName(strings.ToLower(product.Unit.Name))
		var unit model.Unit
		if err != nil || len(units) == 0 {
			unit = product.Unit
			unit, err = s.unit.New(unit)
			if err != nil {
				errIndex = append(errIndex, product.Name)
				continue
			}
		} else {
			unit = units[0]
		}
		product.UnitID = int64(unit.ID)
		product.Unit.ID = unit.ID
		product, err = s.product.New(product)
		if err != nil {
			errIndex = append(errIndex, product.Name)
			continue
		}
		productCounter++
	}
	if productCounter != len(products) {
		warnText := fmt.Sprintf("There are %v rows, but only %v products were created.", len(products), productCounter)
		golog.Warn(warnText)
		warnText = fmt.Sprintf("These are the name of products of unimported rows: %s", strings.Join(errIndex, ", "))
		golog.Warn(warnText)
	}
	golog.Infof("%v products imported!", productCounter)
	return productCounter, errIndex
}

func (service *ProductService) UpdateProduct(productData string) (model.Product, error) {
	var product model.Product
	productDataByte := []byte(productData)
	err := json.Unmarshal(productDataByte, &product)
	if product.ID < 1 {
		return product, errors.New("Empty ID")
	}
	if err != nil {
		return product, err
	}
	return service.product.Update(product)
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
		code := sql.NullString{}
		codeString := row[0]
		if codeString != "" {
			code = sql.NullString{
				String: codeString,
				Valid:  true,
			}
		}
		name := row[1]
		if name == "" {
			errIndex = append(errIndex, name)
			continue
		}

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
		unitCode := row[7]
		unitTotalPcs, err := strconv.ParseInt(row[8], 10, 64)
		if err != nil {
			unitTotalPcs = 1
		}

		unit := model.Unit{
			Name:     unitName,
			Code:     unitCode,
			TotalPcs: int32(unitTotalPcs),
		}

		var productPrices model.ProductPrices
		productPrices = productPrices.FromStringJson(row[9])

		var isOpenPrice bool
		if row[10] == "1" {
			isOpenPrice = true
		}

		product := model.Product{
			Code: code,
			Name: name,
			SellPrice: sql.NullInt64{
				Int64: sellPrice,
				Valid: true,
			},
			Quantity: sql.NullInt64{
				Int64: quantity,
				Valid: true,
			},
			Category: model.Category{
				Name: categoryName,
			},
			BuyPrice: sql.NullInt64{
				Int64: buyPrice,
				Valid: true,
			},
			Unit:          unit,
			ProductPrices: productPrices,
			IsOpenPrice:   isOpenPrice,
		}

		products = append(products, product)
	}
	return products, errIndex
}
