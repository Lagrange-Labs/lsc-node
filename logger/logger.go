package logger

import (
	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func init() {
	Log = logrus.New()

	// Configure log rotation
	Log.SetOutput(&lumberjack.Logger{
		Filename:   "/var/log/lagrange-node/lagrange-node.log",
		MaxSize:    1,  // Megabytes
		MaxBackups: 10, // max number of log rotated log files to keep
		MaxAge:     7,  // Days
		Compress:   true,
	})

	// Set the log level and format
	Log.SetLevel(logrus.DebugLevel)
	Log.SetFormatter(&logrus.JSONFormatter{})
}
