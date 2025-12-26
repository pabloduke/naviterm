package app

import (
	"naviterm/internal/data"
	"strconv"

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

func PrintText(x int, y int, text string) {
	drawText(x, y, text, termbox.ColorWhite, termbox.ColorDefault)
	Flush()
}

func drawMenu(x int, y int, menu data.Menu, sitem selectedItem) selectedItem {
	for i, item := range menu.MenuItems {
		if i == sitem.itemNumber {
			drawText(x, y+i, strconv.Itoa(i+1)+"). "+item.Name, termbox.ColorBlack, termbox.ColorWhite)
		} else {
			drawText(x, y+i, strconv.Itoa(i+1)+"). "+item.Name, item.Color, termbox.ColorDefault)
		}
	}

	Flush()

	for {
		event := termbox.PollEvent()

		if event.Ch >= '1' && event.Ch <= '9' {
			index := int(event.Ch - '1')
			if index < len(menu.MenuItems) {
				return selectedItem{itemNumber: index, selected: false}
			}
		}

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

func ResetColor() {
	print("\033[0m")
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
	renderTitle(x, y, menu)

	for {
		sitem = drawMenu(x, y, menu, sitem)
		if sitem.selected {
			return menu.MenuItems[sitem.itemNumber]
		}
	}
}

func getUserInputFromMenu(x int, y int, menu data.Menu, sitem selectedItem) data.MenuItem {
	renderBorder(x, y, menu)
	renderTitle(x, y, menu)

	for {
		sitem = drawMenu(x, y, menu, sitem)
		if sitem.selected {
			return menu.MenuItems[sitem.itemNumber]
		}
	}
}

func renderBorder(x int, y int, menu data.Menu) {
	//Determine longest name
	longestName := len(menu.Title)
	for i := 0; i < len(menu.MenuItems); i++ {
		if len(menu.MenuItems[i].Name) > longestName {
			longestName = len(menu.MenuItems[i].Name)
		}
	}

	//Draw border from X Coor  past long menu name length (wheter item or title)
	for ix := 0 - menu.Hpad; ix <= longestName+menu.Hpad; ix++ {
		//top
		drawText(ix+x, y-menu.Vpad, data.Hbar, menu.BorderColor, termbox.ColorDefault)

		//bottom
		drawText(ix+x, y+len(menu.MenuItems)+menu.Vpad, data.Hbar, menu.BorderColor, termbox.ColorDefault)
	}

	//Draw border on left and right
	for jy := 0 - menu.Vpad; jy <= len(menu.MenuItems)+menu.Vpad; jy++ {
		//left
		drawText(x-menu.Hpad, jy+y, data.Vbar, menu.BorderColor, termbox.ColorDefault)

		//right
		drawText(longestName+menu.Hpad+x, jy+y, data.Vbar, menu.BorderColor, termbox.ColorDefault)
	}

	//Draw Corners
	drawText(x-menu.Hpad, y-menu.Vpad, data.TopLeft, menu.BorderColor, termbox.ColorDefault)
	drawText(x+menu.Hpad+longestName, y-menu.Vpad, data.TopRight, menu.BorderColor, termbox.ColorDefault)
	drawText(x+menu.Hpad+longestName, y+len(menu.MenuItems)+menu.Vpad, data.BottomRight, menu.BorderColor, termbox.ColorDefault)
	drawText(x-menu.Hpad, y+len(menu.MenuItems)+menu.Vpad, data.BottomLeft, menu.BorderColor, termbox.ColorDefault)
}

func renderTitle(x int, y int, menu data.Menu) {
	drawText(x, y-menu.Vpad, menu.Title, menu.TitleColor, termbox.ColorDefault)
}
