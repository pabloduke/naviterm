package data

import "github.com/nsf/termbox-go"

type Menu struct {
	Title       string
	MenuItems   []MenuItem
	BorderColor termbox.Attribute
	Vpad        int
	Hpad        int
}

type MenuItem struct {
	Name  string
	Color termbox.Attribute
}
