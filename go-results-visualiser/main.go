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

func firstPage(w http.ResponseWriter, _ *http.Request) {
	page := components.NewPage()

	testCaseDir := "../project/4cpu/get"
	resultsMap, scenarioNames := buildScenariosMap(testCaseDir)

	log.Println("ScenarioNames: ", scenarioNames)

	page.AddCharts(
		createExecutionTimeGraph(scenarioNames, resultsMap),
		createMeanLatencyGraph(scenarioNames, resultsMap),
		createMaxLatencyGraph(scenarioNames, resultsMap),
		createMeanReqGraph(scenarioNames, resultsMap),
		createMaxReqGraph(scenarioNames, resultsMap),
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

func main() {

	http.HandleFunc("/", httpserver)
	http.HandleFunc("/firstPage", firstPage)
	err := http.ListenAndServe(":7000", nil)
	if err != nil {
		return
	}
}

func buildScenariosMap(testCaseDir string) (map[string]models.Scenario, []string) {
	testCases := findTestCaseDirectories(testCaseDir)

	log.Println(testCases)

	resultsMap := make(map[string]models.Scenario)
	scenarioNames := make([]string, 0)
	for _, tc := range testCases {
		scenario := csv.FindTestCaseCsv(testCaseDir, tc)
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
