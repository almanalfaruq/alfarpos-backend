package model

type Order struct {
	Template
	Invoice      string        `gorm:"unique_index" json:"invoice"`
	UserID       int           `json:"user_id"`
	User         User          `json:"user"`
	CustomerID   int           `json:"customer_id"`
	Customer     Customer      `json:"customer"`
	Total        int           `json:"total"`
	AmountPaid   int           `json:"amount_paid"`
	PaymentID    int           `json:"payment_id"`
	Payment      Payment       `json:"payment"`
	OrderDetails []OrderDetail `gorm:"foreignkey:order_id" json:"order_details"`
}
