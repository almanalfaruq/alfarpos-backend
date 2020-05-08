package model

type Order struct {
	Template
	Invoice      string        `gorm:"unique_index" json:"invoice" example:"5-9-2020-000001"`
	UserID       int64         `json:"user_id" example:"1"`
	User         User          `gorm:"foreignkey:UserID" json:"user"`
	CustomerID   int64         `json:"customer_id" example:"1"`
	Customer     Customer      `gorm:"foreignkey:CustomerID" json:"customer"`
	Total        int64         `json:"total" example:"130000"`
	AmountPaid   int64         `json:"amount_paid" example:"150000"`
	TotalChange  int64         `json:"total_change" example:"20000"`
	PPN          int64         `json:"ppn" example:"13000"`
	Discount     float32       `gorm:"default:0.00" json:"discount" example:"0.00"`
	PaymentID    int64         `json:"payment_id" example:"1"`
	Payment      Payment       `gorm:"foreignkey:PaymentID" json:"payment"`
	OrderDetails []OrderDetail `gorm:"foreignkey:order_id" json:"order_details"`
}
