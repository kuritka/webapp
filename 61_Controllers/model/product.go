package model

import "webapp/54_ViewModel/viewmodel"

type Category struct {
	Name string
	Id   int
}

var (
	categories = []Category{{"one drive", 1}, {"Iomega", 7}, {"Notorr", 3}, {"Motorr", 10}}
	colors     = map[string][]viewmodel.Color{
		"one drive": {{"Red", true}, {"Blue", true}, {"Black", true}, {"White", true}},
		"Iomega":    {{"Grey", true}, {"Brown", true}, {"Brown", true}, {"Brown", true}},
		"Notorr":    {{"White", true}, {"Blue", true}, {"White", true}, {"Yellow", true}},
		"Empty":     {{"Transient", true}},
	}
)

func GetCategory(categoryId int) Category {
	for _, cat := range categories {
		if cat.Id == categoryId {
			return cat
		}
	}
	return Category{"Empty", 0}
}

func GetColors(categoryId int) []viewmodel.Color {
	cat := GetCategory(categoryId)
	return colors[cat.Name]
}
