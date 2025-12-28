package render

import (
	"github.com/nsf/termbox-go"
	"github.com/pabloduke/naviterm/data"
	"github.com/pabloduke/naviterm/internal/terminal"
	"github.com/pabloduke/naviterm/internal/types"
)

func DrawMenu(x int, y int, menu data.Menu) {
	drawMenuBorder(x+menu.Hpad, y+menu.Vpad, menu)
	drawMenuTitle(x+menu.Hpad, y+menu.Vpad, menu)
	DrawMenuItems(
		x+menu.Hpad,
		y+menu.Vpad,
		menu,
		types.MenuCursor{
			Position: 0,
			Offset:   0,
			Selected: false,
		},
	)

	terminal.Flush()
}

func drawMenuTitle(x int, y int, menu data.Menu) {
	DrawText(x, y-menu.Vpad, menu.Title, termbox.Attribute(menu.TitleColor.Attr()), termbox.ColorDefault)
}

// TODO: Render single and double borders
func drawMenuBorder(x int, y int, menu data.Menu) {
	//Determine longest name
	longestName := len(menu.Title)
	longestName = determineLongestName(longestName, menu)
	//Draw border from X Coor  past long menu name length (wheter item or title)
	for ix := 0 - menu.Hpad; ix <= longestName+menu.Hpad; ix++ {
		//top
		DrawText(
			ix+x,
			y-menu.Vpad,
			data.Hbar,
			termbox.Attribute(menu.BorderColor.Attr()),
			termbox.ColorDefault,
		)

		//bottom
		DrawText(
			ix+x,
			y+menu.MaxHeight+menu.Vpad,
			data.Hbar,
			termbox.Attribute(menu.BorderColor.Attr()),
			termbox.ColorDefault,
		)
	}

	//Draw border on left and right
	for jy := 0 - menu.Vpad; jy <= menu.MaxHeight+menu.Vpad; jy++ {
		//left
		DrawText(
			x-menu.Hpad,
			jy+y,
			data.Vbar,
			termbox.Attribute(menu.BorderColor.Attr()),
			termbox.ColorDefault,
		)

		//right
		DrawText(
			longestName+menu.Hpad+x,
			jy+y,
			data.Vbar,
			termbox.Attribute(menu.BorderColor.Attr()),
			termbox.ColorDefault,
		)
	}

	//Draw Corners
	DrawText(x-menu.Hpad, y-menu.Vpad, data.TopLeft, termbox.Attribute(menu.BorderColor.Attr()), termbox.ColorDefault)
	DrawText(x+menu.Hpad+longestName, y-menu.Vpad, data.TopRight, termbox.Attribute(menu.BorderColor.Attr()), termbox.ColorDefault)
	DrawText(x+menu.Hpad+longestName, y+menu.MaxHeight+menu.Vpad, data.BottomRight, termbox.Attribute(menu.BorderColor.Attr()), termbox.ColorDefault)
	DrawText(x-menu.Hpad, y+menu.MaxHeight+menu.Vpad, data.BottomLeft, termbox.Attribute(menu.BorderColor.Attr()), termbox.ColorDefault)
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
