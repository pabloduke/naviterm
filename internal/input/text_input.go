package input

import (
	"strings"

	"github.com/nsf/termbox-go"
	"github.com/pabloduke/naviterm/internal/render"
	"github.com/pabloduke/naviterm/internal/terminal"
)

var Cursor = "█"

func GetTextInput(x int, y int, prompt string) string {
	render.DrawText(x, y, prompt+Cursor, termbox.ColorWhite, termbox.ColorDefault)
	input := ""
	done := make(chan struct{})
	go render.Spinner(x-1, y, done)

	for {
		event := termbox.PollEvent()

		if event.Key == termbox.KeyEnter {
			render.DrawText(x-1, y, " ", termbox.ColorWhite, termbox.ColorDefault)
			render.DrawText(x+len(prompt), y, getBlankLine(), termbox.ColorWhite, termbox.ColorDefault)
			render.DrawText(x+len(prompt), y, input, termbox.ColorWhite, termbox.ColorDefault)
			close(done)
			terminal.Flush()
			return input
		}

		if event.Ch != 0 {
			input += string(event.Ch)
		}

		if event.Key == termbox.KeyBackspace2 || event.Key == termbox.KeyBackspace {
			input = chopLast(input)
		}

		if event.Key == termbox.KeySpace {
			input += " "
		}

		render.DrawText(x+len(prompt), y, getBlankLine(), termbox.ColorWhite, termbox.ColorDefault)
		render.DrawText(x+len(prompt), y, input+Cursor, termbox.ColorWhite, termbox.ColorDefault)
		terminal.Flush()
	}

}

func getBlankLine() string {
	return strings.Repeat(" ", terminal.Width())
}

func chopLast(s string) string {
	r := []rune(s)
	if len(r) == 0 {
		return s
	}
	return string(r[:len(r)-1])
}
