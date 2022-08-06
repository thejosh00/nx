package util

type NoLogger struct {
}

func (l *NoLogger) Write(p []byte) (n int, err error) {
	return len(p), nil
}
