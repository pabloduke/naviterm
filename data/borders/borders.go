package borders

import (
	"github.com/pabloduke/naviterm/data"
	"github.com/pabloduke/naviterm/data/symbols"
)

var SquareBorder = data.MenuBorder{
	TopLeftCorner:     symbols.TopLeft,
	TopRightCorner:    symbols.TopRight,
	BottomLeftCorner:  symbols.BottomLeft,
	BottomRightCorner: symbols.BottomRight,
	TopBorder:         symbols.Hbar,
	BottomBorder:      symbols.Hbar,
	LeftBorder:        symbols.Vbar,
	RightBorder:       symbols.Vbar,
}

var HeavySquareBorder = data.MenuBorder{
	TopLeftCorner:     symbols.HeavyTopLeft,
	TopRightCorner:    symbols.HeavyTopRight,
	BottomLeftCorner:  symbols.HeavyBottomLeft,
	BottomRightCorner: symbols.HeavyBottomRight,
	TopBorder:         symbols.HeavyHbar,
	BottomBorder:      symbols.HeavyHbar,
	LeftBorder:        symbols.HeavyVbar,
	RightBorder:       symbols.HeavyVbar,
}

var RoundedBorder = data.MenuBorder{
	TopLeftCorner:     symbols.RoundedTopLeft,
	TopRightCorner:    symbols.RoundedTopRight,
	BottomLeftCorner:  symbols.RoundedBottomLeft,
	BottomRightCorner: symbols.RoundedBottomRight,
	TopBorder:         symbols.Hbar,
	BottomBorder:      symbols.Hbar,
	LeftBorder:        symbols.Vbar,
	RightBorder:       symbols.Vbar,
}

var BlockBorder = data.MenuBorder{
	TopLeftCorner:     symbols.BlockFull,
	TopRightCorner:    symbols.BlockFull,
	BottomLeftCorner:  symbols.BlockFull,
	BottomRightCorner: symbols.BlockFull,
	TopBorder:         symbols.BlockUpperHalf,
	BottomBorder:      symbols.BlockLowerHalf,
	LeftBorder:        symbols.BlockLeftHalf,
	RightBorder:       symbols.BlockRightHalf,
}

var DoubleBorder = data.MenuBorder{
	TopLeftCorner:     symbols.DoubleTopLeft,
	TopRightCorner:    symbols.DoubleTopRight,
	BottomLeftCorner:  symbols.DoubleBottomLeft,
	BottomRightCorner: symbols.DoubleBottomRight,
	TopBorder:         symbols.DoubleHbar,
	BottomBorder:      symbols.DoubleHbar,
	LeftBorder:        symbols.DoubleVbar,
	RightBorder:       symbols.DoubleVbar,
}
