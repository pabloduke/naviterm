package symbols

//Horizontal: ━ (U+2501)
//Vertical: ┃ (U+2503)
//Top-left: ┏ (U+250F)
//Top-right: ┓ (U+2513)
//Bottom-left: ┗ (U+2517)
//Bottom-right: ┛ (U+251B)

const (
	Hbar        = "─"
	Vbar        = "│"
	TopLeft     = "┌"
	TopRight    = "┐"
	BottomLeft  = "└"
	BottomRight = "┘"

	// Double Line
	DoubleHbar        = "═"
	DoubleVbar        = "║"
	DoubleTopLeft     = "╔"
	DoubleTopRight    = "╗"
	DoubleBottomLeft  = "╚"
	DoubleBottomRight = "╝"
	DoubleTeeLeft     = "╠"
	DoubleTeeRight    = "╣"
	DoubleTeeDown     = "╦"
	DoubleTeeUp       = "╩"
	DoubleCross       = "╬"

	// Heavy Line
	HeavyHbar        = "━"
	HeavyVbar        = "┃"
	HeavyTopLeft     = "┏"
	HeavyTopRight    = "┓"
	HeavyBottomLeft  = "┗"
	HeavyBottomRight = "┛"
	HeavyTeeLeft     = "┣"
	HeavyTeeRight    = "┫"
	HeavyTeeDown     = "┳"
	HeavyTeeUp       = "┻"
	HeavyCross       = "╋"

	// Rounded
	RoundedTopLeft     = "╭"
	RoundedTopRight    = "╮"
	RoundedBottomLeft  = "╰"
	RoundedBottomRight = "╯"

	// Mixed: Double Horizontal, Single Vertical
	MixedDHSVTopLeft     = "╒"
	MixedDHSVTopRight    = "╕"
	MixedDHSVBottomLeft  = "╘"
	MixedDHSVBottomRight = "╛"

	// Mixed: Single Horizontal, Double Vertical
	MixedSHDVTopLeft     = "╓"
	MixedSHDVTopRight    = "╖"
	MixedSHDVBottomLeft  = "╙"
	MixedSHDVBottomRight = "╜"

	// Dashed Light
	DashedLightHbar = "┄"
	DashedLightVbar = "┆"
	DottedLightHbar = "┈"
	DottedLightVbar = "┊"

	// Dashed Heavy
	DashedHeavyHbar = "┅"
	DashedHeavyVbar = "┇"
	DottedHeavyHbar = "┉"
	DottedHeavyVbar = "┋"

	// Block Elements
	BlockFull        = "█"
	BlockUpperHalf   = "▀"
	BlockLowerHalf   = "▄"
	BlockLeftHalf    = "▌"
	BlockRightHalf   = "▐"
	BlockShadeLight  = "░"
	BlockShadeMedium = "▒"
	BlockShadeDark   = "▓"
)
