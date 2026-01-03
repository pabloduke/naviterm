package render

import (
	"github.com/pabloduke/naviterm/api"
	"github.com/pabloduke/naviterm/internal/terminal"
)

func DrawBarChart(x int, y int, barChart api.BarChart) {
	height := calcChartHeight(barChart.Items)
	drawBorder(x, y-1, barChart)
	drawTitle(x, y, height, barChart)
	drawXLabel(x, y, height, barChart)
	for index, item := range barChart.Items {
		DrawRectangle(index+(index*2)+x, y-2, 1, item.Value, item.Color)
	}

	terminal.Flush()
}

func calcChartWidth(barChart api.BarChart) int {
	return len(barChart.Items) * barChart.Spacing
}

func calcChartHeight(items []api.BarChartItem) int {
	height := items[0].Value
	for _, i := range items {
		if i.Value > height {
			height = i.Value
		}
	}

	return height
}

func drawBorder(x int, y int, barChart api.BarChart) {
	width := calcChartWidth(barChart)
	height := calcChartHeight(barChart.Items)

	var XPos = x - 2

	for ; XPos < width+x+5; XPos++ {
		DrawText(XPos, y, api.Hbar, terminal.Attribute(api.ColorRed), terminal.Attribute(api.ColorDefault))
		DrawText(XPos, y-height-1, api.Hbar, terminal.Attribute(api.ColorLightCya), terminal.Attribute(api.ColorDefault))
	}
}

func drawTitle(x int, y int, height int, barChart api.BarChart) {
	DrawText(x, y-height-2, barChart.Title, terminal.Attribute(api.ColorWhite), terminal.Attribute(api.ColorDefault))
}

func drawXLabel(x int, y int, height int, barChart api.BarChart) {
	DrawText(x-10, y-(height/2), barChart.XLabel, terminal.Attribute(api.ColorWhite), terminal.Attribute(api.ColorDefault))
}
