package terminal

// Screen provides a backend-agnostic interface for terminal operations
type Screen interface {
	Init() error
	Close()
	Clear()
	Flush()
	Size() (width int, height int)
	SetCell(x, y int, ch rune, fg, bg Attribute)
	PollEvent() Event
}

// Attribute represents a color/style attribute
// Currently uint16 to match termbox, will support tcell colors
type Attribute uint16
