package main

import (
	"bufio"
	"os"
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
	BASE_PATH     = "./external/jmeter/" // NOTE: Run CLI from root folder
	BASE_FILENAME = "rcie-api-no"
	SCRIPT_FOLDER = "script/jmx/"
	LOG_FOLDER    = "log/"
	REPORT_FOLDER = "report/"
	JMX_FILE_TYPE = ".jmx"
	LOG_FILE_TYPE = ".jtl"
	TIME_FORMAT   = "2006-01-02_15-04-05"
)

// TO FUTURE ME: IF STILL FREEZE, TRY TO RUN THE COMMAND WITH SET HEAP SIZE
func main() {
	meta := generateMeta()

	if cli.GetArg(2, "true") == "false" {
		console.App.Log("Skip Running JMeter CLI...")
		console.App.Log("File name", meta.FileName)
		console.App.Log("CLI Command:", "jmeter -n -t", meta.Script, "-l", meta.Log, "-e -o", meta.Report)
		return
	}

	console.App.Log("Start Running JMeter CLI at", meta.ExecutionTime)
	console.App.Log("Running File", meta.FileName)
	console.App.Log("Script Location", meta.Script)
	console.App.Log("Log Location", meta.Log)
	console.App.Log("Report Location", meta.Report)

	cmd := exec.Command("jmeter", "-n", "-t", meta.Script, "-l", meta.Log, "-e", "-o", meta.Report)

	console.App.Log("Running Command", cmd.String())

	cmdReader, err := cmd.StdoutPipe()
	if err != nil {
		console.App.Fatal(os.Stderr, "Error Creating StdoutPipe for CMD", err)
		return
	}

	scanner := bufio.NewScanner(cmdReader)
	go func() {
		for scanner.Scan() {
			console.App.Logf("\t > %s\n", scanner.Text())
		}
	}()

	err = cmd.Start()
	if err != nil {
		console.App.Log(os.Stderr, "Error Starting CMD", err)
		return
	}

	err = cmd.Wait()
	if err != nil {
		console.App.Log(os.Stderr, "Error Waiting for CMD", err)
		return
	}

	console.App.Log("Done Running JMeter CLI with", meta.FileNameWithTime)
}

func generateMeta() *FileMeta {
	currentTime := time.Now().Format(TIME_FORMAT)
	logFileType := env.GetEnv("LOG_TYPE", LOG_FILE_TYPE)
	api := cli.GetArg(1, "1")
	fileName := BASE_FILENAME + api
	fileNameWithTime := currentTime + "_" + fileName

	return &FileMeta{
		ExecutionTime:    currentTime,
		FileName:         fileName,
		FileNameWithTime: fileNameWithTime,
		Script:           BASE_PATH + SCRIPT_FOLDER + fileName + JMX_FILE_TYPE,
		Log:              BASE_PATH + LOG_FOLDER + fileNameWithTime + logFileType,
		Report:           BASE_PATH + REPORT_FOLDER + fileNameWithTime,
	}
}
