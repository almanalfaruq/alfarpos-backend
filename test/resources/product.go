package resources

import (
	"github.com/almanalfaruq/alfarpos-backend/model"
)

func helperToInt64(number int64) *int64 {
	return &number
}

var Product1 = model.Product{
	Template:   model.Template{ID: 1},
	Code:       "Product1",
	Name:       "Product1",
	BuyPrice:   helperToInt64(1000),
	SellPrice:  helperToInt64(1500),
	Quantity:   helperToInt64(10),
	CategoryID: 1,
	Category:   Category1,
	UnitID:     1,
	Unit:       Unit1,
	ImageUrl:   "",
	Discount:   0.00,
}

var Product2 = model.Product{
	Template:   model.Template{ID: 2},
	Code:       "Product2",
	Name:       "Product2",
	BuyPrice:   helperToInt64(2000),
	SellPrice:  helperToInt64(2500),
	Quantity:   helperToInt64(10),
	CategoryID: 2,
	Category:   Category2,
	UnitID:     2,
	Unit:       Unit2,
	ImageUrl:   "",
	Discount:   0.00,
}

var Product3 = model.Product{
	Template:   model.Template{ID: 3},
	Code:       "Product3",
	Name:       "Product3",
	BuyPrice:   helperToInt64(3000),
	SellPrice:  helperToInt64(3500),
	Quantity:   helperToInt64(10),
	CategoryID: 3,
	Category:   Category3,
	UnitID:     3,
	Unit:       Unit3,
	ImageUrl:   "",
	Discount:   0.00,
}

var Product4 = model.Product{
	Template:   model.Template{ID: 4},
	Code:       "Product4",
	Name:       "Product4",
	BuyPrice:   helperToInt64(4000),
	SellPrice:  helperToInt64(4500),
	Quantity:   helperToInt64(10),
	CategoryID: 4,
	Category:   Category4,
	UnitID:     4,
	Unit:       Unit4,
	ImageUrl:   "",
	Discount:   0.00,
}

var Products = []model.Product{
	Product1,
	Product2,
	Product3,
	Product4,
}
