package render

import (
	"strconv"

	"github.com/pabloduke/naviterm/api"
	"github.com/pabloduke/naviterm/internal/terminal"
	"github.com/pabloduke/naviterm/internal/types"
)

func DrawMenuItems(x int, y int, menu api.Menu, menuCursor types.MenuCursor) {
	// Clear the menu content area to avoid leftover characters from longer names
	// Determine interior width based on longest name (title or items)
	width := CalculateMenuWidth(menu)
	ClearArea(x, y, width, menu.MaxHeight+menu.Vpad)

	// Determine the visible window based on offset and MaxHeight
	start := menuCursor.Offset
	if start < 0 {
		start = 0
	}
	end := menuCursor.Offset + menu.MaxHeight
	if end > len(menu.MenuItems) {
		end = len(menu.MenuItems)
	}

	// Draw only visible items, positioning them relative to the offset
	for i := start; i < end; i++ {
		item := menu.MenuItems[i]
		if menu.IsNumbered {
			numberedPrefix := strconv.Itoa(i+1) + menu.Prefix
			drawMenuItem(x, y, i, menuCursor, item, numberedPrefix)
		} else {
			drawMenuItem(x, y, i, menuCursor, item, menu.Prefix)
		}
	}
	terminal.Flush()
}

func drawMenuItem(x int, y int, i int, menuCursor types.MenuCursor, item api.MenuItem, prefix string) {
	drawY := y + (i - menuCursor.Offset)
	if i == menuCursor.Position {
		DrawText(x, drawY, prefix+item.Name, terminal.Attribute(api.ColorBlack), terminal.Attribute(api.ColorWhite))
	} else {
		DrawText(x, drawY, prefix+item.Name, terminal.Attribute(item.Color.Attr()), terminal.Attribute(api.ColorDefault))
	}
}
