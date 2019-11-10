package resources

import (
	"github.com/almanalfaruq/alfarpos-backend/model"
)

var Order1 = model.Order{
	Template:     model.Template{ID: 1},
	Invoice:      "Order1",
	UserID:       1,
	User:         User1,
	CustomerID:   1,
	Customer:     Customer1,
	Total:        17000,
	AmountPaid:   20000,
	TotalChange:  3000,
	PPN:          0,
	Discount:     0.00,
	PaymentID:    1,
	Payment:      Payment1,
	OrderDetails: OrderDetails1,
}

var Order2 = model.Order{
	Template:     model.Template{ID: 2},
	Invoice:      "Order2",
	UserID:       2,
	User:         User2,
	CustomerID:   2,
	Customer:     Customer2,
	Total:        28500,
	AmountPaid:   30000,
	TotalChange:  1500,
	PPN:          0,
	Discount:     0.00,
	PaymentID:    2,
	Payment:      Payment2,
	OrderDetails: OrderDetails2,
}
