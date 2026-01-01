package render

import (
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
	DrawText(x, y-menu.Vpad, menu.Title, menu.TitleColor.Attr(), terminal.ColorDefault)
}

// TODO: Render single and double borders
func drawMenuBorder(x int, y int, menu data.Menu) {
	//Determine longest name

	width := CalculateMenuWidth(menu)
	//Draw border from X Coor  past long menu name length (wheter item or title)
	for ix := 0 - menu.Hpad; ix <= width; ix++ {
		//top
		DrawText(
			ix+x,
			y-menu.Vpad,
			menu.MenuBorder.TopBorder,
			menu.BorderColor.Attr(),
			terminal.ColorDefault,
		)

		//bottom
		DrawText(
			ix+x,
			y+menu.MaxHeight+menu.Vpad,
			menu.MenuBorder.BottomBorder,
			menu.BorderColor.Attr(),
			terminal.ColorDefault,
		)
	}

	//Draw border on left and right
	for jy := 0 - menu.Vpad; jy <= menu.MaxHeight+menu.Vpad; jy++ {
		//left
		DrawText(
			x-menu.Hpad,
			jy+y,
			menu.MenuBorder.LeftBorder,
			menu.BorderColor.Attr(),
			terminal.ColorDefault,
		)

		//right
		DrawText(
			x+width,
			jy+y,
			menu.MenuBorder.RightBorder,
			menu.BorderColor.Attr(),
			terminal.ColorDefault,
		)
	}

	//Draw Corners
	DrawText(x-menu.Hpad, y-menu.Vpad, menu.MenuBorder.TopLeftCorner, menu.BorderColor.Attr(), terminal.ColorDefault)
	DrawText(x+width, y-menu.Vpad, menu.MenuBorder.TopRightCorner, menu.BorderColor.Attr(), terminal.ColorDefault)
	DrawText(x+width, y+menu.MaxHeight+menu.Vpad, menu.MenuBorder.BottomRightCorner, menu.BorderColor.Attr(), terminal.ColorDefault)
	DrawText(x-menu.Hpad, y+menu.MaxHeight+menu.Vpad, menu.MenuBorder.BottomLeftCorner, menu.BorderColor.Attr(), terminal.ColorDefault)
}
