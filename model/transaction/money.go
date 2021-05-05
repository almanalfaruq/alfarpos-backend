package transaction

import "github.com/almanalfaruq/alfarpos-backend/model"

const (
	TypeIn  = 1
	TypeOut = -1
)

type Money struct {
	model.Template
	Type   int32   `db:"type" json:"type" example:"1"`
	Amount float64 `db:"amount" json:"amount" example:"1000000"`
	Note   string  `db:"note" json:"note" example:"Payment for unilever"`
}

type GetMoneyWithFilterReq struct {
	Types     []int32 `json:"types"`
	StartDate string  `json:"start_date"`
	EndDate   string  `json:"end_date"`
}
