package forceuser

import (
	"github.com/pabloduke/naviterm/api"
	"github.com/pabloduke/naviterm/demo/menus/menuitems"
)

func FactionSelectMenu() api.Menu {
	return api.Menu{
		Title:      "Choose your faction:",
		TitleColor: api.MAGENTA,
		TitleStyle: api.NavitermStyle{
			FgColor: api.RED,
			Bold:    true,
			Italic:  true,
		},
		MenuItems: []api.MenuItem{
			menuitems.Jedi,
			menuitems.Sith,
		},
		BorderColor: api.WHITE,
		Vpad:        1,
		Hpad:        4,
		IsNumbered:  false,
		MenuBorder:  api.HeavySquareBorder,
	}
}

func SithSaberMenu() api.Menu {
	return api.Menu{
		Title:       "Choose your lightsaber:",
		TitleColor:  api.MAGENTA,
		MenuItems:   menuitems.GetSithSabers(),
		BorderColor: api.RED,
		Vpad:        1,
		Hpad:        4,
		IsNumbered:  true,
		Prefix:      "⚔ ",
		MenuBorder:  api.BlockBorder,
	}
}

func JediSaberMenu() api.Menu {
	return api.Menu{
		Title:       "Choose your lightsaber:",
		TitleColor:  api.MAGENTA,
		TitleStyle:  menuitems.JediTextStyle,
		MenuItems:   menuitems.GetJediSabers(),
		BorderColor: api.BLUE,
		Vpad:        1,
		Hpad:        4,
		IsNumbered:  true,
		Prefix:      "⚔ ",
		MenuBorder:  api.BlockBorder,
	}
}

func HomeworldMenu() api.Menu {
	return api.Menu{
		Title:       "Choose Your Homeworld:",
		TitleColor:  api.MAGENTA,
		BorderColor: api.GREEN,
		MaxHeight:   5,
		Vpad:        2,
		Hpad:        4,
		MenuItems:   menuitems.GetHomeworlds(),
		IsNumbered:  true,
		Prefix:      "●",
		MenuBorder:  api.RoundedBorder,
	}
}

func SpeciesMenu() api.Menu {
	return api.Menu{
		Title:       "Choose Your Species:",
		TitleColor:  api.YELLOW,
		BorderColor: api.WHITE,
		MaxHeight:   5,
		Vpad:        2,
		Hpad:        4,
		MenuItems:   menuitems.GetSpecies(),
		IsNumbered:  true,
		Prefix:      api.Lozenge,
		MenuBorder:  api.DoubleBorder,
	}
}
