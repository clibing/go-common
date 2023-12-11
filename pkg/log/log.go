package log

import (
	"io"

	"github.com/sirupsen/logrus"
)

var (
	logger = logrus.New()
)

// // Hook is a hook that writes logs of specified LogLevels to specified Writer
type Hook struct {
	Writer    io.Writer
	LogLevels []logrus.Level
}

// Fire will be called when some logging function is called with current hook
// It will format log entry to string and write it to appropriate writer
func (hook *Hook) Fire(entry *logrus.Entry) error {
	line, err := entry.Bytes()
	if err != nil {
		return err
	}
	_, err = hook.Writer.Write(line)
	return err
}

// Levels define on which log levels this hook would trigger
func (hook *Hook) Levels() []logrus.Level {
	return hook.LogLevels
}

type AddHook interface {
}

func LogLevel(level string) *logrus.Logger {
	if level == "debug" {
		logger.Level = logrus.DebugLevel
	} else if level == "info" {
		logger.Level = logrus.InfoLevel
	} else if level == "trace" {
		logger.Level = logrus.TraceLevel
	} else if level == "warn" {
		logger.Level = logrus.WarnLevel
	} else if level == "error" {
		logger.Level = logrus.ErrorLevel
	}
	return logger
}

func Logger() *logrus.Logger {
	return logger
}
