// Package main contains a tiny interactive cmd showcasing how to use the
// naviterm app helpers and types types to build simple terminal menus with
// termbox-go. Run this to pick a faction (Jedi/Sith) and then a lightsaber.
package main

import (
	"github.com/pabloduke/naviterm"
	"github.com/pabloduke/naviterm/data"
	"github.com/pabloduke/naviterm/data/color"
	"github.com/pabloduke/naviterm/demo/menus/forceuser"
	"github.com/pabloduke/naviterm/demo/menus/menuitems"
)

func main() {
	// Initialize naviterm and app-level state; ensure cleanup on exit.
	err := naviterm.Init()
	if err != nil {
		return
	}
	defer naviterm.Close()

	/*Set a Custom Spinner, else default will be used*/
	//app.SetSpinner([]rune{'|', '/', '-', '\\'}, 100)
	//app.SetSpinner([]rune{'<', '^', '>', 'v'}, 100)

	// Use arrow keys to move or enter a numbered selection (works even when numbers are not shown),
	//Enter to confirm
	factionSelection := naviterm.GetMenuInput(10, 10, forceuser.FactionSelectMenu())

	var saberSelection data.MenuItem

	// Based on the first choice, present a second menu
	if factionSelection == menuitems.Jedi {
		saberSelection = naviterm.GetMenuInput(40, 10, forceuser.JediSaberMenu())
	} else {
		saberSelection = naviterm.GetMenuInput(40, 10, forceuser.SithSaberMenu())
	}

	homeworldSelection := naviterm.GetMenuInput(75, 10, forceuser.HomeworldMenu())

	naviterm.ResetColor()
	//naviterm.PrintText(10, 20, "You have have selected to be a "+factionSelection.Name)
	//naviterm.PrintText(10, 21, "You will wield a "+saberSelection.Name)
	//naviterm.PrintText(10, 22, "You are from "+homeworldSelection.Name)

	userName := naviterm.GetTextInput(10, 23, "Enter your name: ")

	if factionSelection == menuitems.Jedi {
		naviterm.PrintText(10, 24, "Welcome, Master "+userName+"! You are a Jedi Knight!")
	} else {
		naviterm.PrintText(10, 24, "Welcome, Darth "+userName+"! You are a Sith Lord!")
	}

	naviterm.PrintTextWithSpinner(10, 26, "Press any key to continue...")

	selectionsViewMenu := data.Menu{
		Title:       userName,
		TitleColor:  color.WHITE,
		BorderColor: color.WHITE,
		MaxHeight:   5,
		Vpad:        1,
		Hpad:        4,
		IsNumbered:  false,
		Prefix:      "",
		MenuItems:   []data.MenuItem{factionSelection, saberSelection, homeworldSelection},
	}
	naviterm.DrawMenuAsView(10, 30, selectionsViewMenu)

	naviterm.PrintTextWithSpinner(10, 40, "Press any key to continue...")

}
