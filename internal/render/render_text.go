package render

import (
	"github.com/nsf/termbox-go"
	"github.com/pabloduke/naviterm/internal/terminal"
)

func ClearArea(x int, y int, w int, h int) {
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			DrawText(x+j, y+i, " ", termbox.ColorDefault, termbox.ColorDefault)
		}
	}

	terminal.Flush()
}

func DrawText(x, y int, text string, fg termbox.Attribute, bg termbox.Attribute) {
	for i, ch := range text {
		termbox.SetCell(x+i, y, ch, fg, bg)
	}
}
