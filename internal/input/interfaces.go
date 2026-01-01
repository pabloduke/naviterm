package input

import "github.com/nsf/termbox-go"

// EventSource provides terminal event polling capabilities
type EventSource interface {
	PollEvent() termbox.Event
}

// Renderer handles drawing text to the terminal
type Renderer interface {
	DrawText(x, y int, text string, fg, bg termbox.Attribute)
	Flush()
}

// TerminalInfo provides terminal dimensions
type TerminalInfo interface {
	Width() int
}

// SpinnerRunner manages spinner animation lifecycle
type SpinnerRunner interface {
	RunSpinner(x, y int, done chan struct{})
}
