package main

import (
	"bufio"
	"encoding/csv"
	"io"
	"os"

	"github.com/nuttchai/go-rest/internal/shared/console"
	"github.com/nuttchai/go-rest/internal/utils/cli"
)

const (
	BASE_FILENAME = "rcie-api-"
	BASE_PATH     = "./external/jmeter"

	LOG_FOLDER     = "/log/"
	GRAFANA_FOLDER = "/grafana/"

	CPU_FILE     = "cpu"
	MEMORY_FILE  = "memory"
	NETWORK_FILE = "network"
	FILE_TYPE    = ".csv"
)

func main() {
	dateTime := cli.GetArg(1, "")
	if dateTime == "" {
		console.App.Fatal("Please provide date time")
		return
	}

	baseName := getBaseName(dateTime)
	grafana := BASE_PATH + GRAFANA_FOLDER + baseName + "/"
	logPath := BASE_PATH + LOG_FOLDER + baseName + FILE_TYPE
	cpuPath := grafana + CPU_FILE + FILE_TYPE
	memoryPath := grafana + MEMORY_FILE + FILE_TYPE
	networkPath := grafana + NETWORK_FILE + FILE_TYPE

	if !verifyPath(logPath, cpuPath, memoryPath, networkPath) {
		return
	}

	console.App.Log("Base Name: " + baseName)
	console.App.Log("Log File Path: " + logPath)
	console.App.Log("CPU File Path: " + cpuPath)
	console.App.Log("Memory File Path: " + memoryPath)
	console.App.Log("Network File Path: " + networkPath)

	readFile(cpuPath)
}

func extractFile(basePath string, paths ...string) ([][]string, error) {
	timestamp := []string{}
	baseFile, err := os.Open(basePath)

	if err != nil {
		return nil, err
	}
	defer baseFile.Close()

	r := csv.NewReader(bufio.NewReader(baseFile))
	r.Comma = ','
	r.LazyQuotes = true

	// NOTE: Skip first line
	if _, err = r.Read(); err != nil {
		return nil, err
	}

	// NOTE: Store timestamp
	for {
		record, err := r.Read()

		if err == io.EOF {
			break
		}
		if err != nil {
			console.App.Fatal(err)
		}
		timestamp = append(timestamp, record[0])
	}

	columns := []string{}
	for _, path := range paths {
		file, err := os.Open(path)

		if err != nil {
			return nil, err
		}
		defer file.Close()

		r := csv.NewReader(bufio.NewReader(file))
		r.Comma = ','
		r.LazyQuotes = true

		// NOTE: Store column name
		columnRecord, err := r.Read()
		if len(columns) == 0 {
			columns = columnRecord
		} else {
			columnsRecordWithoutTime := columnRecord[1:]
			columns = append(columns, columnsRecordWithoutTime...)
		}

		// NOTE: Store data
	}
	return nil, nil
}

func readFile(path string) string {
	file, err := os.Open(path)

	if err != nil {
		console.App.Fatal(err)
		return ""
	}
	defer file.Close()

	r := csv.NewReader(bufio.NewReader(file))
	r.Comma = ','
	r.LazyQuotes = true

	for {
		record, err := r.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			console.App.Fatal(err)
		}

		console.App.Log("record:", record)
		for value := range record {
			console.App.Log("value:", value)
			console.App.Log("record-value:", record[value])
		}
	}

	return ""
}

func getBaseName(dateTime string) string {
	fileName := dateTime + "_" + BASE_FILENAME
	api := cli.GetArg(2, "default")
	if api == "default" {
		fileName += "default"
	} else {
		fileName += "no" + api
	}

	return fileName
}

func isPathExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func verifyPath(paths ...string) bool {
	for _, path := range paths {
		if isExist, _ := isPathExist(path); !isExist {
			console.App.Fatal("Path not found: " + path)
			return false
		}
	}

	return true
}
