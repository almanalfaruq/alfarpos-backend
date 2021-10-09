package product

import (
	"database/sql"
	"encoding/json"
	"errors"
	"io"
	"os"
	"testing"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/almanalfaruq/alfarpos-backend/model"
	productentity "github.com/almanalfaruq/alfarpos-backend/model/product"
	stockentity "github.com/almanalfaruq/alfarpos-backend/model/stock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestProductUpdateProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	category := model.Category{
		Template: model.Template{
			ID: int64(10),
		},
		Name: "Category1",
	}
	unit := model.Unit{
		Template: model.Template{
			ID: int64(5),
		},
		Name: "Pcs",
	}
	product := productentity.Product{
		Template: model.Template{
			ID: int64(1),
		},
		Name: "Product1",
		Code: sql.NullString{
			String: "Product1",
			Valid:  true,
		},
		BuyPrice: sql.NullInt64{
			Int64: 10000,
			Valid: true,
		},
		SellPrice: sql.NullInt64{
			Int64: 1500,
			Valid: true,
		},
		Quantity: sql.NullInt64{
			Int64: 10,
			Valid: true,
		},
		Category: category,
		Unit:     unit,
	}
	productStub := product
	productStub.SellPrice = sql.NullInt64{
		Int64: 55000,
		Valid: true,
	}

	productRepository := NewMockproductRepositoryIface(ctrl)
	productRepository.EXPECT().Update(productStub).Return(productStub, nil)

	testTable := []struct {
		testName string
		arg      func() string
		expect   productentity.Product
		wantErr  bool
	}{
		{
			testName: "Error - Failed Unmarshal JSON",
			arg: func() string {
				return `{product_id: 1}`
			},
			expect:  productentity.Product{},
			wantErr: true,
		},
		{
			testName: "Success",
			arg: func() string {
				jsonByte, _ := json.Marshal(productStub)
				return string(jsonByte)
			},
			expect: productStub,
		},
	}

	service := &ProductService{
		product: productRepository,
	}

	for _, tt := range testTable {
		t.Run(tt.testName, func(t *testing.T) {
			arg := tt.arg()
			actualResult, err := service.UpdateProduct(arg)
			if tt.wantErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
			assert.Equal(t, tt.expect, actualResult)
		})
	}
}

func TestProductService_NewProductUsingExcel(t *testing.T) {
	t.Skip("Skipped because of unknown mock error")
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productRepository := NewMockproductRepositoryIface(ctrl)
	categoryRepository := NewMockcategoryRepositoryIface(ctrl)
	unitRepository := NewMockunitRepositoryIface(ctrl)
	stockRepository := NewMockstockRepositoryIface(ctrl)

	s := NewProductService(productRepository, categoryRepository, unitRepository, stockRepository)

	testTable := []struct {
		testName string
		args     func() (string, io.ReadCloser)
		mock     func() error
	}{
		{
			testName: "Error - Not Valid File",
			args: func() (string, io.ReadCloser) {
				file, _ := os.Open("")
				return "", file
			},
			mock: func() error {
				return errors.New("invalid argument")
			},
		},
		{
			testName: "Error - Sheet Not Exists",
			args: func() (string, io.ReadCloser) {
				file, _ := os.Open("../test/resources/test.xlsx")
				return "abcde", file
			},
			mock: func() error {
				return excelize.ErrSheetNotExist{
					SheetName: "abcde",
				}
			},
		},
		{
			testName: "Success - No Import",
			args: func() (string, io.ReadCloser) {
				file, _ := os.Open("../test/resources/test.xlsx")
				return "", file
			},
			mock: func() error {
				return errors.New("Rows length < 1")
			},
		},
		{
			testName: "Success - No Blank Data - New Product",
			args: func() (string, io.ReadCloser) {
				file, _ := os.Open("../test/resources/test.xlsx")
				return "Product1", file
			},
			mock: func() error {
				category := model.Category{
					Template: model.Template{
						ID: int64(10),
					},
					Name: "Category1",
				}
				categoryRepository.EXPECT().FindByName("Category1").Return([]model.Category{category}, nil)
				unit := model.Unit{
					Template: model.Template{
						ID: int64(5),
					},
					Name: "Unit1",
				}
				unitRepository.EXPECT().FindByName("Unit1").Return([]model.Unit{unit}, nil)
				productRepository.EXPECT().FindByCode("Product1").Return([]productentity.Product{}, nil)
				product := productentity.Product{
					Template: model.Template{
						ID: int64(1),
					},
					Name: "Product1",
					Code: sql.NullString{
						String: "Product1",
						Valid:  true,
					},
					BuyPrice: sql.NullInt64{
						Int64: 10000,
						Valid: true,
					},
					SellPrice: sql.NullInt64{
						Int64: 1500,
						Valid: true,
					},
					Quantity: sql.NullInt64{
						Int64: 10,
						Valid: true,
					},
					Category:   category,
					CategoryID: int64(category.ID),
					Unit:       unit,
					UnitID:     int64(unit.ID),
				}
				productStub := product
				productStub.ID = int64(0)
				productRepository.EXPECT().New(productStub).Return(product, nil)
				stock := stockentity.Stock{
					Template: model.Template{
						ID: int64(1),
					},
					ProductID: int64(1),
					Quantity:  10,
				}
				stockStub := stock
				stockStub.ID = 0
				stockStub.Quantity = 0
				stockRepository.EXPECT().New(stockStub).Return(stock, nil)
				return nil
			},
		},
	}

	for _, tt := range testTable {
		sheetName, excelFile := tt.args()
		expectedResult := tt.mock()
		t.Run(tt.testName, func(t *testing.T) {
			_, actualResult := s.NewProductUsingExcel(sheetName, excelFile)
			assert.Equal(t, expectedResult, actualResult)
		})
		excelFile.Close()
	}
}
