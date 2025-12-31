package render

import (
	"github.com/nsf/termbox-go"
	"github.com/pabloduke/naviterm/data/symbols"
)

func DrawRectangle(x int, y int, w int, h int) {
	for width := 0; width < w; width++ {
		for height := 0; height < h; height++ {
			DrawText(x+width, y+height, symbols.BlockFull, termbox.ColorDefault, termbox.ColorDefault)
		}
	}
}
