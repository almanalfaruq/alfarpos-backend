package resources

import (
	"fmt"
	"time"

	"../../model"
)

var Customer1 = model.Customer{
	Template: model.Template{ID: 1},
	Code:     fmt.Sprintf("CUST-%s", time.Now().Format("20060102150405")),
	Name:     "Customer1",
	Address:  "Boyolali",
	Phone:    "081225812599",
}

var Customer2 = model.Customer{
	Template: model.Template{ID: 2},
	Code:     fmt.Sprintf("CUST-%s", time.Now().Format("20060102150405")),
	Name:     "Customer2",
	Address:  "Boyolali",
	Phone:    "081225812599",
}

var Customer2Updated = model.Customer{
	Template: model.Template{ID: 2},
	Code:     fmt.Sprintf("CUST-%s", time.Now().Format("20060102150405")),
	Name:     "Customer2Updated",
	Address:  "Boyolali",
	Phone:    "081225812599",
}

var Customer3 = model.Customer{
	Template: model.Template{ID: 3},
	Code:     fmt.Sprintf("CUST-%s", time.Now().Format("20060102150405")),
	Name:     "Customer3",
	Address:  "Boyolali",
	Phone:    "081225812599",
}

var Customer4 = model.Customer{
	Template: model.Template{ID: 4},
	Code:     fmt.Sprintf("CUST-%s", time.Now().Format("20060102150405")),
	Name:     "Customer4",
	Address:  "Boyolali",
	Phone:    "081225812599",
}

var Customers = []model.Customer{
	Customer1,
	Customer2,
	Customer3,
	Customer4,
}
