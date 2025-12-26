package forceuser

import (
	"naviterm/data"
	menuitems2 "naviterm/demo/menus/menuitems"

	"github.com/nsf/termbox-go"
)

func FactionSelectMenu() data.Menu {
	return data.Menu{
		Title:      "Choose your faction:",
		TitleColor: termbox.ColorMagenta,
		MenuItems: []data.MenuItem{
			menuitems2.Jedi,
			menuitems2.Sith,
		},
		BorderColor: termbox.ColorWhite, // menu border color
		Vpad:        1,
		Hpad:        4,
		IsNumbered:  true,
	}
}

func SithSaberMenu() data.Menu {
	return data.Menu{
		Title:       "Choose your lightsaber:",
		TitleColor:  termbox.ColorMagenta,
		MenuItems:   menuitems2.GetSithSabers(),
		BorderColor: termbox.ColorRed, // menu border color
		Vpad:        1,
		Hpad:        4,
		IsNumbered:  true,
		Prefix:      "⚔ ",
	}
}

func JediSaberMenu() data.Menu {
	return data.Menu{
		Title:       "Choose your lightsaber:",
		TitleColor:  termbox.ColorMagenta,
		MenuItems:   menuitems2.GetJediSabers(),
		BorderColor: termbox.ColorBlue,
		Vpad:        1,
		Hpad:        4,
		IsNumbered:  true,
		Prefix:      "⚔ ",
	}
}

func HomeworldMenu() data.Menu {
	return data.Menu{
		Title:       "Choose Your Homeworld:",
		TitleColor:  termbox.ColorMagenta,
		BorderColor: termbox.ColorGreen,
		Vpad:        1,
		Hpad:        4,
		MenuItems:   menuitems2.GetHomeworlds(),
		IsNumbered:  true,
		Prefix:      "●",
	}
}
