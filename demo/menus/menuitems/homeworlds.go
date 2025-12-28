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
var Naboo = data.MenuItem{Name: "Naboo", Color: color.WHITE}
var Mustafar = data.MenuItem{Name: "Mustafar", Color: color.WHITE}
var Endor = data.MenuItem{Name: "Endor", Color: color.WHITE}
var Dathomir = data.MenuItem{Name: "Dathomir", Color: color.WHITE}
var Kashyyyk = data.MenuItem{Name: "Kashyyyk", Color: color.WHITE}

func GetHomeworlds() []data.MenuItem {
	return []data.MenuItem{
		Tatooine,
		Corellia,
		Coruscant,
		Ruusan,
		Apatros,
		Naboo,
		Mustafar,
		Endor,
		Dathomir,
		Kashyyyk,
	}
}
