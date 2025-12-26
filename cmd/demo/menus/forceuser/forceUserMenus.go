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
		IsNumbered:  false,
	}
}

func SithSaberMenu() data.Menu {
	return data.Menu{
		Title:       "Choose your lightsaber:",
		TitleColor:  termbox.ColorMagenta,
		MenuItems:   menuitems.GetSithSabers(),
		BorderColor: termbox.ColorRed, // menu border color
		Vpad:        1,
		Hpad:        4,
		IsNumbered:  true,
	}
}

func JediSaberMenu() data.Menu {
	return data.Menu{
		Title:       "Choose your lightsaber:",
		TitleColor:  termbox.ColorMagenta,
		MenuItems:   menuitems.GetJediSabers(),
		BorderColor: termbox.ColorBlue,
		Vpad:        1,
		Hpad:        4,
		IsNumbered:  true,
	}
}

func HomeworldMenu() data.Menu {
	return data.Menu{
		Title:       "Choose Your Homeworld:",
		TitleColor:  termbox.ColorMagenta,
		BorderColor: termbox.ColorGreen,
		Vpad:        1,
		Hpad:        4,
		MenuItems:   menuitems.GetHomeworlds(),
		IsNumbered:  true,
	}
}
