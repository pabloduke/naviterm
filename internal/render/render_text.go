package render

import (
	"github.com/nsf/termbox-go"
	"github.com/pabloduke/naviterm/internal/terminal"
)

func DrawText(x, y int, text string, fg termbox.Attribute, bg termbox.Attribute) {
	for i, ch := range text {
		termbox.SetCell(x+i, y, ch, fg, bg)
	}

	terminal.Flush()
}
