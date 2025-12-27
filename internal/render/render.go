package render

import (
	"github.com/nsf/termbox-go"
	data2 "naviterm/data"
	"naviterm/internal/types"
	"strconv"
)

func DrawMenuTitle(x int, y int, menu data2.Menu) {
	DrawText(x, y-menu.Vpad, menu.Title, termbox.Attribute(menu.TitleColor.Attr()), termbox.ColorDefault)
}

func DrawText(x, y int, text string, fg termbox.Attribute, bg termbox.Attribute) {
	for i, ch := range text {
		termbox.SetCell(x+i, y, ch, fg, bg)
	}

	Flush()
}

func DrawMenu(x int, y int, menu data2.Menu, sitem types.SelectedItem) {
	for i, item := range menu.MenuItems {
		if menu.IsNumbered {
			drawMenuItemNumbered(x, y, i, sitem, item, menu.Prefix)
		} else {
			drawMenuItem(x, y, i, sitem, item, menu.Prefix)
		}
	}

	Flush()
}

// TODO: Render single and double borders
func DrawMenuBorder(x int, y int, menu data2.Menu) {
	//Determine longest name
	longestName := len(menu.Title)
	longestName = determineLongestName(longestName, menu)
	//Draw border from X Coor  past long menu name length (wheter item or title)
	for ix := 0 - menu.Hpad; ix <= longestName+menu.Hpad; ix++ {
		//top
		DrawText(ix+x, y-menu.Vpad, data2.Hbar, termbox.Attribute(menu.BorderColor.Attr()), termbox.ColorDefault)

		//bottom
		DrawText(ix+x, y+len(menu.MenuItems)+menu.Vpad, data2.Hbar, termbox.Attribute(menu.BorderColor.Attr()), termbox.ColorDefault)
	}

	//Draw border on left and right
	for jy := 0 - menu.Vpad; jy <= len(menu.MenuItems)+menu.Vpad; jy++ {
		//left
		DrawText(x-menu.Hpad, jy+y, data2.Vbar, termbox.Attribute(menu.BorderColor.Attr()), termbox.ColorDefault)

		//right
		DrawText(longestName+menu.Hpad+x, jy+y, data2.Vbar, termbox.Attribute(menu.BorderColor.Attr()), termbox.ColorDefault)
	}

	//Draw Corners
	DrawText(x-menu.Hpad, y-menu.Vpad, data2.TopLeft, termbox.Attribute(menu.BorderColor.Attr()), termbox.ColorDefault)
	DrawText(x+menu.Hpad+longestName, y-menu.Vpad, data2.TopRight, termbox.Attribute(menu.BorderColor.Attr()), termbox.ColorDefault)
	DrawText(x+menu.Hpad+longestName, y+len(menu.MenuItems)+menu.Vpad, data2.BottomRight, termbox.Attribute(menu.BorderColor.Attr()), termbox.ColorDefault)
	DrawText(x-menu.Hpad, y+len(menu.MenuItems)+menu.Vpad, data2.BottomLeft, termbox.Attribute(menu.BorderColor.Attr()), termbox.ColorDefault)
}

func drawMenuItemNumbered(x int, y int, i int, sitem types.SelectedItem, item data2.MenuItem, prefix string) {
	if i == sitem.ItemNumber {
		DrawText(x, y+i, strconv.Itoa(i+1)+prefix+item.Name, termbox.ColorBlack, termbox.ColorWhite)
	} else {
		DrawText(x, y+i, strconv.Itoa(i+1)+prefix+item.Name, termbox.Attribute(item.Color.Attr()), termbox.ColorDefault)
	}
}

func drawMenuItem(x int, y int, i int, sitem types.SelectedItem, item data2.MenuItem, prefix string) {
	if i == sitem.ItemNumber {
		DrawText(x, y+i, prefix+item.Name, termbox.ColorBlack, termbox.ColorWhite)
	} else {
		DrawText(x, y+i, prefix+item.Name, termbox.Attribute(item.Color.Attr()), termbox.ColorDefault)
	}

	Flush()
}

func determineLongestName(longestName int, menu data2.Menu) int {
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

func Flush() {
	termbox.Flush()
}
