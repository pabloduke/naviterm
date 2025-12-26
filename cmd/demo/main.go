// Package main contains a tiny interactive demo showcasing how to use the
// naviterm app helpers and data types to build simple terminal menus with
// termbox-go. Run this to pick a faction (Jedi/Sith) and then a lightsaber.
package main

import (
	"naviterm/cmd/demo/menus/forceuser"
	"naviterm/cmd/demo/menus/menuitems"
	"naviterm/internal/app"
	"naviterm/internal/data"
)

func main() {
	factionMenu := forceuser.FactionSelectMenu()
	jediMenu := forceuser.JediSaberMenu()
	sithMenu := forceuser.SithSaberMenu()

	// Initialize termbox and app-level state; ensure cleanup on exit.
	app.Init()
	defer app.Close()

	// Show the main menu at terminal coordinates (x=10, y=10) and get a selection.
	// Use arrow keys to move, Enter to confirm (see internal/app for details).
	factionSelection := app.GetUserInput(10, 10, factionMenu)

	var saberHomeworldSelection []data.MenuItem

	// Based on the first choice, present a second menu just below (y=20).
	if factionSelection[0] == menuitems.JediItem {
		saberHomeworldSelection = app.GetUserInput(10, 20, jediMenu)
	} else {
		saberHomeworldSelection = app.GetUserInput(10, 20, sithMenu)
	}

	app.ResetColor()
	app.PrintText(10, 40, "You have have selected to be a "+factionSelection[0].Name)
	app.PrintText(10, 41, "You will wield a "+saberHomeworldSelection[0].Name)
	app.PrintText(10, 42, "You are from "+saberHomeworldSelection[1].Name)

	// Flush final frame and wait for one more event before exiting,
	// so the user can see the final output.
	app.Flush()
	app.PollEvent()
}
