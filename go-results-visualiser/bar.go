package main

import (
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func createExecutionTimeGraph(scenarioNames []string, resultsMap map[string]Scenario) *charts.Bar {
	bar := createBarWithTitle(
		"Execution time",
		"Execution time measured in seconds (less is better)",
	)

	for _, scenarioName := range scenarioNames {

		scenarios := resultsMap[scenarioName]

		items := make([]opts.BarData, 0)
		for _, scenario := range scenarios.stats {
			items = append(items,
				opts.BarData{
					Name:  scenario.ServerName,
					Value: scenario.stats.TimeTaken,
				})
		}

		bar.AddSeries(scenarioName, items)
	}
	return bar
}

func createMeanLatencyGraph(scenarioNames []string, resultsMap map[string]Scenario) *charts.Bar {
	bar := createBarWithTitle(
		"Mean Latency",
		"Mean latency measured in ms (less is better)",
	)
	for _, scenarioName := range scenarioNames {

		scenarios := resultsMap[scenarioName]

		items := make([]opts.BarData, 0)
		for _, scenario := range scenarios.stats {
			items = append(items,
				opts.BarData{
					Name:  scenario.ServerName,
					Value: unitToFloat(scenario.stats.LatencyMean),
				})
		}

		bar.AddSeries(scenarioName, items)
	}
	return bar
}

func createMaxLatencyGraph(scenarioNames []string, resultsMap map[string]Scenario) *charts.Bar {
	bar := createBarWithTitle(
		"Max Latency",
		"Max latency measured in ms (less is better)",
	)
	for _, scenarioName := range scenarioNames {

		scenarios := resultsMap[scenarioName]

		items := make([]opts.BarData, 0)
		for _, scenario := range scenarios.stats {
			items = append(items,
				opts.BarData{
					Name:  scenario.ServerName,
					Value: unitToFloat(scenario.stats.LatencyMax),
				})
		}

		bar.AddSeries(scenarioName, items)
	}
	return bar
}

func createMeanReqGraph(scenarioNames []string, resultsMap map[string]Scenario) *charts.Bar {
	bar := createBarWithTitle(
		"Mean Req/s",
		"Mean number of requests executed in 1 second",
	)
	for _, scenarioName := range scenarioNames {

		scenarios := resultsMap[scenarioName]

		items := make([]opts.BarData, 0)
		for _, scenario := range scenarios.stats {
			items = append(items,
				opts.BarData{
					Name:  scenario.ServerName,
					Value: scenario.stats.ReqMean,
				})
		}

		bar.AddSeries(scenarioName, items)
	}
	return bar
}

func createMaxReqGraph(scenarioNames []string, resultsMap map[string]Scenario) *charts.Bar {
	bar := createBarWithTitle(
		"Max Req/s",
		"Max number of requests executed in 1 second",
	)
	for _, scenarioName := range scenarioNames {

		scenarios := resultsMap[scenarioName]

		items := make([]opts.BarData, 0)
		for _, scenario := range scenarios.stats {
			items = append(items,
				opts.BarData{
					Name:  scenario.ServerName,
					Value: scenario.stats.ReqMax,
				})
		}

		bar.AddSeries(scenarioName, items)
	}
	return bar
}

func createBarWithTitle(title string, subtitle string) *charts.Bar {
	bar := generateBarChart(title, subtitle, constServerNames)
	bar.SetGlobalOptions(
		charts.WithTooltipOpts(opts.Tooltip{Show: true}),
		charts.WithLegendOpts(opts.Legend{Show: true, Right: "80px"}),
	)
	return bar
}

func generateBarChart(title string, subtitle string, labels []string) *charts.Bar {
	bar := charts.NewBar()

	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: title, Subtitle: subtitle}),
	)

	bar.SetXAxis(labels)
	bar.XYReversal()
	return bar
}
