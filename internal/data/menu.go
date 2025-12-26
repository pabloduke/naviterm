package data

import "github.com/nsf/termbox-go"

type Menu struct {
	Title       string
	TitleColor  termbox.Attribute
	MenuItems   []MenuItem
	BorderColor termbox.Attribute
	Vpad        int
	Hpad        int
}

type MenuItem struct {
	Name  string
	Color termbox.Attribute
}
