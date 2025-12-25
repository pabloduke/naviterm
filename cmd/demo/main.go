package main

import (
	"naviterm/internal/app"
	"naviterm/internal/data"

	"github.com/nsf/termbox-go"
)

func main() {
	jediItem := data.MenuItem{Name: "JEDI", Color: termbox.ColorBlue}
	sithItem := data.MenuItem{Name: "SITH", Color: termbox.ColorBlue}

	redSaber := data.MenuItem{Name: "Red Saber", Color: termbox.ColorRed}
	blueSaber := data.MenuItem{Name: "Blue Saber", Color: termbox.ColorBlue}
	greenSaber := data.MenuItem{Name: "Green Saber", Color: termbox.ColorGreen}
	yellowSaber := data.MenuItem{Name: "Yellow Saber", Color: termbox.ColorYellow}
	purpleSaber := data.MenuItem{Name: "Purple Saber", Color: termbox.ColorMagenta}
	crimsonSaber := data.MenuItem{Name: "Crimson Saber", Color: termbox.ColorRed}

	sithItems := []data.MenuItem{redSaber, purpleSaber, crimsonSaber}
	sithMenu := data.Menu{
		MenuItems:   sithItems,
		BorderColor: termbox.ColorWhite,
	}

	jediMenu := data.Menu{
		MenuItems:   []data.MenuItem{blueSaber, greenSaber, yellowSaber},
		BorderColor: termbox.ColorWhite,
	}

	items := []data.MenuItem{}

	items = append(items, jediItem, sithItem)

	mainmenu := data.Menu{
		MenuItems:   items,
		BorderColor: termbox.ColorWhite,
	}

	app.Init()
	defer app.Close()

	userSelection := app.GetUserInput(10, 10, mainmenu)

	if userSelection.Name == jediItem.Name {
		userSelection = app.GetUserInput(10, 20, jediMenu)
	} else {
		userSelection = app.GetUserInput(10, 20, sithMenu)
	}

	if userSelection.Name == blueSaber.Name {
		println("You chose Blue Saber!")
	}

	app.Flush()
	app.PollEvent()
}
