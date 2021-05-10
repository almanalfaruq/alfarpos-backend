package transaction

import "time"

const (
	TypeIn  = 1
	TypeOut = -1
)

type Money struct {
	ID        int64      `gorm:"primary_key" json:"id" example:"1"`
	Type      int32      `db:"type" gorm:"index:idx_type_time" json:"type" example:"1"`
	Amount    float64    `db:"amount" gorm:"type:decimal(10,2)" json:"amount" example:"1000000"`
	Note      string     `db:"note" json:"note" example:"Payment for unilever"`
	CreatedAt time.Time  `gorm:"not null;default:CURRENT_TIMESTAMP;index:idx_type_time" json:"created_at" example:""`
	UpdatedAt time.Time  `json:"updated_at" example:""`
	DeletedAt *time.Time `json:"deleted_at" example:""`
}

type GetMoneyWithFilterReq struct {
	Types     []int32 `json:"types"`
	StartDate string  `json:"start_date"`
	EndDate   string  `json:"end_date"`
}
