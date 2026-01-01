package render

import (
	"github.com/nsf/termbox-go"
	"github.com/pabloduke/naviterm/data/color"
	"github.com/pabloduke/naviterm/data/symbols"
)

func DrawRectangle(x int, y int, w int, h int, color color.NaviTermColor) {
	for width := 0; width < w; width++ {
		for height := h - 1; height >= 0; height-- {
			DrawText(
				x+width,
				y-height,
				symbols.BlockFull,
				termbox.Attribute(color.Attr()),
				termbox.ColorDefault,
			)
		}
	}
}
