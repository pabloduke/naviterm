package render

import (
	"github.com/pabloduke/naviterm/data/chart"
	"github.com/pabloduke/naviterm/internal/terminal"
)

func DrawBarChart(barChart chart.BarChart) {
	for x, item := range barChart.Items {
		DrawRectangle(x*2, 10, 4, item.Value)
	}

	terminal.Flush()
}
