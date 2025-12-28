package input

import (
	"github.com/nsf/termbox-go"
	"github.com/pabloduke/naviterm/data"
	"github.com/pabloduke/naviterm/internal/types"
)

func GetMenuInput(menu data.Menu, sitem types.SelectedItem) types.SelectedItem {
	for {
		event := termbox.PollEvent()

		if event.Ch >= '1' && event.Ch <= '9' {
			index := int(event.Ch - '1')
			if index < len(menu.MenuItems) {
				return types.SelectedItem{ItemNumber: index, Selected: false}
			}
		}

		if event.Key == termbox.KeyArrowDown {
			if sitem.ItemNumber+1 > len(menu.MenuItems)-1 {
				return types.SelectedItem{ItemNumber: 0, Selected: false}
			}

			return types.SelectedItem{ItemNumber: sitem.ItemNumber + 1, Selected: false}
		}

		if event.Key == termbox.KeyArrowUp {
			if sitem.ItemNumber-1 < 0 {
				return types.SelectedItem{ItemNumber: len(menu.MenuItems) - 1, Selected: false}
			}
			return types.SelectedItem{ItemNumber: sitem.ItemNumber - 1, Selected: false}
		}

		if event.Key == termbox.KeyEnter {
			return types.SelectedItem{ItemNumber: sitem.ItemNumber, Selected: true}
		}
	}
}
