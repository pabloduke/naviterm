package render

import (
	"github.com/pabloduke/naviterm/api"
	"github.com/pabloduke/naviterm/internal/terminal"
	"github.com/pabloduke/naviterm/internal/types"
)

func DrawMenu(x int, y int, menu api.Menu) {
	drawMenuBorder(x+menu.Hpad, y+menu.Vpad, menu)
	//drawMenuTitle(x+menu.Hpad, y+menu.Vpad, menu)
	drawStyledMenuTitle(x+menu.Hpad, y+menu.Vpad, menu)
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

func drawMenuTitle(x int, y int, menu api.Menu) {
	DrawText(x, y-menu.Vpad, menu.Title, terminal.Attribute(menu.TitleColor.Attr()), terminal.Attribute(api.ColorDefault))
}

func drawStyledMenuTitle(x int, y int, menu api.Menu) {
	DrawTextStyled(x, y-menu.Vpad, menu.Title, menu.TitleStyle)
}

// TODO: Render single and double borders
func drawMenuBorder(x int, y int, menu api.Menu) {
	//Determine longest name

	width := CalculateMenuWidth(menu)
	//Draw border from X Coor  past long menu name length (wheter item or title)
	for ix := 0 - menu.Hpad; ix <= width; ix++ {
		//top
		DrawText(
			ix+x,
			y-menu.Vpad,
			menu.MenuBorder.TopBorder,
			terminal.Attribute(menu.BorderColor.Attr()),
			terminal.Attribute(api.ColorDefault),
		)

		//bottom
		DrawText(
			ix+x,
			y+menu.MaxHeight+menu.Vpad,
			menu.MenuBorder.BottomBorder,
			terminal.Attribute(menu.BorderColor.Attr()),
			terminal.Attribute(api.ColorDefault),
		)
	}

	//Draw border on left and right
	for jy := 0 - menu.Vpad; jy <= menu.MaxHeight+menu.Vpad; jy++ {
		//left
		DrawText(
			x-menu.Hpad,
			jy+y,
			menu.MenuBorder.LeftBorder,
			terminal.Attribute(menu.BorderColor.Attr()),
			terminal.Attribute(api.ColorDefault),
		)

		//right
		DrawText(
			x+width,
			jy+y,
			menu.MenuBorder.RightBorder,
			terminal.Attribute(menu.BorderColor.Attr()),
			terminal.Attribute(api.ColorDefault),
		)
	}

	//Draw Corners
	DrawText(x-menu.Hpad, y-menu.Vpad, menu.MenuBorder.TopLeftCorner, terminal.Attribute(menu.BorderColor.Attr()), terminal.Attribute(api.ColorDefault))
	DrawText(x+width, y-menu.Vpad, menu.MenuBorder.TopRightCorner, terminal.Attribute(menu.BorderColor.Attr()), terminal.Attribute(api.ColorDefault))
	DrawText(x+width, y+menu.MaxHeight+menu.Vpad, menu.MenuBorder.BottomRightCorner, terminal.Attribute(menu.BorderColor.Attr()), terminal.Attribute(api.ColorDefault))
	DrawText(x-menu.Hpad, y+menu.MaxHeight+menu.Vpad, menu.MenuBorder.BottomLeftCorner, terminal.Attribute(menu.BorderColor.Attr()), terminal.Attribute(api.ColorDefault))
}
