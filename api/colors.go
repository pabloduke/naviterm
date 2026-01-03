package api

// Attribute is naviterm's internal representation of a terminal color/style.
// Now uses terminal.Attribute which abstracts backend-specific color types.

type NaviTermColor struct {
	attr Attribute
}

type Attribute uint16

func (c NaviTermColor) Attr() Attribute { return c.attr }
func (c NaviTermColor) IsZero() bool    { return c.attr == 0 }

// Basic colors (values match terminal constants).
const (
	Default Attribute = ColorDefault
	Black   Attribute = ColorBlack
	Red     Attribute = ColorRed
	Green   Attribute = ColorGreen
	Yellow  Attribute = ColorYellow
	Blue    Attribute = ColorBlue
	Magenta Attribute = ColorMagenta
	Cyan    Attribute = ColorCyan
	White   Attribute = ColorWhite
)

const (
	ColorDefault  Attribute = 0
	ColorBlack    Attribute = 1
	ColorRed      Attribute = 2
	ColorGreen    Attribute = 3
	ColorYellow   Attribute = 4
	ColorBlue     Attribute = 5
	ColorMagenta  Attribute = 6
	ColorCyan     Attribute = 7
	ColorWhite    Attribute = 8
	ColorLightCya Attribute = 14
)

var RED = NaviTermColor{attr: Red}
var GREEN = NaviTermColor{attr: Green}
var BLUE = NaviTermColor{attr: Blue}
var WHITE = NaviTermColor{attr: White}
var BLACK = NaviTermColor{attr: Black}
var MAGENTA = NaviTermColor{attr: Magenta}
var YELLOW = NaviTermColor{attr: Yellow}
var CYAN = NaviTermColor{attr: Cyan}
var DEFAULT = NaviTermColor{attr: Default}
var NON = NaviTermColor{attr: 0}
