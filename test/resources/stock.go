package resources

import (
	"github.com/almanalfaruq/alfarpos-backend/model"
)

var Stock1 = model.Stock{
	Template:  model.Template{ID: 1},
	ProductID: 1,
	Product:   Product1,
	Quantity:  100,
}

var Stock2 = model.Stock{
	Template:  model.Template{ID: 2},
	ProductID: 2,
	Product:   Product2,
	Quantity:  200,
}

var Stock3 = model.Stock{
	Template:  model.Template{ID: 3},
	ProductID: 3,
	Product:   Product3,
	Quantity:  300,
}

var Stock4 = model.Stock{
	Template:  model.Template{ID: 4},
	ProductID: 4,
	Product:   Product4,
	Quantity:  400,
}

var Stocks = []model.Stock{
	Stock1,
	Stock2,
	Stock3,
	Stock4,
}
