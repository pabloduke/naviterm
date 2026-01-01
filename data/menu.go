package data

import "github.com/pabloduke/naviterm/data/color"

type Menu struct {
	Title       string
	TitleColor  color.NaviTermColor
	BorderColor color.NaviTermColor
	MaxHeight   int
	Vpad        int
	Hpad        int
	IsNumbered  bool
	Prefix      string
	MenuItems   []MenuItem
	MenuBorder  BorderStyle
}

type MenuItem struct {
	Name  string
	Color color.NaviTermColor
}

type BorderStyle struct {
	Color color.NaviTermColor

	TopLeftCorner,
	TopRightCorner,
	BottomLeftCorner,
	BottomRightCorner,
	TopBorder,
	BottomBorder,
	LeftBorder,
	RightBorder string
}
