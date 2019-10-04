package model

type Order struct {
	Template
	Invoice    string
	UserID     int
	User       User
	CustomerID int
	Customer   Customer
	Total      int
	AmountPaid int
	PaymentID  int
	Payment    Payment
}
