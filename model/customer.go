package model

type Customer struct {
	Template
	Code    string `json:"code"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}
