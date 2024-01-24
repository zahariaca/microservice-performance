package csv

import (
	"fmt"
	"github.com/gocarina/gocsv"
	"go-results-visualiser/models"
	"go-results-visualiser/util"
	"log"
	"os"
	"path/filepath"
)

func ReadCsvFile(filePath string) []models.ScenarioStats {
	f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	results := []models.ScenarioStats{}

	err = gocsv.UnmarshalFile(f, &results)

	if err != nil {
		log.Fatal("Unable to unmarshal csv")
	}

	return results
}

func FindTestCaseCsv(testCaseDir string, testCase string) models.Scenario {
	tcPath := fmt.Sprintf("%s/%s/", testCaseDir, testCase)
	pattern := "*-req.csv"
	files, err := os.ReadDir(tcPath)
	util.CheckForError(err)

	var matchingFiles []string
	for _, f := range files {
		match, err := filepath.Match(pattern, f.Name())
		util.CheckForError(err)

		if match {
			matchingFiles = append(matchingFiles, f.Name())
		}
	}

	resultsMap := make(map[string]models.ScenarioStats)

	var combinedRecords []models.ServerStats

	for _, tc := range matchingFiles {
		records := ReadCsvFile(fmt.Sprintf("%s%s", tcPath, tc))
		//fmt.Printf("TestCase: %s, values: %s\n", tc, records)
		if len(records) > 1 {
			panic(fmt.Sprintf("Invalid csv found: %s\n", tc))
		}

		resultsMap[tc] = records[0]
		record := records[0]
		combinedRecords = append(combinedRecords, models.ServerStats{
			ServerName: tc,
			Stats:      &record,
		})
	}

	scenario := models.Scenario{
		Name:  testCase,
		Stats: combinedRecords,
	}

	return scenario
}
