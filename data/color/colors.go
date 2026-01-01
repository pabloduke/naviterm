package color

import "github.com/pabloduke/naviterm/internal/terminal"

// Attribute is naviterm's internal representation of a terminal color/style.
// Now uses terminal.Attribute which abstracts backend-specific color types.
type Attribute = terminal.Attribute

type NaviTermColor struct {
	attr Attribute
}

func (c NaviTermColor) Attr() Attribute { return c.attr }
func (c NaviTermColor) IsZero() bool    { return c.attr == 0 }

// Basic colors (values match terminal constants).
const (
	Default Attribute = terminal.ColorDefault
	Black   Attribute = terminal.ColorBlack
	Red     Attribute = terminal.ColorRed
	Green   Attribute = terminal.ColorGreen
	Yellow  Attribute = terminal.ColorYellow
	Blue    Attribute = terminal.ColorBlue
	Magenta Attribute = terminal.ColorMagenta
	Cyan    Attribute = terminal.ColorCyan
	White   Attribute = terminal.ColorWhite
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
