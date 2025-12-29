package chart

import "github.com/pabloduke/naviterm/data/color"

type BarChart struct {
	Title  string
	XLabel string
	YLabel string
	Items  []BarChartItem
}

type BarChartItem struct {
	Label string
	Value int
	Color color.NaviTermColor
}
