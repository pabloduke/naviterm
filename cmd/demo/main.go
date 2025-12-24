package main

import (
	"naviterm/internal/app"
	"naviterm/internal/data"

	"github.com/nsf/termbox-go"
)

func main() {
	item1 := data.MenuItem{Name: "JEDI", Color: termbox.ColorBlue}
	item2 := data.MenuItem{Name: "SITH", Color: termbox.ColorBlue}

	items := []data.MenuItem{}

	items = append(items, item1, item2)

	mainmenu := data.Menu{
		MenuItems:   items,
		BorderColor: termbox.ColorWhite,
	}

	app.Init()
	defer app.Close()

	app.DrawMenu(10, 10, mainmenu)

	app.Flush()
	termbox.PollEvent()

}
