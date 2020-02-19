package viewmodel

type Base struct {
	Title  string
	Colors []Color
}

type Color struct {
	Color string
	Bold  bool
}

func NewBase() *Base {
	//Data should be taken from model and transform here
	colors := []Color{{"Red", true}, {"Blue", false}, {"White", true}, {"Black", false}}

	return &Base{
		Title:  "Lemonade",
		Colors: colors,
	}
}

func NewColors(colors []Color) *Base {
	//Data should be taken from model and transform here
	return &Base{
		Title:  "Lemonade",
		Colors: colors,
	}
}

func NewHome() *Base {
	//Data should be taken from model and transform here
	colors := []Color{{"Grey", true}, {"Yellow", false}, {"BANANA", false}}

	return &Base{
		Title:  "Tea",
		Colors: colors,
	}
}
