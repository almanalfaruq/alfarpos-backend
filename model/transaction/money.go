package transaction

import "github.com/almanalfaruq/alfarpos-backend/model"

type Money struct {
	model.Template
	Type   string  `db:"type" example:"in"`
	Amount float64 `db:"amount" example:"1000000"`
}
