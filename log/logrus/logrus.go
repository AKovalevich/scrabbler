package logrus

import (
	"github.com/AKovalevich/scrabbler/log"
	"github.com/Sirupsen/logrus"
)

var (
	Do	log.Logger
)

func init() {
	Do = createNew(&log.Config{
		Level: log.LevelDebug,
		Time:  true,
		UTC:   true,
	})
}

func createNew(config *log.Config) log.Logger {
	logger := logrus.New()
	logger.Level = logrusLevelConverter(config.Level)
	logger.WithFields(logrus.Fields(config.Fields))
	return logger
}

func logrusLevelConverter(level log.Level) logrus.Level {
	switch level {
	case log.LevelDebug:
		return logrus.DebugLevel
	case log.LevelInfo:
		return logrus.InfoLevel
	case log.LevelWarn:
		return logrus.WarnLevel
	case log.LevelError:
		return logrus.ErrorLevel
	case log.LevelFatal:
		return logrus.FatalLevel
	default:
		return logrus.InfoLevel
	}
}
