package main

import (
	"fmt"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
	"github.com/gocarina/gocsv"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"sort"
)

var (
	constServerNames = []string{"Go+Gin", "SB No VT", "SB VT", "SB Webflux"}
)

type ScenarioStats struct {
	NrOfConn    uint64  `csv:"Connections" json:"nr-of-connections"`
	NrOfReq     uint64  `csv:"NrOfReq" json:"nr-of-req"`
	TimeTaken   float64 `csv:"TimeTaken" json:"time-taken"`
	LatencyMean float64 `csv:"Latency.Mean" json:"latency-mean"`
	LatencyMax  float64 `csv:"Latency.Max" json:"latency-max"`
	ReqMean     float64 `csv:"Req.Mean" json:"req-mean"`
	ReqMax      float64 `csv:"Req.Max" json:"req-max"`
	Req2xx      uint64  `csv:"Req2XX" json:"req2xx"`
	Others      uint64  `csv:"Other" json:"other"`
}

func (scenarioStats *ScenarioStats) String() string {
	return fmt.Sprintf("%d %d %s %s %s %f %f %d %d",
		scenarioStats.NrOfConn,
		scenarioStats.NrOfReq,
		formatTimeUs(scenarioStats.TimeTaken),
		formatTimeUs(scenarioStats.LatencyMean),
		formatTimeUs(scenarioStats.LatencyMax),
		scenarioStats.ReqMean,
		scenarioStats.ReqMax,
		scenarioStats.Req2xx,
		scenarioStats.Others)
}

type ServerStats struct {
	ServerName string
	stats      *ScenarioStats
}

func (serverStat ServerStats) String() string {
	return fmt.Sprintf("ServerName: %s, stats: %v",
		serverStat.ServerName,
		serverStat.stats)
}

type Scenario struct {
	Name  string
	stats []ServerStats
}

func (scenario Scenario) String() string {
	return fmt.Sprintf("ScenarioName: %s, Stats: %v", scenario.Name, scenario.stats)
}

func readCsvFile(filePath string) []ScenarioStats {
	f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	results := []ScenarioStats{}

	err = gocsv.UnmarshalFile(f, &results)

	if err != nil {
		log.Fatal("Unable to unmarshal csv")
	}

	return results
}

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
	line.Render(w)
}

func firstPage(w http.ResponseWriter, _ *http.Request) {
	page := components.NewPage()

	resultsMap, scenarioNames := buildScenariosMap()

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

func buildScenariosMap() (map[string]Scenario, []string) {
	testCaseDir := "../project/4cpu/get"
	testCases := findTestCaseDirectories(testCaseDir)

	log.Println(testCases)

	resultsMap := make(map[string]Scenario)
	scenarioNames := make([]string, 0)
	for _, tc := range testCases {
		scenario := findTestCaseCsv(testCaseDir, tc)
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

func findTestCaseCsv(testCaseDir string, testCase string) Scenario {
	tcPath := fmt.Sprintf("%s/%s/", testCaseDir, testCase)
	pattern := "*-req.csv"
	files, err := os.ReadDir(tcPath)
	CheckForError(err)

	var matchingFiles []string
	for _, f := range files {
		match, err := filepath.Match(pattern, f.Name())
		CheckForError(err)

		if match {
			matchingFiles = append(matchingFiles, f.Name())
		}
	}

	resultsMap := make(map[string]ScenarioStats)

	var combinedRecords []ServerStats

	for _, tc := range matchingFiles {
		records := readCsvFile(fmt.Sprintf("%s%s", tcPath, tc))
		//fmt.Printf("TestCase: %s, values: %s\n", tc, records)
		if len(records) > 1 {
			panic(fmt.Sprintf("Invalid csv found: %s\n", tc))
		}

		resultsMap[tc] = records[0]
		record := records[0]
		combinedRecords = append(combinedRecords, ServerStats{
			ServerName: tc,
			stats:      &record,
		})
	}

	scenario := Scenario{
		Name:  testCase,
		stats: combinedRecords,
	}

	return scenario
}

func CheckForError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
