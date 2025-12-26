// Package main contains a tiny interactive demo showcasing how to use the
// naviterm app helpers and data types to build simple terminal menus with
// termbox-go. Run this to pick a faction (Jedi/Sith) and then a lightsaber.
package main

import (
	"naviterm/app"
	"naviterm/cmd/demo/menus/forceuser"
	"naviterm/cmd/demo/menus/menuitems"
	"naviterm/data"
)

func main() {
	// Initialize naviterm and app-level state; ensure cleanup on exit.
	app.Init()
	defer app.Close()

	// Show the main menu at terminal coordinates (x=10, y=10) and get a selection.
	// Use arrow keys to move, Enter to confirm (see internal/app for details).
	factionSelection := app.GetUserInput(10, 10, forceuser.FactionSelectMenu())

	var saberSelection data.MenuItem

	// Based on the first choice, present a second menu just below (y=20).
	if factionSelection == menuitems.JediItem {
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

	if factionSelection == menuitems.JediItem {
		app.PrintText(10, 24, "Welcome, Master "+userName+"! You are a Jedi Knight!")
	} else {
		app.PrintText(10, 24, "Welcome, Darth "+userName+"! You are a Sith Lord!")
	}
	app.PrintText(10, 26, "Press any key to continue...")

	// Flush final frame and wait for one more event before exiting,
	// so the user can see the final output.
	app.Flush()
	app.PollEvent()
}
