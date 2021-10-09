package order

import (
	"time"

	"github.com/almanalfaruq/alfarpos-backend/model"
	userentity "github.com/almanalfaruq/alfarpos-backend/model/user"
	"gorm.io/gorm"
)

type Order struct {
	ID           uint            `gorm:"primary_key" json:"id" example:"1"`
	Invoice      string          `gorm:"unique_index" json:"invoice" example:"5-9-2020-000001"`
	Status       int32           `gorm:"index:idx_status_time" json:"status"`
	UserID       int64           `json:"user_id" example:"1"`
	User         userentity.User `gorm:"foreignkey:UserID" json:"user"`
	CustomerID   int64           `json:"customer_id" example:"1"`
	Customer     model.Customer  `gorm:"foreignkey:CustomerID" json:"customer"`
	Total        int64           `json:"total" example:"130000"`
	AmountPaid   int64           `json:"amount_paid" example:"150000"`
	TotalChange  int64           `json:"total_change" example:"20000"`
	PPN          int64           `json:"ppn" example:"13000"`
	Discount     float32         `gorm:"default:0.00" json:"discount" example:"0.00"`
	PaymentID    int64           `json:"payment_id" example:"1"`
	Payment      model.Payment   `gorm:"foreignkey:PaymentID" json:"payment"`
	OrderDetails []OrderDetail   `gorm:"foreignkey:order_id" json:"order_details"`
	Note         string          `json:"note" example:"This was a pending transaction for specific customer"`
	CreatedAt    time.Time       `gorm:"index:idx_create_time;index:idx_status_time;not null;default:CURRENT_TIMESTAMP" json:"created_at" example:""`
	UpdatedAt    time.Time       `json:"updated_at" example:""`
	DeletedAt    *time.Time      `json:"deleted_at" example:""`
}

func (o *Order) AfterFind(tx *gorm.DB) (err error) {
	if o.User.Password != "" {
		o.User.Password = ""
	}
	return
}

// Status for order
const (
	StatusCanceled = -1
	StatusPending  = 1
	StatusFinish   = 2
)

// Param for sorting
const (
	SortAsc  = "asc"
	SortDesc = "desc"
)

type GetOrderUsingFilterParam struct {
	Statuses  []int32 `json:"statuses"`
	Invoice   string  `json:"invoice"`
	StartDate string  `json:"start_date"`
	EndDate   string  `json:"end_date"`
	Sort      string  `json:"sort"`
}
