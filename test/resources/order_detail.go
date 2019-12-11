package resources

import (
	"github.com/almanalfaruq/alfarpos-backend/model"
)

var OrderDetail1 = model.OrderDetail{
	Template:  model.Template{ID: 1},
	ProductID: 1,
	Product:   Product1,
	Quantity:  1,
	SubTotal:  int(*Product1.SellPrice) * 1,
	OrderID:   1,
}

var OrderDetail2 = model.OrderDetail{
	Template:  model.Template{ID: 2},
	ProductID: 2,
	Product:   Product2,
	Quantity:  2,
	SubTotal:  int(*Product2.SellPrice) * 2,
	OrderID:   1,
}

var OrderDetail3 = model.OrderDetail{
	Template:  model.Template{ID: 3},
	ProductID: 3,
	Product:   Product3,
	Quantity:  3,
	SubTotal:  int(*Product3.SellPrice) * 3,
	OrderID:   1,
}

var OrderDetail4 = model.OrderDetail{
	Template:  model.Template{ID: 4},
	ProductID: 1,
	Product:   Product1,
	Quantity:  4,
	SubTotal:  int(*Product1.SellPrice) * 4,
	OrderID:   2,
}

var OrderDetail5 = model.OrderDetail{
	Template:  model.Template{ID: 5},
	ProductID: 4,
	Product:   Product4,
	Quantity:  5,
	SubTotal:  int(*Product4.SellPrice) * 5,
	OrderID:   2,
}

var OrderDetails1 = []model.OrderDetail{
	OrderDetail1,
	OrderDetail2,
	OrderDetail3,
}

var OrderDetails2 = []model.OrderDetail{
	OrderDetail4,
	OrderDetail5,
}
