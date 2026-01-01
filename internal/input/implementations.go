package input

import (
	"github.com/pabloduke/naviterm/internal/render"
	"github.com/pabloduke/naviterm/internal/terminal"
)

// TermboxEventSource wraps terminal event polling
type TermboxEventSource struct{}

func (t *TermboxEventSource) PollEvent() terminal.Event {
	return terminal.PollEvent()
}

// TerminalRenderer wraps rendering functionality
type TerminalRenderer struct{}

func (t *TerminalRenderer) DrawText(x, y int, text string, fg, bg terminal.Attribute) {
	render.DrawText(x, y, text, fg, bg)
}

func (t *TerminalRenderer) Flush() {
	terminal.Flush()
}

// TerminalInfoProvider wraps terminal information
type TerminalInfoProvider struct{}

func (t *TerminalInfoProvider) Width() int {
	return terminal.Width()
}

// RealSpinnerRunner wraps the real spinner functionality
type RealSpinnerRunner struct{}

func (r *RealSpinnerRunner) RunSpinner(x, y int, done chan struct{}) {
	go render.Spinner(x, y, done)
}

// NewDefaultDependencies creates production implementations
func NewDefaultDependencies() (EventSource, Renderer, TerminalInfo, SpinnerRunner) {
	return &TermboxEventSource{},
		&TerminalRenderer{},
		&TerminalInfoProvider{},
		&RealSpinnerRunner{}
}
