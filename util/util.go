package util

import (
	"log"
	"os"
)

type NoLogger struct {
}

func (l *NoLogger) Write(p []byte) (n int, err error) {
	return len(p), nil
}

func StopLogging() {
	log.SetOutput(new(NoLogger))
}

func StartLogging() {
	log.SetOutput(os.Stdout)
}
