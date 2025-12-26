package menuitems

import (
	"naviterm/internal/data"

	"github.com/nsf/termbox-go"
)

var RedSaber = data.MenuItem{Name: "Red Saber", Color: termbox.ColorRed}
var BlueSaber = data.MenuItem{Name: "Blue Saber", Color: termbox.ColorBlue}
var GreenSaber = data.MenuItem{Name: "Green Saber", Color: termbox.ColorGreen}
var YellowSaber = data.MenuItem{Name: "Yellow Saber", Color: termbox.ColorYellow}
var PurpleSaber = data.MenuItem{Name: "Purple Saber", Color: termbox.ColorMagenta}
var CrimsonSaber = data.MenuItem{Name: "Crimson Saber", Color: termbox.ColorRed}

func GetJediSabers() []data.MenuItem {
	return []data.MenuItem{BlueSaber, GreenSaber, YellowSaber, PurpleSaber}
}

func GetSithSabers() []data.MenuItem {
	return []data.MenuItem{RedSaber, CrimsonSaber, PurpleSaber}
}
