package main

import (
	"os/exec"
	"time"

	"github.com/nuttchai/go-rest/internal/shared/console"
	"github.com/nuttchai/go-rest/internal/utils/cli"
	"github.com/nuttchai/go-rest/internal/utils/env"
)

type FileMeta struct {
	ExecutionTime    string
	FileName         string
	FileNameWithTime string
	Script           string
	Log              string
	Report           string
}

const (
	BASE_PATH     = "D:/nuttchai/dev/Projects/rcie/jmeter/rcie/"
	BASE_FILENAME = "rcie-api-no"
	SCRIPT_FOLDER = "script/jmx/"
	LOG_FOLDER    = "log/"
	REPORT_FOLDER = "report/"
)

func main() {
	console.App.Log("Start Running JMeter CLI...")

	meta := generateMeta()

	console.App.Log("Running File", meta.FileName)
	console.App.Log("Running Time", meta.ExecutionTime)
	console.App.Log("Script Location", meta.Script)
	console.App.Log("Log Location", meta.Log)
	console.App.Log("Report Location", meta.Report)

	cmd := exec.Command("jmeter", "-n", "-t", meta.Script, "-l", meta.Log, "-e", "-o", meta.Report)

	console.App.Log("Running Command", cmd.String())

	if err := cmd.Run(); err != nil {
		console.App.Fatal("Failed to Execute Command", err)
	}

	console.App.Log("Done Running JMeter CLI with", meta.FileNameWithTime)
}

func generateMeta() *FileMeta {
	currentTime := time.Now().Format("2006:01:02-15:04:05")
	logFileType := env.GetEnv("LOG_TYPE", "csv")
	api := cli.GetArg(1, "1")
	fileName := BASE_FILENAME + api
	fileNameWithTime := currentTime + "-" + fileName

	return &FileMeta{
		ExecutionTime:    currentTime,
		FileName:         fileName,
		FileNameWithTime: fileNameWithTime,
		Script:           BASE_PATH + SCRIPT_FOLDER + fileName + ".jmx",
		Log:              BASE_PATH + LOG_FOLDER + fileNameWithTime + "." + logFileType,
		Report:           BASE_PATH + REPORT_FOLDER + fileNameWithTime,
	}
}
