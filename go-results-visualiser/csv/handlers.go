package csv

import (
	"encoding/csv"
	"fmt"
	"github.com/gocarina/gocsv"
	"go-results-visualiser/models"
	"go-results-visualiser/util"
	"io"
	"log"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
)

func ReadScenarioStatsCsvFile(filePath string) []models.ScenarioStats {
	f, err := os.OpenFile(filePath, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Println(err)
		}
	}(f)

	results := make([]models.ScenarioStats, 0)

	gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
		r := csv.NewReader(in)
		r.Comma = ','
		return r // Allows use pipe as delimiter
	})

	err = gocsv.UnmarshalFile(f, &results)

	if err != nil {
		log.Fatal("Unable to unmarshal csv, err:", err)
	}

	return results
}

func ReadUsageStatsCsvFile(filePath string) []models.RawUsageStats {
	f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Println(err)
		}
	}(f)

	results := make([]models.RawUsageStats, 0)

	gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
		r := csv.NewReader(in)
		r.Comma = '|'
		return r // Allows use pipe as delimiter
	})

	err = gocsv.UnmarshalWithoutHeaders(f, &results)

	if err != nil {
		log.Fatal("Unable to unmarshal csv")
	}

	return results
}

func FindTestCaseScenarioCsv(testCaseDir string, testCase string, pattern string) models.Scenario {
	tcPath := fmt.Sprintf("%s/%s/", testCaseDir, testCase)
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

	var combinedRecords []models.ServerStats

	for _, tc := range matchingFiles {
		records := ReadScenarioStatsCsvFile(fmt.Sprintf("%s%s", tcPath, tc))
		//fmt.Printf("TestCase: %s, values: %s\n", tc, records)
		if len(records) > 1 {
			panic(fmt.Sprintf("Invalid csv found: %s\n", tc))
		}

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

func FindTestCaseResourceUsageCsv(testCaseDir string, testCase string, pattern string) models.Scenario {
	tcPath := fmt.Sprintf("%s/%s/", testCaseDir, testCase)
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

	var combinedRecords []models.UsageStats
	for _, tc := range matchingFiles {
		rawUsageStats := ReadUsageStatsCsvFile(fmt.Sprintf("%s%s", tcPath, tc))
		cpuUsageArray := make([]float64, 0)
		memoryUsageArray := make([]float64, 0)
		memoryUsageMiBArray := make([]float64, 0)
		for _, t := range rawUsageStats {
			cpuValue := strings.TrimSuffix(t.CpuUsagePercentage, "%")
			if f, err := strconv.ParseFloat(cpuValue, 64); err == nil {
				cpuUsageArray = append(cpuUsageArray, f)
			}

			memoryValue := strings.TrimSuffix(t.MemoryUsagePercentage, "%")
			if f, err := strconv.ParseFloat(memoryValue, 64); err == nil {
				memoryUsageArray = append(memoryUsageArray, f)
			}

			memoryValueMiB := strings.TrimSuffix(t.MemoryUsage, "MiB")
			if f, err := strconv.ParseFloat(memoryValueMiB, 64); err == nil {
				memoryUsageMiBArray = append(memoryUsageMiBArray, f)
			}

		}
		usageStats := models.UsageStats{
			ServerName:        tc,
			MinCpuUsage:       slices.Min(cpuUsageArray),
			MaxCpuUsage:       slices.Max(cpuUsageArray),
			MinMemoryUsage:    slices.Min(memoryUsageArray),
			MaxMemoryUsage:    slices.Max(memoryUsageArray),
			MinMemoryUsageMiB: slices.Min(memoryUsageMiBArray),
			MaxMemoryUsageMiB: slices.Max(memoryUsageMiBArray),
		}
		combinedRecords = append(combinedRecords, usageStats)
	}

	scenario := models.Scenario{
		Name:       testCase,
		UsageStats: combinedRecords,
	}

	return scenario
}
