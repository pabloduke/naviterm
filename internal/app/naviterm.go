package app

import (
	"naviterm/internal/data"

	"github.com/nsf/termbox-go"
)

var currentMenu = data.Menu{}

func Init() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}

}

func Close() {
	termbox.Close()
}

func Flush() {
	termbox.Flush()
}

func DrawMenu(x int, y int, menu data.Menu) {

	for i, item := range menu.MenuItems {
		drawText(x, y+i, item.Name, item.Color)
	}
}

func drawText(x, y int, text string, fg termbox.Attribute) {
	for i, ch := range text {
		termbox.SetCell(x+i, y, ch, fg, termbox.ColorDefault)
	}
}
