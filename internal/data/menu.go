package data

import "github.com/nsf/termbox-go"

type Menu struct {
	Title       string
	MenuItems   []MenuItem
	BorderColor termbox.Attribute
}

type MenuItem struct {
	Name  string
	Color termbox.Attribute
}
