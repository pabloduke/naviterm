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

func drawMenu(x int, y int, menu data.Menu, sitem selectedItem) {
	for i, item := range menu.MenuItems {
		if menu.IsNumbered {
			drawMenuItemNumbered(x, y, i, sitem, item, menu.Prefix)
		} else {
			drawMenuItem(x, y, i, sitem, item, menu.Prefix)
		}
	}

	Flush()
}

func getUserInput(menu data.Menu, sitem selectedItem) selectedItem {
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

func drawMenuItemNumbered(x int, y int, i int, sitem selectedItem, item data.MenuItem, prefix string) {
	if i == sitem.itemNumber {
		drawText(x, y+i, strconv.Itoa(i+1)+prefix+item.Name, termbox.ColorBlack, termbox.ColorWhite)
	} else {
		drawText(x, y+i, strconv.Itoa(i+1)+prefix+item.Name, item.Color, termbox.ColorDefault)
	}
}

func drawMenuItem(x int, y int, i int, sitem selectedItem, item data.MenuItem, prefix string) {
	if i == sitem.itemNumber {
		drawText(x, y+i, prefix+item.Name, termbox.ColorBlack, termbox.ColorWhite)
	} else {
		drawText(x, y+i, prefix+item.Name, item.Color, termbox.ColorDefault)
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
	menu = defaultMenu(menu)
	sitem := selectedItem{
		itemNumber: 0,
		selected:   false,
	}

	renderBorder(x+menu.Hpad, y+menu.Vpad, menu)
	renderTitle(x+menu.Hpad, y+menu.Vpad, menu)

	for {
		drawMenu(x+menu.Hpad, y+menu.Vpad, menu, sitem)
		sitem = getUserInput(menu, sitem)
		if sitem.selected {
			return menu.MenuItems[sitem.itemNumber]
		}
	}
}

// Sets defaults for menu values not passed in by user
func defaultMenu(menu data.Menu) data.Menu {
	if menu.TitleColor == 0 {
		menu.TitleColor = termbox.ColorWhite
	}
	if menu.BorderColor == 0 {
		menu.BorderColor = termbox.ColorWhite
	}

	if menu.Title == "" {
		menu.Title = "Menu"
	}

	if menu.Prefix == "" && menu.IsNumbered {
		menu.Prefix = ".) "
	}

	for i := 0; i < len(menu.MenuItems); i++ {
		if menu.MenuItems[i].Color == 0 {
			menu.MenuItems[i].Color = termbox.ColorWhite
		}

		if menu.MenuItems[i].Name == "" {
			menu.MenuItems[i].Name = "Item"
		}
	}

	return menu
}

// TODO: Render single and double borders
func renderBorder(x int, y int, menu data.Menu) {
	//Determine longest name
	longestName := len(menu.Title)
	longestName = determineLongestName(longestName, menu)

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

func determineLongestName(longestName int, menu data.Menu) int {
	var currentName int
	for i := 0; i < len(menu.MenuItems); i++ {
		if i+1 > 9 {
			currentName = len(menu.MenuItems[i].Name) + len(menu.Prefix) + 1
		} else {
			currentName = len(menu.MenuItems[i].Name) + len(menu.Prefix) + 2
		}

		longestName = max(longestName, currentName)
	}

	return longestName
}

func renderTitle(x int, y int, menu data.Menu) {
	drawText(x, y-menu.Vpad, menu.Title, menu.TitleColor, termbox.ColorDefault)
}

func GetTextInput(x int, y int, prompt string) string {
	drawText(x, y, prompt, termbox.ColorWhite, termbox.ColorDefault)
	Flush()
	input := ""
	for {
		event := termbox.PollEvent()

		if event.Key == termbox.KeyEnter {
			return input
		}

		if event.Ch != 0 {
			input += string(event.Ch)
		}

		if event.Key == termbox.KeyBackspace2 {
			input = chopLast(input)
		}

		drawText(x+len(prompt), y, getBlankLine(), termbox.ColorWhite, termbox.ColorDefault)
		drawText(x+len(prompt), y, input, termbox.ColorWhite, termbox.ColorDefault)
		Flush()
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
