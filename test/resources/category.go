package resources

import "../../model"

var Category1 = model.Category{
	Template: model.Template{ID: 1},
	Name:     "Category1",
}

var Category2 = model.Category{
	Template: model.Template{ID: 2},
	Name:     "Category2",
}

var Category2Updated = model.Category{
	Template: model.Template{ID: 2},
	Name:     "Category2Updated",
}

var Category3 = model.Category{
	Template: model.Template{ID: 3},
	Name:     "Category3",
}

var Category4 = model.Category{
	Template: model.Template{ID: 4},
	Name:     "Category4",
}

var Categories = []model.Category{
	Category1,
	Category2,
	Category3,
	Category4,
}
