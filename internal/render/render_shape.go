package render

import (
	"github.com/pabloduke/naviterm/api"
	"github.com/pabloduke/naviterm/internal/terminal"
)

func DrawRectangle(x int, y int, w int, h int, color api.NaviTermColor) {
	for width := 0; width < w; width++ {
		for height := h - 1; height >= 0; height-- {
			DrawText(
				x+width,
				y-height,
				api.BlockFull,
				terminal.Attribute(color.Attr()),
				terminal.Attribute(api.ColorDefault),
			)
		}
	}
}
