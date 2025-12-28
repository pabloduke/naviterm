package forceuser

import (
	"github.com/pabloduke/naviterm/data"
	"github.com/pabloduke/naviterm/data/color"
	"github.com/pabloduke/naviterm/demo/menus/menuitems"
)

func FactionSelectMenu() data.Menu {
	return data.Menu{
		Title:      "Choose your faction:",
		TitleColor: color.MAGENTA,
		MenuItems: []data.MenuItem{
			menuitems.Jedi,
			menuitems.Sith,
		},
		BorderColor: color.WHITE,
		Vpad:        1,
		Hpad:        4,
		IsNumbered:  false,
	}
}

func SithSaberMenu() data.Menu {
	return data.Menu{
		Title:       "Choose your lightsaber:",
		TitleColor:  color.MAGENTA,
		MenuItems:   menuitems.GetSithSabers(),
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
		MenuItems:   menuitems.GetJediSabers(),
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
		MenuItems:   menuitems.GetHomeworlds(),
		IsNumbered:  true,
		Prefix:      "●",
	}
}
