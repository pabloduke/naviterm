package menuitems

import (
	"github.com/pabloduke/naviterm/data"
	"github.com/pabloduke/naviterm/data/color"
)

var Tatooine = data.MenuItem{Name: "Tatooine", Color: color.WHITE}
var Corellia = data.MenuItem{Name: "Corellia", Color: color.WHITE}
var Coruscant = data.MenuItem{Name: "Coruscant", Color: color.WHITE}
var Ruusan = data.MenuItem{Name: "Ruusan", Color: color.WHITE}
var Apatros = data.MenuItem{Name: "Apatros", Color: color.WHITE}
var Earth = data.MenuItem{Name: "Earth", Color: color.WHITE}
var Caprica = data.MenuItem{Name: "Caprica", Color: color.WHITE}
var Vulcan = data.MenuItem{Name: "Vulcan", Color: color.WHITE}

func GetHomeworlds() []data.MenuItem {
	return []data.MenuItem{Tatooine, Corellia, Coruscant, Ruusan, Apatros, Earth, Caprica, Vulcan}
}
