package console

import (
	"fmt"
	"log"
	"os"
)

type Console struct {
	Logger      *log.Logger
	ErrorLogger *log.Logger
}

var App *Console

func init() {
	App = &Console{
		Logger:      log.New(os.Stdout, "APP_LOG: ", log.Ldate|log.Ltime),
		ErrorLogger: log.New(os.Stdout, "APP_ERROR_LOG: ", log.Ldate|log.Ltime),
	}
}

func (s *Console) Log(messages ...any) {
	var logMsg string
	for index, message := range messages {
		if index == 0 {
			logMsg += fmt.Sprintf("%v", message)
			continue
		}
		logMsg += fmt.Sprintf(" %v", message)
	}

	s.Logger.Println(logMsg)
}

func (s *Console) Logf(message string, options ...any) {
	s.Logger.Printf(message, options...)
}

func (s *Console) Fatal(messages ...any) {
	var fatalMsg string
	for index, message := range messages {
		if index == 0 {
			fatalMsg += fmt.Sprintf("%v", message)
			continue
		}
		fatalMsg += fmt.Sprintf(" %v", message)
	}

	s.ErrorLogger.Fatal(fatalMsg)
}

func (s *Console) Fatalf(message string, options ...any) {
	s.ErrorLogger.Fatalf(message, options...)
}
