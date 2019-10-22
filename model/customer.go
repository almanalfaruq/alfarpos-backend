package model

type Customer struct {
	Template
	Code    string `gorm:"unique_index" json:"code"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}
