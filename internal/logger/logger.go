package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func InitLogger() {
	Log = logrus.New()

	Log.SetOutput(os.Stdout)

	level := os.Getenv("LOG_LEVEL")
	if level == "" {
		level = "info"
	}

	parsedLevel, err := logrus.ParseLevel(level)
	if err != nil {
		Log.Fatalf("Invalid log level: %v", err)
	}
	Log.SetLevel(parsedLevel)

	format := os.Getenv("LOG_FORMAT")
	if format == "json" {
		Log.SetFormatter(&logrus.JSONFormatter{})
	} else {
		Log.SetFormatter(&logrus.TextFormatter{})
	}
}
