package main

import (
	"database/sql"
	"flag"
	"fmt"
	"strconv"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/almanalfaruq/alfarpos-backend/model"
	productentity "github.com/almanalfaruq/alfarpos-backend/model/product"
)

/*
	Script used to transform excel file from existing data
	so it can be used for the new cashier program.
*/

var sheetColumnName = map[string]string{
	"A1": "Barcode",
	"B1": "Nama Barang",
	"C1": "Harga Jual",
	"D1": "Stok",
	"E1": "Jenis Barang",
	"F1": "Harga Beli",
	"G1": "Satuan",
	"H1": "Kode Satuan",
	"I1": "Jumlah Satuan",
	"J1": "Harga Khusus",
	"K1": "Harga Terbuka",
}

func main() {
	var excelFile, sheetName string
	flag.StringVar(&excelFile, "excelfile", "", "path to excel file")
	flag.StringVar(&sheetName, "sheetname", "", "sheet name of excel file")
	flag.Parse()

	if excelFile == "" {
		panic("Choose excelfile using -excelfile=path/to/excel")
	}

	if sheetName == "" {
		sheetName = "Sheet1"
	}

	excel, err := excelize.OpenFile(excelFile)
	if err != nil {
		panic(err)
	}

	rows := excel.GetRows(sheetName)
	if len(rows) < 1 {
		panic("Rows length < 1")
	}

	products := parseExcelRowsToProduct(rows)

	sheetName = "Products"
	xlsx := excelize.NewFile()

	xlsx.SetSheetName(xlsx.GetSheetName(1), sheetName)
	for key, val := range sheetColumnName {
		xlsx.SetCellValue(sheetName, key, val)
	}

	for i, product := range products {
		xlsx.SetCellValue(sheetName, fmt.Sprintf("A%d", i+2), product.Code.String)
		xlsx.SetCellValue(sheetName, fmt.Sprintf("B%d", i+2), product.Name)
		xlsx.SetCellValue(sheetName, fmt.Sprintf("C%d", i+2), product.SellPrice.Int64)
		xlsx.SetCellValue(sheetName, fmt.Sprintf("D%d", i+2), product.Quantity.Int64)
		xlsx.SetCellValue(sheetName, fmt.Sprintf("E%d", i+2), product.Category.Name)
		xlsx.SetCellValue(sheetName, fmt.Sprintf("F%d", i+2), product.BuyPrice.Int64)
		xlsx.SetCellValue(sheetName, fmt.Sprintf("G%d", i+2), product.Unit.Name)
		xlsx.SetCellValue(sheetName, fmt.Sprintf("H%d", i+2), product.Unit.Code)
		xlsx.SetCellValue(sheetName, fmt.Sprintf("I%d", i+2), product.Unit.TotalPcs)
		strProductPrices, err := product.ProductPrices.ToString()
		if err != nil {
			continue
		}
		xlsx.SetCellValue(sheetName, fmt.Sprintf("J%d", i+2), strProductPrices)
		xlsx.SetCellValue(sheetName, fmt.Sprintf("K%d", i+2), product.IsOpenPrice)
	}

	err = xlsx.SaveAs("./exported-product.xlsx")
	if err != nil {
		panic(err)
	}
}

func parseExcelRowsToProduct(rows [][]string) []productentity.Product {
	var products []productentity.Product
	// skip index 0 - Header
	for _, row := range rows[1:] {
		code := row[9]
		sellPrice, err := strconv.ParseInt(row[10], 10, 64)
		if err != nil {
			continue
		}
		categoryName := "Umum"
		buyPrice, err := strconv.ParseInt(row[14], 10, 64)
		if err != nil {
			continue
		}
		unitName := strings.Title(strings.ToLower(row[12]))

		name := row[1]
		if name == "" {
			continue
		}

		var isOpenPrice bool
		if row[23] == "1" {
			isOpenPrice = true
		}

		// for other product price based on its qty
		var productPrices productentity.ProductPrices
		multiplierStr := row[81]
		if multiplierStr != "" {
			multiplier, _ := strconv.ParseInt(multiplierStr, 10, 64)
			sellPrice, _ := strconv.ParseInt(row[86], 10, 64)
			productPrices = append(productPrices, productentity.ProductPrice{
				QuantityMultiplier: int32(multiplier),
				PricePerPacket: sql.NullInt64{
					Int64: sellPrice,
					Valid: true,
				},
			})
		}

		product := productentity.Product{
			Code:      getSqlNullString(code),
			Name:      name,
			SellPrice: getSqlNullInt64(sellPrice),
			Quantity:  getSqlNullInt64(10),
			Category: model.Category{
				Name: categoryName,
			},
			BuyPrice: getSqlNullInt64(buyPrice),
			Unit: model.Unit{
				Name:     unitName,
				Code:     strings.ToUpper(unitName),
				TotalPcs: 1,
			},
			ProductPrices: productPrices,
			IsOpenPrice:   isOpenPrice,
		}

		products = append(products, product)

		// import other size product
		unitCode := row[69]
		if unitCode != "" {
			sellPrice, err := strconv.ParseInt(row[70], 10, 64)
			if err != nil {
				continue
			}

			totalPcs, err := strconv.ParseInt(row[71], 10, 64)
			if err != nil {
				continue
			}

			unitName = strings.ToLower(row[72])
			product := productentity.Product{
				Code:      getSqlNullString(code),
				Name:      name,
				SellPrice: getSqlNullInt64(sellPrice),
				Quantity:  getSqlNullInt64(10),
				Category: model.Category{
					Name: categoryName,
				},
				BuyPrice: getSqlNullInt64(buyPrice),
				Unit: model.Unit{
					Name:     strings.Title(unitName),
					Code:     strings.ToUpper(unitCode),
					TotalPcs: int32(totalPcs),
				},
				IsOpenPrice: isOpenPrice,
			}

			products = append(products, product)

		}
	}

	return products
}

func getSqlNullString(s string) sql.NullString {
	if s == "" {
		return sql.NullString{}
	}
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}

func getSqlNullInt64(i int64) sql.NullInt64 {
	if i == 0 {
		return sql.NullInt64{}
	}
	return sql.NullInt64{
		Int64: i,
		Valid: true,
	}
}
