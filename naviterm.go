package naviterm

import (
	"github.com/nsf/termbox-go"
	"github.com/pabloduke/naviterm/data"
	"github.com/pabloduke/naviterm/data/color"
	"github.com/pabloduke/naviterm/internal/input"
	"github.com/pabloduke/naviterm/internal/render"
	"github.com/pabloduke/naviterm/internal/terminal"
	"github.com/pabloduke/naviterm/internal/types"
)

func SetCursor(c string) {
	input.Cursor = c
}

func SetSpinner(frames []rune, delayMS int) {
	render.SpinnerFrames = frames
	render.Delay = delayMS
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
	go render.Spinner(x-1, y, done)
	render.DrawText(x, y, text, termbox.ColorWhite, termbox.ColorDefault)
	termbox.PollEvent()
	render.DrawText(x-1, y, " ", termbox.ColorWhite, termbox.ColorDefault)
	close(done)
}

func ClearArea(x int, y int, w int, h int) {
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			render.DrawText(x+j, y+i, " ", termbox.ColorDefault, termbox.ColorDefault)
		}
	}

	terminal.Flush()
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

		terminal.Flush()
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
	return input.GetTextInput(x, y, prompt)
}
