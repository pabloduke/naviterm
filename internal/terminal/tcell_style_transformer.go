package terminal

import (
	"github.com/gdamore/tcell/v3"
	"github.com/gdamore/tcell/v3/color"
	"github.com/pabloduke/naviterm/api"
)

func TransformToTcell(style api.NavitermStyle) tcell.Style {
	tCellStyle := tcell.StyleDefault.
		Foreground(color.Color(style.FgColor.Attr())).
		Background(color.Color(style.BgColor.Attr())).
		Bold(style.Bold).
		Blink(style.Blink).
		Reverse(style.Reverse).
		Underline(style.Underline).
		StrikeThrough(style.Strikethrough).
		Dim(style.Dim)

	return tCellStyle
}
