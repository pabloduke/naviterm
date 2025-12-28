package borders

import "github.com/pabloduke/naviterm/data"

var SquareBorder = data.MenuBorder{
	TopLeftCorner:     TopLeft,
	TopRightCorner:    TopRight,
	BottomLeftCorner:  BottomLeft,
	BottomRightCorner: BottomRight,
	TopBorder:         Hbar,
	BottomBorder:      Hbar,
	LeftBorder:        Vbar,
	RightBorder:       Vbar,
}

var HeavySquareBorder = data.MenuBorder{
	TopLeftCorner:     HeavyTopLeft,
	TopRightCorner:    HeavyTopRight,
	BottomLeftCorner:  HeavyBottomLeft,
	BottomRightCorner: HeavyBottomRight,
	TopBorder:         HeavyHbar,
	BottomBorder:      HeavyHbar,
	LeftBorder:        HeavyVbar,
	RightBorder:       HeavyVbar,
}

var RoundedBorder = data.MenuBorder{
	TopLeftCorner:     RoundedTopLeft,
	TopRightCorner:    RoundedTopRight,
	BottomLeftCorner:  RoundedBottomLeft,
	BottomRightCorner: RoundedBottomRight,
	TopBorder:         Hbar,
	BottomBorder:      Hbar,
	LeftBorder:        Vbar,
	RightBorder:       Vbar,
}

var BlockBorder = data.MenuBorder{
	TopLeftCorner:     BlockFull,
	TopRightCorner:    BlockFull,
	BottomLeftCorner:  BlockFull,
	BottomRightCorner: BlockFull,
	TopBorder:         BlockUpperHalf,
	BottomBorder:      BlockLowerHalf,
	LeftBorder:        BlockLeftHalf,
	RightBorder:       BlockRightHalf,
}
