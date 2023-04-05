package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

func init() {
	log = logrus.New()

	// Set the logger output to stdout
	log.SetOutput(os.Stdout)

	// Set the log level and format
	log.SetLevel(logrus.DebugLevel)
	log.SetFormatter(&logrus.JSONFormatter{})
}

func Debug(args ...interface{}) {
	log.Debug(args...)
}

func Debugf(str string, args ...interface{}) {
	log.Debugf(str, args...)
}

func Info(args ...interface{}) {
	log.Info(args...)
}

func Infof(str string, args ...interface{}) {
	log.Infof(str, args...)
}

func Warn(args ...interface{}) {
	log.Warn(args...)
}

func Warnf(str string, args ...interface{}) {
	log.Warnf(str, args...)
}

func Error(args ...interface{}) {
	log.Error(args...)
}

func Errorf(str string, args ...interface{}) {
	log.Errorf(str, args...)
}

func Fatal(args ...interface{}) {
	log.Fatal(args...)
}

func Fatalf(str string, args ...interface{}) {
	log.Fatalf(str, args...)
}

func Panic(args ...interface{}) {
	log.Panic(args...)
}

func Panicf(str string, args ...interface{}) {
	log.Panicf(str, args...)
}
