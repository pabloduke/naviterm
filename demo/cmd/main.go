// Package main contains a tiny interactive cmd showcasing how to use the
// naviterm app helpers and types types to build simple terminal menus with
// termbox-go. Run this to pick a faction (Jedi/Sith) and then a lightsaber.
package main

import (
	"github.com/pabloduke/naviterm"
	"github.com/pabloduke/naviterm/data"
	"github.com/pabloduke/naviterm/data/chart"
	"github.com/pabloduke/naviterm/data/color"
	"github.com/pabloduke/naviterm/demo/menus/forceuser"
	"github.com/pabloduke/naviterm/demo/menus/menuitems"
	"github.com/pabloduke/naviterm/internal/render"
)

func main() {
	// Initialize naviterm and app-level state; ensure cleanup on exit.
	err := naviterm.Init()
	if err != nil {
		return
	}

	defer naviterm.Close()

	/*Set a Custom Spinner, else default will be used*/
	naviterm.SetSpinner([]rune{'|', '/', '-', '\\'}, 100)
	//app.SetSpinner([]rune{'<', '^', '>', 'v'}, 100)
	testBarChart()
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

	speciesSelection := naviterm.GetMenuInput(10, 20, forceuser.SpeciesMenu())
	homeworldSelection := naviterm.GetMenuInput(40, 20, forceuser.HomeworldMenu())

	naviterm.ResetColor()

	userName := naviterm.GetTextInput(10, 30, "Enter your name: ")

	if factionSelection == menuitems.Jedi {
		userName = "Master " + userName
		naviterm.PrintText(10, 35, "Welcome, "+userName+"! You are a Jedi Knight!")
	} else {
		userName = "Darth " + userName
		naviterm.PrintText(10, 36, "Welcome, "+userName+"! You are a Sith Lord!")
	}

	selectionsViewMenu := data.Menu{
		Title:       "Your Character:",
		TitleColor:  color.WHITE,
		BorderColor: color.WHITE,
		MaxHeight:   5,
		Vpad:        1,
		Hpad:        4,
		IsNumbered:  false,
		Prefix:      "",
		MenuItems: []data.MenuItem{
			{Name: "Name: " + userName, Color: color.WHITE},
			{Name: "Faction: " + factionSelection.Name, Color: color.WHITE},
			{Name: "Saber: " + saberSelection.Name, Color: color.WHITE},
			{Name: "Species: " + speciesSelection.Name, Color: color.WHITE},
			{Name: "HomeWorld: " + homeworldSelection.Name, Color: color.WHITE},
		},
	}
	naviterm.DrawMenuAsView(75, 15, selectionsViewMenu)

	naviterm.PrintTextWithSpinner(10, 38, "Press any key to continue...")
	naviterm.ClearArea(0, 0, 200, 200)
}

func testBarChart() {
	bcItem := chart.BarChartItem{
		Label: "Test",
		Value: 3,
		Color: color.MAGENTA,
	}

	bcItem2 := chart.BarChartItem{
		Label: "Test",
		Value: 10,
		Color: color.GREEN,
	}

	bcItem3 := chart.BarChartItem{
		Label: "Test",
		Value: 6,
		Color: color.BLUE,
	}

	bChart := chart.BarChart{
		Title:   "CHART",
		XLabel:  "Amount",
		YLabel:  "ITEM NAME",
		Items:   []chart.BarChartItem{bcItem, bcItem2, bcItem3},
		Spacing: 2,
	}
	naviterm.ClearScreen()
	render.DrawBarChart(50, 30, bChart)

	naviterm.PrintTextWithSpinner(10, 40, "Press any key to exit...")
}
