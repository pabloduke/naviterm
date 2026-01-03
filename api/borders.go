package api

var SquareBorder = BorderStyle{
	TopLeftCorner:     TopLeft,
	TopRightCorner:    TopRight,
	BottomLeftCorner:  BottomLeft,
	BottomRightCorner: BottomRight,
	TopBorder:         Hbar,
	BottomBorder:      Hbar,
	LeftBorder:        Vbar,
	RightBorder:       Vbar,
}

var HeavySquareBorder = BorderStyle{
	TopLeftCorner:     HeavyTopLeft,
	TopRightCorner:    HeavyTopRight,
	BottomLeftCorner:  HeavyBottomLeft,
	BottomRightCorner: HeavyBottomRight,
	TopBorder:         HeavyHbar,
	BottomBorder:      HeavyHbar,
	LeftBorder:        HeavyVbar,
	RightBorder:       HeavyVbar,
}

var RoundedBorder = BorderStyle{
	TopLeftCorner:     RoundedTopLeft,
	TopRightCorner:    RoundedTopRight,
	BottomLeftCorner:  RoundedBottomLeft,
	BottomRightCorner: RoundedBottomRight,
	TopBorder:         Hbar,
	BottomBorder:      Hbar,
	LeftBorder:        Vbar,
	RightBorder:       Vbar,
}

var BlockBorder = BorderStyle{
	TopLeftCorner:     BlockFull,
	TopRightCorner:    BlockFull,
	BottomLeftCorner:  BlockFull,
	BottomRightCorner: BlockFull,
	TopBorder:         BlockUpperHalf,
	BottomBorder:      BlockLowerHalf,
	LeftBorder:        BlockLeftHalf,
	RightBorder:       BlockRightHalf,
}

var DoubleBorder = BorderStyle{
	TopLeftCorner:     DoubleTopLeft,
	TopRightCorner:    DoubleTopRight,
	BottomLeftCorner:  DoubleBottomLeft,
	BottomRightCorner: DoubleBottomRight,
	TopBorder:         DoubleHbar,
	BottomBorder:      DoubleHbar,
	LeftBorder:        DoubleVbar,
	RightBorder:       DoubleVbar,
}
