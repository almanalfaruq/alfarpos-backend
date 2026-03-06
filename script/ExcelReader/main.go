package main

import (
	"flag"
	"fmt"

	"github.com/xuri/excelize/v2"
)

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

	rows, err := excel.GetRows(sheetName)
	if err != nil {
		panic(err)
	}
	if len(rows) < 1 {
		panic("Rows length < 1")
	}

	fmt.Println(rows[:10])
}
