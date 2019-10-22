package resources

import "../../model"

var Payment1 = model.Payment{
	Template: model.Template{ID: 1},
	Name:     "Payment1",
}

var Payment2 = model.Payment{
	Template: model.Template{ID: 2},
	Name:     "Payment2",
}

var Payment2Updated = model.Payment{
	Template: model.Template{ID: 2},
	Name:     "Payment2Updated",
}

var Payment3 = model.Payment{
	Template: model.Template{ID: 3},
	Name:     "Payment3",
}

var Payment4 = model.Payment{
	Template: model.Template{ID: 4},
	Name:     "Payment4",
}

var Payments = []model.Payment{
	Payment1,
	Payment2,
	Payment3,
	Payment4,
}
