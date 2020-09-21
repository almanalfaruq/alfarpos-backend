package model

type Role int

const (
	RoleAdmin Role = iota + 1
	RoleManager
	RoleCashier
)

const CTX_USER = "user"

type User struct {
	Template
	Username string `gorm:"unique_index" json:"username" example:"cashier"`
	Password string `json:"password,omitempty"`
	FullName string `json:"full_name" example:"Cashier Primary"`
	Address  string `json:"address" example:"Ketaon, Banyudono, Boyolali"`
	Phone    string `json:"phone" example:"0276 3283720"`
	RoleID   Role   `json:"role_id" example:"1"`
}

func (u *User) HasRole(roles ...Role) bool {
	for _, role := range roles {
		if u.RoleID == role {
			return true
		}
	}
	return false
}
