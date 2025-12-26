package forceuser

import (
	"naviterm/cmd/demo/menus/menuitems"
	"naviterm/data"

	"github.com/nsf/termbox-go"
)

// No title defines will default to "Menu"
func FactionSelectMenu() data.Menu {
	return data.Menu{
		//Title:      "Choose your faction:",
		TitleColor: termbox.ColorMagenta,
		MenuItems: []data.MenuItem{
			menuitems.JediItem,
			menuitems.SithItem,
		},
		BorderColor: termbox.ColorWhite, // menu border color
		Vpad:        1,
		Hpad:        4,
		IsNumbered:  false,
		Prefix:      "-->",
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
		Prefix:      "!",
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
		Prefix:      "! ",
	}
}

// IsNumbered set to true without a define prefix will defult to 1.), 2.), etc...
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
