package services

import (
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

// DataBricksJobBar ...
type DataBricksJobBar struct {
	executionTime int
	spotConfig    string
}

// ChartService is the service
type ChartService struct{}

// NewChartService ...
func NewChartService() ChartService {
	return ChartService{}
}

// ConvertToBarItems ...
func (c ChartService) ConvertToBarItems(results []int) []opts.BarData {
	items := make([]opts.BarData, 0)

	itemStyles := opts.ItemStyle{
		Color: "#006666",
	}
	for _, result := range results {
		// values := []DataBricksJobBar{}
		items = append(items, opts.BarData{Value: result, ItemStyle: &itemStyles})
	}

	return items
}

// GenerateJobChart ...
func (c ChartService) GenerateJobChart(dates []string, executionTimes []int) string {
	chartValues := make(map[string][]opts.BarData)
	barData := c.ConvertToBarItems(executionTimes)

	chartValues["Execution Times"] = barData
	fileLocation := RenderBarChart(chartValues, dates)

	return fileLocation
}

// RenderBarChart ...
func RenderBarChart(series map[string][]opts.BarData, xLabels []string) string {
	fileName := "bar.html"

	// create a new bar instance
	bar := charts.NewBar()

	// set some global options like Title/Legend/ToolTip or anything else
	// AxisLabel *Label `json:"axisLabel,omitempty"`
	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "Databricks Job Execution times",
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Type: "value",
			AxisLabel: &opts.Label{
				Formatter: "{value} minutes",
			},
		}),
		charts.WithTooltipOpts(opts.Tooltip{
			Show:      true,
			Formatter: "Hello {a} {b} {c} {d}",
		}),
	)

	// iowriter
	f, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}

	bar.SetXAxis(xLabels)

	// Put some data in instance
	for key, item := range series {
		bar.AddSeries(key, item)
	}

	// Where the magic happens
	bar.Render(f)

	return fileName
}
