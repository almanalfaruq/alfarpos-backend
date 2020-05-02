package model

type Order struct {
	Template
	Invoice      string        `gorm:"unique_index" json:"invoice"`
	UserID       int64         `json:"user_id"`
	User         User          `gorm:"foreignkey:UserID" json:"user"`
	CustomerID   int64         `json:"customer_id"`
	Customer     Customer      `gorm:"foreignkey:CustomerID" json:"customer"`
	Total        int64         `json:"total"`
	AmountPaid   int64         `json:"amount_paid"`
	TotalChange  int64         `json:"total_change"`
	PPN          int64         `json:"ppn"`
	Discount     float32       `gorm:"default:0.00" json:"discount"`
	PaymentID    int64         `json:"payment_id"`
	Payment      Payment       `gorm:"foreignkey:PaymentID" json:"payment"`
	OrderDetails []OrderDetail `gorm:"foreignkey:order_id" json:"order_details"`
}
