package util

import "log"

type NoLogger struct {
}

func (l *NoLogger) Write(p []byte) (n int, err error) {
	return len(p), nil
}

func StopLogging() {
	log.SetOutput(new(NoLogger))
}
