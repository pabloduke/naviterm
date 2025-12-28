package render

import (
	"strconv"

	"github.com/nsf/termbox-go"
	"github.com/pabloduke/naviterm/data"
	"github.com/pabloduke/naviterm/internal/terminal"
	"github.com/pabloduke/naviterm/internal/types"
)

func DrawMenuItems(x int, y int, menu data.Menu, sitem types.SelectedItem) {
	for i, item := range menu.MenuItems {
		if menu.IsNumbered {
			numberedPrefix := strconv.Itoa(i+1) + menu.Prefix
			drawMenuItem(x, y, i, sitem, item, numberedPrefix)
		} else {
			drawMenuItem(x, y, i, sitem, item, menu.Prefix)
		}
	}

	terminal.Flush()
}

func drawMenuItem(x int, y int, i int, sitem types.SelectedItem, item data.MenuItem, prefix string) {
	if i == sitem.ItemNumber {
		DrawText(x, y+i, prefix+item.Name, termbox.ColorBlack, termbox.ColorWhite)
	} else {
		DrawText(x, y+i, prefix+item.Name, termbox.Attribute(item.Color.Attr()), termbox.ColorDefault)
	}

	terminal.Flush()
}
