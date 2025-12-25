// Package main contains a tiny interactive demo showcasing how to use the
// naviterm app helpers and data types to build simple terminal menus with
// termbox-go. Run this to pick a faction (Jedi/Sith) and then a lightsaber.
package main

import (
	"naviterm/internal/app"
	"naviterm/internal/data"

	"github.com/nsf/termbox-go"
)

func main() {
	// Top-level options (first screen): choose a faction.
	// Each MenuItem has a display Name and a Color used when rendering.
	jediItem := data.MenuItem{Name: "JEDI", Color: termbox.ColorBlue}
	sithItem := data.MenuItem{Name: "SITH", Color: termbox.ColorBlue}

	// Second-level options (shown after selecting faction): saber choices.
	redSaber := data.MenuItem{Name: "Red Saber", Color: termbox.ColorRed}
	blueSaber := data.MenuItem{Name: "Blue Saber", Color: termbox.ColorBlue}
	greenSaber := data.MenuItem{Name: "Green Saber", Color: termbox.ColorGreen}
	yellowSaber := data.MenuItem{Name: "Yellow Saber", Color: termbox.ColorYellow}
	purpleSaber := data.MenuItem{Name: "Purple Saber", Color: termbox.ColorMagenta}
	crimsonSaber := data.MenuItem{Name: "Crimson Saber", Color: termbox.ColorRed}

	// Build the submenu for Sith choices.
	sithItems := []data.MenuItem{redSaber, purpleSaber, crimsonSaber}
	sithMenu := data.Menu{
		Title:       "Choose your lightsaber:",
		MenuItems:   sithItems,
		BorderColor: termbox.ColorWhite, // menu border color
	}

	// Build the submenu for Jedi choices.
	jediMenu := data.Menu{
		Title:       "Choose your lightsaber:",
		MenuItems:   []data.MenuItem{blueSaber, greenSaber, yellowSaber},
		BorderColor: termbox.ColorWhite,
	}

	// Build the main menu from the two top-level options.
	items := []data.MenuItem{}
	items = append(items, jediItem, sithItem)
	mainmenu := data.Menu{
		Title:       "Choose your faction:",
		MenuItems:   items,
		BorderColor: termbox.ColorWhite,
	}

	// Initialize termbox and app-level state; ensure cleanup on exit.
	app.Init()
	defer app.Close()

	// Show the main menu at terminal coordinates (x=10, y=10) and get a selection.
	// Use arrow keys to move, Enter to confirm (see internal/app for details).
	factionSelection := app.GetUserInput(10, 10, mainmenu)

	saberSelection := blueSaber

	// Based on the first choice, present a second menu just below (y=20).
	if factionSelection == jediItem {
		saberSelection = app.GetUserInput(10, 20, jediMenu)
	} else {
		saberSelection = app.GetUserInput(10, 20, sithMenu)
	}

	println("\nYou have have selected to be a " + factionSelection.Name)
	println("You will wield a " + saberSelection.Name)

	// Flush final frame and wait for one more event before exiting,
	// so the user can see the final output.
	app.Flush()
	app.PollEvent()
}
