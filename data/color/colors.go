package color

// Attribute is naviterm's internal representation of a terminal color/style.
// It intentionally does NOT expose termbox types in the public API.
type Attribute uint16

type NaviTermColor struct {
	attr Attribute
}

func (c NaviTermColor) Attr() Attribute { return c.attr }
func (c NaviTermColor) IsZero() bool    { return c.attr == 0 }

// Basic colors (values match termbox-go constants; app converts to termbox.Attribute).
const (
	Default Attribute = 0
	Black   Attribute = 1
	Red     Attribute = 2
	Green   Attribute = 3
	Yellow  Attribute = 4
	Blue    Attribute = 5
	Magenta Attribute = 6
	Cyan    Attribute = 7
	White   Attribute = 8
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
