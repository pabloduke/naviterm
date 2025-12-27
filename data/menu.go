package data

import "github.com/pabloduke/naviterm/data/color"

type Menu struct {
	Title       string
	TitleColor  color.NaviTermColor
	MenuItems   []MenuItem
	BorderColor color.NaviTermColor
	Vpad        int
	Hpad        int
	IsNumbered  bool
	Prefix      string
}

type MenuItem struct {
	Name  string
	Color color.NaviTermColor
}
