package input

import "github.com/nsf/termbox-go"

// mockEventSource simulates terminal events for testing
type mockEventSource struct {
	events []termbox.Event
	index  int
}

func (m *mockEventSource) PollEvent() termbox.Event {
	if m.index >= len(m.events) {
		// Return Enter key if we run out of events (safety)
		return termbox.Event{Key: termbox.KeyEnter}
	}
	event := m.events[m.index]
	m.index++
	return event
}

// mockRenderer captures rendering calls for verification
type mockRenderer struct {
	drawCalls []drawCall
	flushes   int
}

type drawCall struct {
	x    int
	y    int
	text string
	fg   termbox.Attribute
	bg   termbox.Attribute
}

func (m *mockRenderer) DrawText(x, y int, text string, fg, bg termbox.Attribute) {
	m.drawCalls = append(m.drawCalls, drawCall{x, y, text, fg, bg})
}

func (m *mockRenderer) Flush() {
	m.flushes++
}

// mockTerminalInfo simulates terminal dimensions
type mockTerminalInfo struct {
	width int
}

func (m *mockTerminalInfo) Width() int {
	return m.width
}

// mockSpinnerRunner simulates spinner without actually running it
type mockSpinnerRunner struct {
	started   bool
	lastX     int
	lastY     int
	doneChan  chan struct{}
}

func (m *mockSpinnerRunner) RunSpinner(x, y int, done chan struct{}) {
	m.started = true
	m.lastX = x
	m.lastY = y
	m.doneChan = done
}
