package naviterm

import (
	"fmt"
	"time"

	"github.com/nsf/termbox-go"
	"github.com/pabloduke/naviterm/data"
	"github.com/pabloduke/naviterm/data/color"
	"github.com/pabloduke/naviterm/internal/input"
	"github.com/pabloduke/naviterm/internal/render"
	"github.com/pabloduke/naviterm/internal/terminal"
	"github.com/pabloduke/naviterm/internal/types"
)

var cursor = "█"
var spinnerFrames = []rune{'⠋', '⠙', '⠹', '⠸', '⠼', '⠴', '⠦', '⠧', '⠇', '⠏'}
var delay = 100

func SetCursor(c string) {
	cursor = c
}

func SetSpinner(frames []rune, delayMS int) {
	spinnerFrames = frames
	delay = delayMS

}

func Init() error {
	return terminal.Init()
}

func Close() {
	terminal.Close()
}

func PrintText(x int, y int, text string) {
	render.DrawText(x, y, text, termbox.ColorWhite, termbox.ColorDefault)
}

func PrintTextWithSpinner(x int, y int, text string) {
	done := make(chan struct{})
	go spinner(x-1, y, done)
	render.DrawText(x, y, text, termbox.ColorWhite, termbox.ColorDefault)
	termbox.PollEvent()
	render.DrawText(x-1, y, " ", termbox.ColorWhite, termbox.ColorDefault)
	close(done)
}

func PollEvent() {
	termbox.PollEvent()
}

func ResetColor() {
	print("\033[0m")
}

func GetMenuInput(x int, y int, menu data.Menu) data.MenuItem {
	menu = defaultMenu(menu)
	sitem := types.SelectedItem{
		ItemNumber: 0,
		Selected:   false,
	}

	render.DrawMenu(x, y, menu)

	for {
		render.DrawMenuItems(x+menu.Hpad, y+menu.Vpad, menu, sitem)
		sitem = input.GetMenuInput(menu, sitem)
		if sitem.Selected {
			return menu.MenuItems[sitem.ItemNumber]
		}
	}
}

func DrawMenuAsView(x int, y int, menu data.Menu) {
	menu = defaultMenu(menu)
	render.DrawMenu(x, y, menu)
}

// Sets defaults for menu values not passed in by user
func defaultMenu(menu data.Menu) data.Menu {
	if termbox.Attribute(menu.TitleColor.Attr()) == 0 {
		menu.TitleColor = color.WHITE
	}
	if termbox.Attribute(menu.BorderColor.Attr()) == 0 {
		menu.BorderColor = color.WHITE
	}

	if menu.MaxHeight < 1 {
		menu.MaxHeight = len(menu.MenuItems)
	}

	if menu.MaxHeight > len(menu.MenuItems) {
		menu.MaxHeight = len(menu.MenuItems)
	}

	if menu.Title == "" {
		menu.Title = "Menu"
	}

	if menu.Prefix == "" && menu.IsNumbered {
		menu.Prefix = ".) "
	}

	for i := 0; i < len(menu.MenuItems); i++ {
		if termbox.Attribute(menu.MenuItems[i].Color.Attr()) == 0 {
			menu.MenuItems[i].Color = color.WHITE
		}

		if menu.MenuItems[i].Name == "" {
			menu.MenuItems[i].Name = "Item"
		}
	}

	return menu
}

func GetTextInput(x int, y int, prompt string) string {
	render.DrawText(x, y, prompt+cursor, termbox.ColorWhite, termbox.ColorDefault)
	input := ""
	done := make(chan struct{})
	go spinner(x-1, y, done)

	for {
		event := termbox.PollEvent()

		if event.Key == termbox.KeyEnter {
			render.DrawText(x-1, y, " ", termbox.ColorWhite, termbox.ColorDefault)
			render.DrawText(x+len(prompt), y, getBlankLine(), termbox.ColorWhite, termbox.ColorDefault)
			render.DrawText(x+len(prompt), y, input, termbox.ColorWhite, termbox.ColorDefault)
			close(done)
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
		render.DrawText(x+len(prompt), y, input+cursor, termbox.ColorWhite, termbox.ColorDefault)
	}

}

func getBlankLine() string {
	blankLine := ""
	for i := 0; i < 80; i++ {
		blankLine += " "
	}

	return blankLine
}

func chopLast(s string) string {
	r := []rune(s)
	if len(r) == 0 {
		return s
	}
	return string(r[:len(r)-1])
}

func spinner(x int, y int, done chan struct{}) {
	i := 0

	for {
		select {
		case <-done:
			return
		default:
			sprite := fmt.Sprintf("%c", spinnerFrames[i%len(spinnerFrames)])
			render.DrawText(x, y, sprite, termbox.ColorGreen, termbox.ColorDefault)
			time.Sleep(time.Duration(delay) * time.Millisecond)
			i++
		}
	}
}
