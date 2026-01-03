package menuitems

import (
	"github.com/pabloduke/naviterm/api"
)

var Human = api.MenuItem{Name: "Human", Color: api.GREEN}
var Twilek = api.MenuItem{Name: "Twi'lek", Color: api.GREEN}
var Chiss = api.MenuItem{Name: "Chiss", Color: api.GREEN}
var Togruta = api.MenuItem{Name: "Togruta", Color: api.GREEN}
var Rodian = api.MenuItem{Name: "Rodian", Color: api.GREEN}
var Gungan = api.MenuItem{Name: "Gungan", Color: api.GREEN}
var Jawa = api.MenuItem{Name: "Jawa", Color: api.GREEN}
var Wookiee = api.MenuItem{Name: "Wookiee", Color: api.GREEN}

func GetSpecies() []api.MenuItem {
	return []api.MenuItem{Human, Twilek, Chiss, Togruta, Rodian, Gungan, Jawa, Wookiee}
}
