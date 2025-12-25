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

func GetUserInput(x int, y int, menu data.Menu) data.MenuItem {
	sitem := selectedItem{
		itemNumber: 0,
		selected:   false,
	}

	renderBorder(x, y, menu)

	for {
		sitem = drawMenu(x, y, menu, sitem)
		if sitem.selected {
			return menu.MenuItems[sitem.itemNumber]
		}
	}
}

func renderBorder(x int, y int, menu data.Menu) {
	longestName := 0
	for i := 0; i < len(menu.MenuItems); i++ {
		if len(menu.MenuItems[i].Name) > longestName {
			longestName = len(menu.MenuItems[i].Name)
		}
	}

	for ix := 0; ix < longestName+4; ix++ {
		drawText(ix+x, y-1, "*", menu.BorderColor, termbox.ColorDefault)
		drawText(ix+x, y+len(menu.MenuItems), "*", menu.BorderColor, termbox.ColorDefault)
	}

	for jy := y - 1; jy < len(menu.MenuItems)+1+y; jy++ {
		drawText(x-1, jy, "*", menu.BorderColor, termbox.ColorDefault)
		drawText(longestName+4+x, jy, "*", menu.BorderColor, termbox.ColorDefault)
	}
}
