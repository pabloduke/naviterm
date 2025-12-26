package menuitems

import (
	"naviterm/internal/data"

	"github.com/nsf/termbox-go"
)

var Tatooine = data.MenuItem{Name: "Tatooine", Color: termbox.ColorWhite}
var Corellia = data.MenuItem{Name: "Corellia", Color: termbox.ColorWhite}
var Coruscant = data.MenuItem{Name: "Coruscant", Color: termbox.ColorWhite}
var Ruusan = data.MenuItem{Name: "Ruusan", Color: termbox.ColorWhite}
var Apatros = data.MenuItem{Name: "Apatros", Color: termbox.ColorWhite}

func GetHomeworlds() []data.MenuItem {
	return []data.MenuItem{Tatooine, Corellia, Coruscant, Ruusan, Apatros}
}
