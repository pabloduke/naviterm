package render

import (
	"github.com/nsf/termbox-go"
	"github.com/pabloduke/naviterm/data/chart"
	"github.com/pabloduke/naviterm/data/symbols"
	"github.com/pabloduke/naviterm/internal/terminal"
)

func DrawBarChart(x int, y int, barChart chart.BarChart) {
	height := calcChartHeight(barChart.Items)
	drawBorder(x, y-1, barChart)
	drawTitle(x, y, height, barChart)
	drawXLabel(x, y, height, barChart)
	for index, item := range barChart.Items {
		DrawRectangle(index+(index*2)+x, y-2, 1, item.Value, item.Color)
	}

	terminal.Flush()
}

func calcChartWidth(barChart chart.BarChart) int {
	return len(barChart.Items) * barChart.Spacing
}

func calcChartHeight(items []chart.BarChartItem) int {
	height := items[0].Value
	for _, i := range items {
		if i.Value > height {
			height = i.Value
		}
	}

	return height
}

func drawBorder(x int, y int, barChart chart.BarChart) {
	width := calcChartWidth(barChart)
	height := calcChartHeight(barChart.Items)

	var XPos = x - 2

	for ; XPos < width+x+5; XPos++ {
		DrawText(XPos, y, symbols.Hbar, termbox.ColorRed, termbox.ColorDefault)
		DrawText(XPos, y-height-1, symbols.Hbar, termbox.ColorLightCyan, termbox.ColorDefault)
	}
}

func drawTitle(x int, y int, height int, barChart chart.BarChart) {
	DrawText(x, y-height-2, barChart.Title, termbox.ColorWhite, termbox.ColorDefault)
}

func drawXLabel(x int, y int, height int, barChart chart.BarChart) {
	DrawText(x-10, y-(height/2), barChart.XLabel, termbox.ColorWhite, termbox.ColorDefault)
}
