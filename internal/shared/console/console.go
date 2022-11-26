package console

import (
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

func (s *Console) Log(message string) {
	s.Logger.Println(message)
}

func (s *Console) Logf(message string, options ...any) {
	s.Logger.Printf(message, options...)
}

func (s *Console) Fatal(message string) {
	s.ErrorLogger.Fatal(message)
}

func (s *Console) Fatalf(message string, options ...any) {
	s.ErrorLogger.Fatalf(message, options...)
}
