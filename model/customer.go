package model

type Customer struct {
	Template
	Code    string `gorm:"unique_index" json:"code" example:"code-1"`
	Name    string `json:"name" example:"Customer"`
	Address string `json:"address" example:"Pengging, Banyudono, Boyolali"`
	Phone   string `json:"phone" example:"081234567890"`
}
