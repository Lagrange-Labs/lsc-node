package logger

import (
	"fmt"
	"os"
	"strings"

	"github.com/hermeznetwork/tracerr"
	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

func init() {
	log = logrus.New()

	// Set the logger output to stdout
	log.SetOutput(os.Stdout)

	// Set the log level and format
	log.SetLevel(logrus.DebugLevel)
	log.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02T15:04:05.000Z07:00",
		FullTimestamp:   true,
	})
}

func sprintStackTrace(st []tracerr.Frame) string {
	builder := strings.Builder{}
	// Skip deepest frame because it belongs to the go runtime and we don't
	// care about it.
	if len(st) > 0 {
		st = st[:len(st)-1]
	}
	for _, f := range st {
		builder.WriteString(fmt.Sprintf("\n%s:%d %s()", f.Path, f.Line, f.Func))
	}
	builder.WriteString("\n")
	return builder.String()
}

// appendStackTraceMaybeArgs will append the stacktrace to the args
func appendStackTraceMaybeArgs(args []interface{}) []interface{} {
	for i := range args {
		if err, ok := args[i].(error); ok {
			err = tracerr.Wrap(err)
			st := tracerr.StackTrace(err)
			stackTrace := sprintStackTrace(st)
			args[i] = fmt.Sprintf("%v\nStack trace: %s", err, stackTrace)
		}
	}
	return args
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
	args = appendStackTraceMaybeArgs(args)
	log.Warn(args...)
}

func Warnf(str string, args ...interface{}) {
	args = appendStackTraceMaybeArgs(args)
	log.Warnf(str, args...)
}

func Error(args ...interface{}) {
	args = appendStackTraceMaybeArgs(args)
	log.Error(args...)
}

func Errorf(str string, args ...interface{}) {
	args = appendStackTraceMaybeArgs(args)
	log.Errorf(str, args...)
}

func Fatal(args ...interface{}) {
	args = appendStackTraceMaybeArgs(args)
	log.Fatal(args...)
}

func Fatalf(str string, args ...interface{}) {
	args = appendStackTraceMaybeArgs(args)
	log.Fatalf(str, args...)
}

func Panic(args ...interface{}) {
	args = appendStackTraceMaybeArgs(args)
	log.Panic(args...)
}

func Panicf(str string, args ...interface{}) {
	args = appendStackTraceMaybeArgs(args)
	log.Panicf(str, args...)
}
