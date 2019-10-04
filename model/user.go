package model

type Role int

const (
	Admin Role = iota + 1
	Manager
	Cashier
)

type User struct {
	Template
	Username string `json:"name"`
	Password string `json:"-"`
	FullName string `json:"full_name"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
	RoleID   Role   `sql:"not null;type:ENUM(1, 2, 3)" json:"role_id"`
}
