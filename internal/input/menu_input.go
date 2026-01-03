package input

import (
	"github.com/pabloduke/naviterm/api"
	"github.com/pabloduke/naviterm/internal/terminal"
	"github.com/pabloduke/naviterm/internal/types"
)

func GetMenuInput(menu api.Menu, menuCursor types.MenuCursor) types.MenuCursor {
	for {
		event := terminal.PollEvent()

		if event.Ch >= '1' && event.Ch <= '9' {
			index := int(event.Ch - '1')
			if index < len(menu.MenuItems) {
				// Jump directly to a numbered item and adjust offset to keep it visible
				newOffset := menuCursor.Offset
				maxOffset := 0
				if len(menu.MenuItems) > menu.MaxHeight {
					maxOffset = len(menu.MenuItems) - menu.MaxHeight
				}
				if index < newOffset {
					newOffset = index
				} else if index >= newOffset+menu.MaxHeight {
					newOffset = index - menu.MaxHeight + 1
				}
				if newOffset < 0 {
					newOffset = 0
				}
				if newOffset > maxOffset {
					newOffset = maxOffset
				}
				return types.MenuCursor{Position: index, Offset: newOffset, Selected: false}
			}
		}

		if event.Key == terminal.KeyArrowDown {
			if len(menu.MenuItems) == 0 {
				return menuCursor
			}
			newPos := menuCursor.Position + 1
			if newPos >= len(menu.MenuItems) {
				// wrap to top
				newPos = 0
			}
			newOffset := menuCursor.Offset
			// Ensure selection is visible within the window
			if newPos < newOffset {
				newOffset = newPos
			}
			if newPos >= newOffset+menu.MaxHeight {
				newOffset = newPos - menu.MaxHeight + 1
			}
			if newOffset < 0 {
				newOffset = 0
			}
			maxOffset := 0
			if len(menu.MenuItems) > menu.MaxHeight {
				maxOffset = len(menu.MenuItems) - menu.MaxHeight
			}
			if newOffset > maxOffset {
				newOffset = maxOffset
			}
			return types.MenuCursor{Position: newPos, Offset: newOffset, Selected: false}
		}

		if event.Key == terminal.KeyArrowUp {
			if len(menu.MenuItems) == 0 {
				return menuCursor
			}
			newPos := menuCursor.Position - 1
			if newPos < 0 {
				// wrap to bottom
				newPos = len(menu.MenuItems) - 1
			}
			newOffset := menuCursor.Offset
			// Ensure selection is visible within the window
			if newPos < newOffset {
				newOffset = newPos
			} else if newPos >= newOffset+menu.MaxHeight {
				newOffset = newPos - menu.MaxHeight + 1
			}
			if newOffset < 0 {
				newOffset = 0
			}
			maxOffset := 0
			if len(menu.MenuItems) > menu.MaxHeight {
				maxOffset = len(menu.MenuItems) - menu.MaxHeight
			}
			if newOffset > maxOffset {
				newOffset = maxOffset
			}
			return types.MenuCursor{Position: newPos, Offset: newOffset, Selected: false}
		}

		if event.Key == terminal.KeyEnter {
			return types.MenuCursor{
				Position: menuCursor.Position,
				Offset:   menuCursor.Offset,
				Selected: true,
			}
		}
	}
}
