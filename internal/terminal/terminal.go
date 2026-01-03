package terminal

import (
	//"github.com/pabloduke/naviterm/internal/render"
	"github.com/pabloduke/naviterm/api"
)

// Package-level screen instance
var screen Screen

// Init initializes the terminal screen
func Init() error {
	tcellScreen, err := NewTcellScreen()
	if err != nil {
		return err
	}
	screen = tcellScreen
	return screen.Init()
}

// Close shuts down the terminal screen
func Close() {
	screen.Close()
}

// Flush flushes the screen buffer
func Flush() {
	screen.Flush()
}

// Width returns the terminal width
func Width() int {
	w, _ := screen.Size()
	return w
}

// Height returns the terminal height
func Height() int {
	_, h := screen.Size()
	return h
}

// SetCell sets a cell at the specified position
func SetCell(x, y int, ch rune, fg, bg Attribute) {
	screen.SetCell(x, y, ch, fg, bg)
}

func SetStyledCell(x int, y int, ch rune, style api.NavitermStyle) {
	tcellStyle := TransformToTcell(style)
	screen.SetStyledCell(x, y, ch, tcellStyle)
}

// PollEvent polls for a terminal event
func PollEvent() Event {
	return screen.PollEvent()
}
