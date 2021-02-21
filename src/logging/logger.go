// Package logging provides structured logging with logrus.
package logging

import (
	"log"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Log implements the logger
type Log struct {
	Log *logrus.Logger
}

// NewLogger creates and configures a new logrus Logger.
func NewLogger() *Log {
	// create new logger
	var l Log
	l.Log = logrus.New()
	l.Log.SetReportCaller(true)

	filename := viper.GetString("logging.log_filename")
	// Create the log file if doesn't exist. And append to it if it already exists.
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalf("Failed to create log file with error %v", err)
	}
	l.Log.SetOutput(f)

	if viper.GetBool("logging.textlogging") {
		l.Log.Formatter = &logrus.TextFormatter{
			DisableTimestamp: false,
		}
	} else {
		l.Log.Formatter = &logrus.JSONFormatter{
			DisableTimestamp: false,
		}
	}

	level := viper.GetString("logging.log_level")
	if level == "" {
		level = "error"
	}
	logLevel, err := logrus.ParseLevel(level)
	if err != nil {
		log.Fatal(err)
	}
	l.Log.Level = logLevel

	return &l
}

// Info adds info logs
func (l *Log) Info(txID string) *logrus.Entry {
	return l.Log.WithField("txID", txID)
}

// Debug adds debug logs
func (l *Log) Debug(txID string, message string) *logrus.Entry {
	return l.Log.WithField("txID", txID)
}

// Error adds error logs
func (l *Log) Error(txID string, errorCode string) *logrus.Entry {
	return l.Log.WithField("txID", txID).WithField("errorCode", errorCode)
}
