package render

import "github.com/pabloduke/naviterm/data"

func CalculateMenuWidth(menu data.Menu) int {
	var currentName int
	longestName := len(menu.Title)
	for i := 0; i < len(menu.MenuItems); i++ {
		if i+1 > 9 {
			currentName = len(menu.MenuItems[i].Name) + len(menu.Prefix) + 1
		} else {
			currentName = len(menu.MenuItems[i].Name) + len(menu.Prefix) + 2
		}

		longestName = max(longestName, currentName)
	}

	return longestName + menu.Hpad
}
