package main

import (
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"go-results-visualiser/models"
)

func createExecutionTimeGraph(scenarioNames []string, resultsMap map[string]models.Scenario) *charts.Bar {
	bar := createBarWithTitle(
		"Execution time",
		"Execution time measured in seconds (less is better)",
	)

	for _, scenarioName := range scenarioNames {

		scenarios := resultsMap[scenarioName]

		items := make([]opts.BarData, 0)
		for _, scenario := range scenarios.Stats {
			items = append(items,
				opts.BarData{
					Name:  scenario.ServerName,
					Value: scenario.Stats.TimeTaken,
				})
		}

		bar.AddSeries(scenarioName, items)
	}
	return bar
}

func createMeanLatencyGraph(scenarioNames []string, resultsMap map[string]models.Scenario) *charts.Bar {
	bar := createBarWithTitle(
		"Mean Latency",
		"Mean latency measured in ms (less is better)",
	)
	for _, scenarioName := range scenarioNames {

		scenarios := resultsMap[scenarioName]

		items := make([]opts.BarData, 0)
		for _, scenario := range scenarios.Stats {
			items = append(items,
				opts.BarData{
					Name:  scenario.ServerName,
					Value: models.UnitToFloat(scenario.Stats.LatencyMean),
				})
		}

		bar.AddSeries(scenarioName, items)
	}
	return bar
}

func createMaxLatencyGraph(scenarioNames []string, resultsMap map[string]models.Scenario) *charts.Bar {
	bar := createBarWithTitle(
		"Max Latency",
		"Max latency measured in ms (less is better)",
	)
	for _, scenarioName := range scenarioNames {

		scenarios := resultsMap[scenarioName]

		items := make([]opts.BarData, 0)
		for _, scenario := range scenarios.Stats {
			items = append(items,
				opts.BarData{
					Name:  scenario.ServerName,
					Value: models.UnitToFloat(scenario.Stats.LatencyMax),
				})
		}

		bar.AddSeries(scenarioName, items)
	}
	return bar
}

func createMeanReqGraph(scenarioNames []string, resultsMap map[string]models.Scenario) *charts.Bar {
	bar := createBarWithTitle(
		"Mean Req/s",
		"Mean number of requests executed in 1 second",
	)
	for _, scenarioName := range scenarioNames {

		scenarios := resultsMap[scenarioName]

		items := make([]opts.BarData, 0)
		for _, scenario := range scenarios.Stats {
			items = append(items,
				opts.BarData{
					Name:  scenario.ServerName,
					Value: scenario.Stats.ReqMean,
				})
		}

		bar.AddSeries(scenarioName, items)
	}
	return bar
}

func createMaxReqGraph(scenarioNames []string, resultsMap map[string]models.Scenario) *charts.Bar {
	bar := createBarWithTitle(
		"Max Req/s",
		"Max number of requests executed in 1 second",
	)
	for _, scenarioName := range scenarioNames {

		scenarios := resultsMap[scenarioName]

		items := make([]opts.BarData, 0)
		for _, scenario := range scenarios.Stats {
			items = append(items,
				opts.BarData{
					Name:  scenario.ServerName,
					Value: scenario.Stats.ReqMax,
				})
		}

		bar.AddSeries(scenarioName, items)
	}
	return bar
}

func createMinCpuUsageGraph(scenarioNames []string, resultsMap map[string]models.Scenario) *charts.Bar {
	bar := createBarWithTitle(
		"Min Cpu Usage %",
		"Minimum cpu usage during test execution, as reported by `docker stats`",
	)
	for _, scenarioName := range scenarioNames {

		scenarios := resultsMap[scenarioName]

		items := make([]opts.BarData, 0)
		for _, scenario := range scenarios.UsageStats {
			items = append(items,
				opts.BarData{
					Name:  scenario.ServerName,
					Value: scenario.MinCpuUsage,
				})
		}

		bar.AddSeries(scenarioName, items)
	}
	return bar
}

func createMaxCpuUsageGraph(scenarioNames []string, resultsMap map[string]models.Scenario) *charts.Bar {
	bar := createBarWithTitle(
		"Max Cpu Usage %",
		"Minimum cpu usage during test execution as percentage of the total allocated cores, reported by `docker stats`",
	)
	for _, scenarioName := range scenarioNames {

		scenarios := resultsMap[scenarioName]

		items := make([]opts.BarData, 0)
		for _, scenario := range scenarios.UsageStats {
			items = append(items,
				opts.BarData{
					Name:  scenario.ServerName,
					Value: scenario.MaxCpuUsage,
				})
		}

		bar.AddSeries(scenarioName, items)
	}
	return bar
}

func createMinMemoryUsageMiBGraph(scenarioNames []string, resultsMap map[string]models.Scenario) *charts.Bar {
	bar := createBarWithTitle(
		"Min Memory Usage (MiB)",
		"Minimum minimum usage during test execution as MiB of the total allocated memory, reported by `docker stats`",
	)
	for _, scenarioName := range scenarioNames {

		scenarios := resultsMap[scenarioName]

		items := make([]opts.BarData, 0)
		for _, scenario := range scenarios.UsageStats {
			items = append(items,
				opts.BarData{
					Name:  scenario.ServerName,
					Value: scenario.MinMemoryUsageMiB,
				})
		}

		bar.AddSeries(scenarioName, items)
	}
	return bar
}

func createMaxMemoryUsageMiBGraph(scenarioNames []string, resultsMap map[string]models.Scenario) *charts.Bar {
	bar := createBarWithTitle(
		"Max Memory Usage (MiB)",
		"Minimum minimum usage during test execution as MiB of the total allocated memory, reported by `docker stats`",
	)
	for _, scenarioName := range scenarioNames {

		scenarios := resultsMap[scenarioName]

		items := make([]opts.BarData, 0)
		for _, scenario := range scenarios.UsageStats {
			items = append(items,
				opts.BarData{
					Name:  scenario.ServerName,
					Value: scenario.MaxMemoryUsageMiB,
				})
		}

		bar.AddSeries(scenarioName, items)
	}
	return bar
}

func createBarWithTitle(title string, subtitle string) *charts.Bar {
	bar := generateBarChart(title, subtitle, constServerNames)
	bar.SetGlobalOptions(
		//charts.WithTooltipOpts(opts.Tooltip{Show: true}),
		charts.WithLegendOpts(opts.Legend{Show: true, Right: "80px"}),
	)
	return bar
}

func generateBarChart(title string, subtitle string, labels []string) *charts.Bar {
	bar := charts.NewBar()

	bar.SetGlobalOptions(
		//charts.WithInitializationOpts(opts.Initialization{
		//	Theme: "dark",
		//}),
		charts.WithTitleOpts(opts.Title{Title: title, Subtitle: subtitle}),
	)

	bar.SetXAxis(labels)
	bar.XYReversal()
	return bar
}
