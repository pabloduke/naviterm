package render

import (
	"github.com/pabloduke/naviterm/api"
	"github.com/pabloduke/naviterm/internal/terminal"
)

func ClearArea(x int, y int, w int, h int) {
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			DrawText(x+j, y+i, " ", terminal.Attribute(api.ColorDefault), terminal.Attribute(api.ColorDefault))
		}
	}

	terminal.Flush()
}

func DrawText(x, y int, text string, fg terminal.Attribute, bg terminal.Attribute) {
	for i, ch := range text {
		terminal.SetCell(x+i, y, ch, fg, bg)
	}
}

func DrawTextStyled(x, y int, text string, style api.NavitermStyle) {
	for i, ch := range text {
		terminal.SetStyledCell(x+i, y, ch, style)
	}
}
