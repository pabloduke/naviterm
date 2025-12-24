package data

import "github.com/nsf/termbox-go"

type Menu struct {
	MenuItems   []MenuItem
	BorderColor termbox.Attribute
}

type MenuItem struct {
	Name  string
	Color termbox.Attribute
}
