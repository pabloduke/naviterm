package forceuser

import (
	"naviterm/data"
	"naviterm/data/color"
	menuitems2 "naviterm/demo/menus/menuitems"
)

func FactionSelectMenu() data.Menu {
	return data.Menu{
		Title:      "Choose your faction:",
		TitleColor: color.MAGENTA,
		MenuItems: []data.MenuItem{
			menuitems2.Jedi,
			menuitems2.Sith,
		},
		BorderColor: color.WHITE,
		Vpad:        1,
		Hpad:        4,
		IsNumbered:  true,
	}
}

func SithSaberMenu() data.Menu {
	return data.Menu{
		Title:       "Choose your lightsaber:",
		TitleColor:  color.MAGENTA,
		MenuItems:   menuitems2.GetSithSabers(),
		BorderColor: color.RED,
		Vpad:        1,
		Hpad:        4,
		IsNumbered:  true,
		Prefix:      "⚔ ",
	}
}

func JediSaberMenu() data.Menu {
	return data.Menu{
		Title:       "Choose your lightsaber:",
		TitleColor:  color.MAGENTA,
		MenuItems:   menuitems2.GetJediSabers(),
		BorderColor: color.BLUE,
		Vpad:        1,
		Hpad:        4,
		IsNumbered:  true,
		Prefix:      "⚔ ",
	}
}

func HomeworldMenu() data.Menu {
	return data.Menu{
		Title:       "Choose Your Homeworld:",
		TitleColor:  color.MAGENTA,
		BorderColor: color.GREEN,
		Vpad:        1,
		Hpad:        4,
		MenuItems:   menuitems2.GetHomeworlds(),
		IsNumbered:  true,
		Prefix:      "●",
	}
}
