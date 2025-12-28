package render

import "github.com/nsf/termbox-go"

func DrawText(x, y int, text string, fg termbox.Attribute, bg termbox.Attribute) {
	for i, ch := range text {
		termbox.SetCell(x+i, y, ch, fg, bg)
	}

	Flush()
}
