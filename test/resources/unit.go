package resources

import "github.com/almanalfaruq/alfarpos-backend/model"

var Unit1 = model.Unit{
	Template: model.Template{ID: 1},
	Name:     "Unit1",
}

var Unit2 = model.Unit{
	Template: model.Template{ID: 2},
	Name:     "Unit2",
}

var Unit2Updated = model.Unit{
	Template: model.Template{ID: 2},
	Name:     "Unit2Updated",
}

var Unit3 = model.Unit{
	Template: model.Template{ID: 3},
	Name:     "Unit3",
}

var Unit4 = model.Unit{
	Template: model.Template{ID: 4},
	Name:     "Unit4",
}

var Units = []model.Unit{
	Unit1,
	Unit2,
	Unit3,
	Unit4,
}
