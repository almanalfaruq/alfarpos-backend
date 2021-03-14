package model

type Unit struct {
	Template
	Name     string `json:"name" example:"Unit"`
	Code     string `json:"code" example:"KRT40"`
	TotalPcs int32  `gorm:"default:1" json:"total_pcs" example:"1"`
}
