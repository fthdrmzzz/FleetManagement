package logger

import (
	"log"
	"os"
)

type Logging struct {
	Info  *log.Logger
	Warn  *log.Logger
	Error *log.Logger
}

func New() Logging {
	flags := log.LstdFlags | log.Lshortfile
	log := Logging{
		Info:  log.New(os.Stdout, "INFO: ", flags),
		Warn:  log.New(os.Stdout, "WARN: ", flags),
		Error: log.New(os.Stdout, "ERROR: ", flags),
	}
	return log
}
