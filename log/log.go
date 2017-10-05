package log

import (
	"io"
	"os"

	"github.com/Sirupsen/logrus"
)

var (
	logger	*logrus.Entry
)

func init() {
	logger = logrus.StandardLogger().WithFields(logrus.Fields{})
	logrus.SetOutput(os.Stdout)
}

func Context(context interface{}) *logrus.Entry {
	return logger.WithField("context", context)
}

func SetOutput(out io.Writer) {
	logrus.SetOutput(out)
}

func SetFormatter(formatter logrus.Formatter) {
	logrus.SetFormatter(formatter)
}

func SetLevel(level string) {
	lvl, err := logrus.ParseLevel(level)
	if err != nil {
		return
	}
	logrus.SetLevel(lvl)
}

func GetLevel() logrus.Level {
	return logrus.GetLevel()
}

func AddHook(hook logrus.Hook) {
	logrus.AddHook(hook)
}

func WithError(err error) *logrus.Entry {
	return logger.WithError(err)
}

func WithField(key string, value interface{}) *logrus.Entry {
	return logger.WithField(key, value)
}

func WithFields(fields logrus.Fields) *logrus.Entry {
	return logger.WithFields(fields)
}

func Debug(args ...interface{}) {
	logger.Debug(args...)
}

func Print(args ...interface{}) {
	logger.Print(args...)
}

func Info(args ...interface{}) {
	logger.Info(args...)
}

func Warn(args ...interface{}) {
	logger.Warn(args...)
}

func Warning(args ...interface{}) {
	logger.Warning(args...)
}

func Error(args ...interface{}) {
	logger.Error(args...)
}

func Panic(args ...interface{}) {
	logger.Panic(args...)
}

func Fatal(args ...interface{}) {
	logger.Fatal(args...)
}

func Debugf(format string, args ...interface{}) {
	logger.Debugf(format, args...)
}

func Printf(format string, args ...interface{}) {
	logger.Printf(format, args...)
}

func Infof(format string, args ...interface{}) {
	logger.Infof(format, args...)
}

func Warnf(format string, args ...interface{}) {
	logger.Warnf(format, args...)
}

func Warningf(format string, args ...interface{}) {
	logger.Warningf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	logger.Errorf(format, args...)
}

func Panicf(format string, args ...interface{}) {
	logger.Panicf(format, args...)
}

func Fatalf(format string, args ...interface{}) {
	logger.Fatalf(format, args...)
}

func Debugln(args ...interface{}) {
	logger.Debugln(args...)
}

func Println(args ...interface{}) {
	logger.Println(args...)
}

func Infoln(args ...interface{}) {
	logger.Infoln(args...)
}

func Warnln(args ...interface{}) {
	logger.Warnln(args...)
}

func Warningln(args ...interface{}) {
	logger.Warningln(args...)
}

func Errorln(args ...interface{}) {
	logger.Errorln(args...)
}

func Panicln(args ...interface{}) {
	logger.Panicln(args...)
}

func Fatalln(args ...interface{}) {
	logger.Fatalln(args...)
}

func Writer() *io.PipeWriter {
	return WriterLevel(logrus.InfoLevel)
}

func WriterLevel(level logrus.Level) *io.PipeWriter {
	return logger.WriterLevel(level)
}
