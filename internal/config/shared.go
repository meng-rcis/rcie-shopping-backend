package config

import (
	"log"
	"os"
)

type Shared struct {
	Logger      *log.Logger
	ErrorLogger *log.Logger
}

var App *Shared

func init() {
	App = &Shared{
		Logger:      log.New(os.Stdout, "APP_LOG: ", log.Ldate|log.Ltime),
		ErrorLogger: log.New(os.Stdout, "APP_ERROR_LOG: ", log.Ldate|log.Ltime),
	}
}

func (s *Shared) Log(message string) {
	s.Logger.Println(message)
}

func (s *Shared) Logf(message string, options ...any) {
	s.Logger.Printf(message, options...)
}

func (s *Shared) Fatal(message string) {
	s.ErrorLogger.Fatal(message)
}

func (s *Shared) Fatalf(message string, options ...any) {
	s.ErrorLogger.Fatalf(message, options...)
}
