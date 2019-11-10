package resources

import (
	"github.com/almanalfaruq/alfarpos-backend/model"
)

var User1 = model.User{
	Template: model.Template{ID: 1},
	Username: "User1",
	Password: "User1",
	FullName: "User1",
	Address:  "Address1",
	Phone:    "User1",
	RoleID:   model.Cashier,
}

var User2 = model.User{
	Template: model.Template{ID: 2},
	Username: "User2",
	Password: "User2",
	FullName: "User2",
	Address:  "Address2",
	Phone:    "User2",
	RoleID:   model.Manager,
}

var User3 = model.User{
	Template: model.Template{ID: 3},
	Username: "User3",
	Password: "User3",
	FullName: "User3",
	Address:  "Address3",
	Phone:    "User3",
	RoleID:   model.Admin,
}

var Users = []model.User{
	User1,
	User2,
	User3,
}
