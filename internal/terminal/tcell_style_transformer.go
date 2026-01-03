package terminal

import (
	"github.com/gdamore/tcell/v3"
	"github.com/pabloduke/naviterm/api"
)

func TransformToTcell(style api.NavitermStyle) tcell.Style {
	tCellStyle := tcell.StyleDefault.
		Foreground(convertToTcellColor(Attribute(style.FgColor.Attr()))).
		Background(convertToTcellColor(Attribute(style.BgColor.Attr()))).
		Bold(style.Bold).
		Blink(style.Blink).
		Reverse(style.Reverse).
		Underline(style.Underline).
		StrikeThrough(style.Strikethrough).
		Dim(style.Dim)

	return tCellStyle
}
