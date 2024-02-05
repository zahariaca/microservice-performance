package main

import (
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
	"go-results-visualiser/csv"
	"go-results-visualiser/models"
	"go-results-visualiser/util"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"sort"
)

var (
	constServerNames = []string{"Go+Gin", "SB No VT", "SB VT", "SB Webflux"}
)

// generate random data for line chart
func generateLineItems() []opts.LineData {
	items := make([]opts.LineData, 0)
	for i := 0; i < 7; i++ {
		items = append(items, opts.LineData{Value: rand.Intn(300)})
	}
	return items
}

func httpserver(w http.ResponseWriter, _ *http.Request) {
	// create a new line instance
	line := charts.NewLine()
	// set some global options like Title/Legend/ToolTip or anything else
	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros}),
		charts.WithTitleOpts(opts.Title{
			Title:    "Line example in Westeros theme",
			Subtitle: "Line chart rendered by the http server this time",
		}))

	// Put data into instance
	line.SetXAxis([]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}).
		AddSeries("Category A", generateLineItems()).
		AddSeries("Category B", generateLineItems()).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: true}))

	err := line.Render(w)
	util.CheckForError(err)
}

func main() {
	http.HandleFunc("/", httpserver)
	http.HandleFunc("/4-cpu-get-results", fourCpuGetResults)
	http.HandleFunc("/auto-4-cpu-get-results", autoFourCpuGetResults)
	err := http.ListenAndServe(":7000", nil)
	if err != nil {
		return
	}
}

func fourCpuGetResults(w http.ResponseWriter, _ *http.Request) {
	//charts.WithInitializationOpts(opts.Initialization{
	//	Theme: "dark",
	//})
	page := components.NewPage()
	page.PageTitle = "4 CPU Get Request Results"

	testCasesDir := "../project/4cpu/get"
	testCases := findTestCaseDirectories(testCasesDir)
	log.Println("Test cases:", testCases)
	resultsMap, scenarioNames := buildScenariosMap(testCasesDir, testCases)
	usageResults, usageScenarioNames := buildResourceUsageMap(testCasesDir, testCases)

	page.AddCharts(
		createExecutionTimeGraph(scenarioNames, resultsMap),
		createMeanLatencyGraph(scenarioNames, resultsMap),
		createMaxLatencyGraph(scenarioNames, resultsMap),
		createMeanReqGraph(scenarioNames, resultsMap),
		createMaxReqGraph(scenarioNames, resultsMap),
		createMinCpuUsageGraph(usageScenarioNames, usageResults),
		createMaxCpuUsageGraph(usageScenarioNames, usageResults),
		createMinMemoryUsageMiBGraph(usageScenarioNames, usageResults),
		createMaxMemoryUsageMiBGraph(usageScenarioNames, usageResults),
	)

	f, err := os.Create("../docs/get-requests-results.html")

	if err != nil {
		panic(err)
	}
	err = page.Render(io.MultiWriter(f))
	if err != nil {
		return
	}

	err = page.Render(w)
	if err != nil {
		return
	}
}
func autoFourCpuGetResults(w http.ResponseWriter, _ *http.Request) {
	//charts.WithInitializationOpts(opts.Initialization{
	//	Theme: "dark",
	//})
	page := components.NewPage()
	page.PageTitle = "4 CPU Get Request Results"

	testCasesDir := "../project/automated/4cpu/get-redo"
	testCases := findTestCaseDirectories(testCasesDir)
	log.Println("Test cases:", testCases)
	resultsMap, scenarioNames := buildScenariosMap(testCasesDir, testCases)
	usageResults, usageScenarioNames := buildResourceUsageMap(testCasesDir, testCases)

	page.AddCharts(
		createExecutionTimeGraph(scenarioNames, resultsMap),
		createMeanLatencyGraph(scenarioNames, resultsMap),
		createMaxLatencyGraph(scenarioNames, resultsMap),
		createMeanReqGraph(scenarioNames, resultsMap),
		createMaxReqGraph(scenarioNames, resultsMap),
		createMinCpuUsageGraph(usageScenarioNames, usageResults),
		createMaxCpuUsageGraph(usageScenarioNames, usageResults),
		createMinMemoryUsageMiBGraph(usageScenarioNames, usageResults),
		createMaxMemoryUsageMiBGraph(usageScenarioNames, usageResults),
	)

	f, err := os.Create("../docs/auto-4cpu-get-requests-results.html")

	if err != nil {
		panic(err)
	}
	err = page.Render(io.MultiWriter(f))
	if err != nil {
		return
	}

	err = page.Render(w)
	if err != nil {
		return
	}
}
func buildScenariosMap(testCasesDir string, testCases []string) (map[string]models.Scenario, []string) {
	pattern := "*-req.csv"

	resultsMap := make(map[string]models.Scenario)
	scenarioNames := make([]string, 0)
	for _, tc := range testCases {
		scenario := csv.FindTestCaseScenarioCsv(testCasesDir, tc, pattern)
		resultsMap[scenario.Name] = scenario
		scenarioNames = append(scenarioNames, scenario.Name)
	}
	sort.Strings(scenarioNames)

	return resultsMap, scenarioNames
}

func buildResourceUsageMap(testCasesDir string, testCases []string) (map[string]models.Scenario, []string) {
	pattern := "*-stats.csv"

	resultsMap := make(map[string]models.Scenario)
	scenarioNames := make([]string, 0)
	for _, tc := range testCases {
		scenario := csv.FindTestCaseResourceUsageCsv(testCasesDir, tc, pattern)
		resultsMap[scenario.Name] = scenario
		scenarioNames = append(scenarioNames, scenario.Name)
	}
	sort.Strings(scenarioNames)

	return resultsMap, scenarioNames
}

func findTestCaseDirectories(testCaseDir string) []string {
	files, err := os.ReadDir(testCaseDir)

	if err != nil {
		log.Fatal(err)
	}

	fileNames := make([]string, 0, 0)
	for _, file := range files {
		log.Println(file.Name())
		fileNames = append(fileNames, file.Name())
	}

	return fileNames
}
