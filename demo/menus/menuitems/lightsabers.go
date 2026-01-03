package menuitems

import (
	"github.com/pabloduke/naviterm/api"
)

var RedSaber = api.MenuItem{Name: "Red Saber", Color: api.RED}
var BlueSaber = api.MenuItem{Name: "Blue Saber", Color: api.BLUE}
var GreenSaber = api.MenuItem{Name: "Green Saber", Color: api.GREEN}
var YellowSaber = api.MenuItem{Name: "Yellow Saber", Color: api.YELLOW}
var PurpleSaber = api.MenuItem{Name: "Purple Saber", Color: api.MAGENTA}
var CrimsonSaber = api.MenuItem{Name: "Crimson Saber", Color: api.RED}

func GetJediSabers() []api.MenuItem {
	return []api.MenuItem{BlueSaber, GreenSaber, YellowSaber, PurpleSaber}
}

func GetSithSabers() []api.MenuItem {
	return []api.MenuItem{RedSaber, CrimsonSaber, PurpleSaber}
}

var JediTextStyle = api.NavitermStyle{
	FgColor:       api.WHITE,
	BgColor:       api.BLUE,
	Underline:     true,
	Bold:          true,
	Blink:         false,
	Italic:        true,
	Strikethrough: false,
	Reverse:       false,
	Dim:           false,
	Hidden:        false,
}
