package forceuser

import (
	"naviterm/cmd/demo/menus/menuitems"
	"naviterm/internal/data"

	"github.com/nsf/termbox-go"
)

func FactionSelectMenu() data.Menu {
	return data.Menu{
		Title:      "Choose your faction:",
		TitleColor: termbox.ColorMagenta,
		MenuItems: []data.MenuItem{
			menuitems.JediItem,
			menuitems.SithItem,
		},
		BorderColor: termbox.ColorWhite, // menu border color
		Vpad:        1,
		Hpad:        4,
	}
}

func SithSaberMenu() data.Menu {
	return data.Menu{
		Title:       "Choose your lightsaber:",
		TitleColor:  termbox.ColorMagenta,
		MenuItems:   menuitems.GetSithSabers(),
		BorderColor: termbox.ColorWhite, // menu border color
		Vpad:        1,
		Hpad:        4,
		SubMenus: []data.Menu{
			HomeworldMenu(),
		},
	}
}

func JediSaberMenu() data.Menu {
	return data.Menu{
		Title:       "Choose your lightsaber:",
		TitleColor:  termbox.ColorMagenta,
		MenuItems:   menuitems.GetJediSabers(),
		BorderColor: termbox.ColorWhite,
		Vpad:        1,
		Hpad:        4,
		SubMenus: []data.Menu{
			HomeworldMenu(),
		},
	}
}

func HomeworldMenu() data.Menu {
	return data.Menu{
		Title:       "Choose Your Homeworld:",
		TitleColor:  termbox.ColorMagenta,
		BorderColor: termbox.ColorWhite,
		Vpad:        1,
		Hpad:        4,
		MenuItems:   menuitems.GetHomeworlds(),
	}
}
