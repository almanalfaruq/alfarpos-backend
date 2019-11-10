package model

type Order struct {
	Template
	Invoice      string        `gorm:"unique_index" json:"invoice"`
	UserID       int           `json:"user_id"`
	User         User          `gorm:"foreignkey:UserID" json:"user"`
	CustomerID   int           `json:"customer_id"`
	Customer     Customer      `gorm:"foreignkey:CustomerID" json:"customer"`
	Total        int           `json:"total"`
	AmountPaid   int           `json:"amount_paid"`
	TotalChange  int           `json:"total_change"`
	PPN          int           `json:"ppn"`
	Discount     float32       `gorm:"default:0.00" json:"discount"`
	PaymentID    int           `json:"payment_id"`
	Payment      Payment       `gorm:"foreignkey:PaymentID" json:"payment"`
	OrderDetails []OrderDetail `gorm:"foreignkey:order_id" json:"order_details"`
}
