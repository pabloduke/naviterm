package input

import (
	"strings"

	"github.com/nsf/termbox-go"
	"github.com/pabloduke/naviterm/internal/terminal"
)

var Cursor = "█"

// GetTextInput reads user text input using injected dependencies
func GetTextInput(x int, y int, prompt string, events EventSource, renderer Renderer, termInfo TerminalInfo, spinner SpinnerRunner) string {
	renderer.DrawText(x, y, prompt+Cursor, terminal.ColorWhite, terminal.ColorDefault)
	input := ""
	done := make(chan struct{})
	spinner.RunSpinner(x-1, y, done)

	for {
		event := events.PollEvent()

		if event.Key == termbox.KeyEnter {
			renderer.DrawText(x-1, y, " ", terminal.ColorWhite, terminal.ColorDefault)
			renderer.DrawText(x+len(prompt), y, getBlankLine(termInfo), terminal.ColorWhite, terminal.ColorDefault)
			renderer.DrawText(x+len(prompt), y, input, terminal.ColorWhite, terminal.ColorDefault)
			close(done)
			renderer.Flush()
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

		renderer.DrawText(x+len(prompt), y, getBlankLine(termInfo), terminal.ColorWhite, terminal.ColorDefault)
		renderer.DrawText(x+len(prompt), y, input+Cursor, terminal.ColorWhite, terminal.ColorDefault)
		renderer.Flush()
	}

}

func getBlankLine(termInfo TerminalInfo) string {
	return strings.Repeat(" ", termInfo.Width())
}

func chopLast(s string) string {
	r := []rune(s)
	if len(r) == 0 {
		return s
	}
	return string(r[:len(r)-1])
}
