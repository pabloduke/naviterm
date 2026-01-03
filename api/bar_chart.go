package api

type BarChart struct {
	Title   string
	XLabel  string
	YLabel  string
	Spacing int
	Items   []BarChartItem
}

type BarChartItem struct {
	Label string
	Value int
	Color NaviTermColor
}
