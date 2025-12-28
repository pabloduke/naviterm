package menuitems

import (
	"github.com/pabloduke/naviterm/data"
	"github.com/pabloduke/naviterm/data/color"
)

var Human = data.MenuItem{Name: "Human", Color: color.GREEN}
var Twilek = data.MenuItem{Name: "Twi'lek", Color: color.GREEN}
var Chiss = data.MenuItem{Name: "Chiss", Color: color.GREEN}
var Togruta = data.MenuItem{Name: "Togruta", Color: color.GREEN}
var Rodian = data.MenuItem{Name: "Rodian", Color: color.GREEN}
var Gungan = data.MenuItem{Name: "Gungan", Color: color.GREEN}
var Jawa = data.MenuItem{Name: "Jawa", Color: color.GREEN}
var Wookiee = data.MenuItem{Name: "Wookiee", Color: color.GREEN}

func GetSpecies() []data.MenuItem {
	return []data.MenuItem{Human, Twilek, Chiss, Togruta, Rodian, Gungan, Jawa, Wookiee}
}
