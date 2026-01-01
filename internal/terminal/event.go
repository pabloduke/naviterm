package terminal

// Event represents a terminal event (keyboard, resize, etc)
type Event struct {
	Type EventType
	Key  Key
	Ch   rune
	Mod  Modifier
}

// EventType represents the type of terminal event
type EventType uint8

const (
	EventKey EventType = iota
	EventResize
	EventMouse
	EventError
)

// Key represents special keyboard keys
type Key uint16

const (
	KeyNone Key = iota
	KeyEnter
	KeyBackspace
	KeyBackspace2
	KeySpace
	KeyArrowUp
	KeyArrowDown
	KeyArrowLeft
	KeyArrowRight
	KeyEsc
	KeyTab
)

// Modifier represents keyboard modifiers (shift, ctrl, alt)
type Modifier uint8

const (
	ModNone Modifier = 0
	ModAlt  Modifier = 1 << iota
	ModCtrl
	ModShift
)
