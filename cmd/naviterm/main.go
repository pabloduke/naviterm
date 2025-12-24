package main

import (
	"github.com/nsf/termbox-go"
)

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}

	defer termbox.Close()
	// Direct control - put text at x,y
	drawText(10, 10, "Hello, world!", termbox.ColorWhite)
	termbox.Flush()
	termbox.PollEvent()
}

func drawText(x, y int, text string, fg termbox.Attribute) {
	for i, ch := range text {
		termbox.SetCell(x+i, y, ch, fg, termbox.ColorDefault)
	}
}
