package menuitems

import (
	"github.com/pabloduke/naviterm/api"
)

var Tatooine = api.MenuItem{Name: "Tatooine", Color: api.WHITE}
var Corellia = api.MenuItem{Name: "Corellia", Color: api.WHITE}
var Coruscant = api.MenuItem{Name: "Coruscant", Color: api.WHITE}
var Ruusan = api.MenuItem{Name: "Ruusan", Color: api.WHITE}
var Apatros = api.MenuItem{Name: "Apatros", Color: api.WHITE}
var Naboo = api.MenuItem{Name: "Naboo", Color: api.WHITE}
var Mustafar = api.MenuItem{Name: "Mustafar", Color: api.WHITE}
var Endor = api.MenuItem{Name: "Endor", Color: api.WHITE}
var Dathomir = api.MenuItem{Name: "Dathomir", Color: api.WHITE}
var Kashyyyk = api.MenuItem{Name: "Kashyyyk", Color: api.WHITE}

func GetHomeworlds() []api.MenuItem {
	return []api.MenuItem{
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
