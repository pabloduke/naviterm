package render

import (
	"fmt"
	"time"

	"github.com/pabloduke/naviterm/internal/terminal"
)

var SpinnerFrames = []rune{'⠋', '⠙', '⠹', '⠸', '⠼', '⠴', '⠦', '⠧', '⠇', '⠏'}
var Delay = 100

func Spinner(x int, y int, done chan struct{}) {
	i := 0

	for {
		select {
		case <-done:
			return
		default:
			sprite := fmt.Sprintf("%c", SpinnerFrames[i%len(SpinnerFrames)])
			DrawText(x, y, sprite, terminal.ColorGreen, terminal.ColorDefault)
			time.Sleep(time.Duration(Delay) * time.Millisecond)
			i++
			terminal.Flush()
		}
	}
}
