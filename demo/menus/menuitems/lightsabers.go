package menuitems

import (
	"github.com/pabloduke/naviterm/data"
	"github.com/pabloduke/naviterm/data/color"
)

var RedSaber = data.MenuItem{Name: "Red Saber", Color: color.RED}
var BlueSaber = data.MenuItem{Name: "Blue Saber", Color: color.BLUE}
var GreenSaber = data.MenuItem{Name: "Green Saber", Color: color.GREEN}
var YellowSaber = data.MenuItem{Name: "Yellow Saber", Color: color.YELLOW}
var PurpleSaber = data.MenuItem{Name: "Purple Saber", Color: color.MAGENTA}
var CrimsonSaber = data.MenuItem{Name: "Crimson Saber", Color: color.RED}

func GetJediSabers() []data.MenuItem {
	return []data.MenuItem{BlueSaber, GreenSaber, YellowSaber, PurpleSaber}
}

func GetSithSabers() []data.MenuItem {
	return []data.MenuItem{RedSaber, CrimsonSaber, PurpleSaber}
}
