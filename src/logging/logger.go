// Package logging provides structured logging with logrus.
package logging

import (
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Log implements the logger
type log struct {
	log *logrus.Logger
}

// NewLogger creates and configures a new logrus Logger.
func NewLogger() Logger {
	// create new logger
	var l log
	l.log = logrus.New()
	l.log.SetReportCaller(true)

	filename := viper.GetString("logging.log_filename")
	// Create the log file if doesn't exist. And append to it if it already exists.
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		l.log.Fatalf("Failed to create log file with error %v", err)
	}
	l.log.SetOutput(f)

	if viper.GetBool("logging.textlogging") {
		l.log.Formatter = &logrus.TextFormatter{
			DisableTimestamp: true,
		}
	} else {
		l.log.Formatter = &logrus.JSONFormatter{
			DisableTimestamp: true,
		}
	}

	level := viper.GetString("logging.log_level")
	if level == "" {
		level = "error"
	}
	logLevel, err := logrus.ParseLevel(level)
	if err != nil {
		l.log.Fatal(err)
	}
	l.log.Level = logLevel

	return &l
}

// Info adds info logs
func (l *log) Info(txID string) *logrus.Entry {
	return l.log.WithField("transaction_id", txID).WithField("timestamp_utc", time.Now().UTC())
}

// Debug adds debug logs
func (l *log) Debug(txID string, message string) *logrus.Entry {
	return l.log.WithField("transaction_id", txID).WithField("timestamp_utc", time.Now().UTC())
}

// Error adds error logs
func (l *log) Error(txID string, errorCode string) *logrus.Entry {
	return l.log.WithField("transaction_id", txID).WithField("errorCode", errorCode).WithField("timestamp_utc", time.Now().UTC())
}

// Fatal adds Fatal logs
func (l *log) Fatal(txID string, errorCode string) *logrus.Entry {
	return l.log.WithField("transaction_id", txID).WithField("errorCode", errorCode).WithField("timestamp_utc", time.Now().UTC())
}

// WithField adds Field
func (l *log) WithField(key string, value interface{}) *logrus.Entry {
	return l.log.WithField("timestamp_utc", time.Now().UTC()).WithField(key, value)
}

// WithFields adds Fields
func (l *log) WithFields(fields logrus.Fields) *logrus.Entry {
	return l.log.WithField("timestamp_utc", time.Now().UTC()).WithFields(fields)
}

func (l *log) Logger() *logrus.Logger {
	return l.log
}
