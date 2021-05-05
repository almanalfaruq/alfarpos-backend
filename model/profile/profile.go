package profile

import "github.com/almanalfaruq/alfarpos-backend/model"

type Profile struct {
	model.Template
	Name            string `db:"name" json:"name"`
	Address         string `db:"address" json:"address"`
	Phone           string `db:"phone" json:"phone"`
	ThankyouMessage string `db:"thankyou_message" json:"thankyou_message"`
	FootNote        string `db:"foot_note" json:"foot_note"`
}
