package api

type Menu struct {
	Title       string
	TitleStyle  NavitermStyle
	TitleColor  NaviTermColor
	BorderColor NaviTermColor
	MaxHeight   int
	Vpad        int
	Hpad        int
	IsNumbered  bool
	Prefix      string
	MenuItems   []MenuItem
	MenuBorder  BorderStyle
}

type MenuItem struct {
	Name      string
	TextStyle NavitermStyle
	Color     NaviTermColor
}

type NavitermStyle struct {
	FgColor       NaviTermColor
	BgColor       NaviTermColor
	Underline     bool
	Bold          bool
	Blink         bool
	Italic        bool
	Strikethrough bool
	Reverse       bool
	Dim           bool
	Hidden        bool
}

type BorderStyle struct {
	Color NaviTermColor

	TopLeftCorner,
	TopRightCorner,
	BottomLeftCorner,
	BottomRightCorner,
	TopBorder,
	BottomBorder,
	LeftBorder,
	RightBorder string
}
