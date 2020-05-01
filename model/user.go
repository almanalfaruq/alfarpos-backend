package model

type Role int

const (
	Admin Role = iota + 1
	Manager
	Cashier
)

type User struct {
	Template
	Username string `gorm:"unique_index" json:"username"`
	Password string `json:"password,omitempty"`
	FullName string `json:"full_name"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
	RoleID   Role   `json:"role_id"`
}

func (u *User) HasRole(roles ...Role) bool {
	for _, role := range roles {
		if u.RoleID == role {
			return true
		}
	}
	return false
}
