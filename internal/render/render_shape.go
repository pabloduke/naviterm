package render

import (
	"github.com/pabloduke/naviterm/data/color"
	"github.com/pabloduke/naviterm/data/symbols"
	"github.com/pabloduke/naviterm/internal/terminal"
)

func DrawRectangle(x int, y int, w int, h int, color color.NaviTermColor) {
	for width := 0; width < w; width++ {
		for height := h - 1; height >= 0; height-- {
			DrawText(
				x+width,
				y-height,
				symbols.BlockFull,
				color.Attr(),
				terminal.ColorDefault,
			)
		}
	}
}
