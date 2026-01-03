package terminal

import (
	"github.com/gdamore/tcell/v3"
)

// Screen provides a backend-agnostic interface for terminal operations
type Screen interface {
	Init() error
	Close()
	Clear()
	Flush()
	Size() (width int, height int)
	SetCell(x, y int, ch rune, fg, bg Attribute)
	SetStyledCell(x, y int, ch rune, style tcell.Style)
	PollEvent() Event
}
type Attribute uint16

// Attribute represents a color/style attribute
// Currently uint16 to match termbox, will support tcell colors
