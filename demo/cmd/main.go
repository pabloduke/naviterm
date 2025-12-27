// Package main contains a tiny interactive cmd showcasing how to use the
// naviterm app helpers and data types to build simple terminal menus with
// termbox-go. Run this to pick a faction (Jedi/Sith) and then a lightsaber.
package main

import (
	"naviterm/app"
	"naviterm/data"
	"naviterm/demo/menus/forceuser"
	"naviterm/demo/menus/menuitems"
)

func main() {
	// Initialize naviterm and app-level state; ensure cleanup on exit.
	app.Init()
	defer app.Close()

	/*Set a Custom Spinner, else default will be used*/
	//app.SetSpinner([]rune{'|', '/', '-', '\\'}, 100)
	//app.SetSpinner([]rune{'<', '^', '>', 'v'}, 100)

	// Use arrow keys to move or enter a numbered selection (works even when numbers are not shown),
	//Enter to confirm
	factionSelection := app.GetUserInput(10, 10, forceuser.FactionSelectMenu())

	var saberSelection data.MenuItem

	// Based on the first choice, present a second menu
	if factionSelection == menuitems.Jedi {
		saberSelection = app.GetUserInput(40, 10, forceuser.JediSaberMenu())
	} else {
		saberSelection = app.GetUserInput(40, 10, forceuser.SithSaberMenu())
	}

	homeworldSelection := app.GetUserInput(75, 10, forceuser.HomeworldMenu())

	app.ResetColor()
	app.PrintText(10, 20, "You have have selected to be a "+factionSelection.Name)
	app.PrintText(10, 21, "You will wield a "+saberSelection.Name)
	app.PrintText(10, 22, "You are from "+homeworldSelection.Name)

	userName := app.GetTextInput(10, 23, "Enter your name: ")

	if factionSelection == menuitems.Jedi {
		app.PrintText(10, 24, "Welcome, Master "+userName+"! You are a Jedi Knight!")
	} else {
		app.PrintText(10, 24, "Welcome, Darth "+userName+"! You are a Sith Lord!")
	}

	app.PrintTextWithSpinner(10, 26, "Press any key to continue...")

}
