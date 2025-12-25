package app

import (
	"naviterm/internal/data"

	"github.com/nsf/termbox-go"
)

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

func drawMenu(x int, y int, menu data.Menu, sitem selectedItem) selectedItem {
	for i, item := range menu.MenuItems {
		if i == sitem.itemNumber {
			drawText(x, y+i, item.Name, item.Color, termbox.ColorWhite)
		} else {
			drawText(x, y+i, item.Name, item.Color, termbox.ColorDefault)
		}
	}

	Flush()

	for {
		event := termbox.PollEvent()

		if event.Key == termbox.KeyArrowDown {
			if sitem.itemNumber+1 > len(menu.MenuItems)-1 {
				return selectedItem{itemNumber: 0, selected: false}
			}

			return selectedItem{itemNumber: sitem.itemNumber + 1, selected: false}
		}

		if event.Key == termbox.KeyArrowUp {
			if sitem.itemNumber-1 < 0 {
				return selectedItem{itemNumber: len(menu.MenuItems) - 1, selected: false}
			}
			return selectedItem{itemNumber: sitem.itemNumber - 1, selected: false}
		}

		if event.Key == termbox.KeyEnter {
			return selectedItem{itemNumber: sitem.itemNumber, selected: true}
		}

	}

}

func drawText(x, y int, text string, fg termbox.Attribute, bg termbox.Attribute) {
	for i, ch := range text {
		termbox.SetCell(x+i, y, ch, fg, bg)
	}
}

func PollEvent() {
	termbox.PollEvent()
}

type selectedItem struct {
	itemNumber int
	selected   bool
}

func GetUserInput(x int, y int, menu data.Menu) int {
	sitem := selectedItem{
		itemNumber: 0,
		selected:   false,
	}

	for {
		sitem = drawMenu(x, y, menu, sitem)
		if sitem.selected {
			return sitem.itemNumber
		}
	}
}

func selectMenuItem(menu data.Menu, selectedItem int) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	termbox.Flush()

}
