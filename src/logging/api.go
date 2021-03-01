package logging

import "github.com/sirupsen/logrus"

type Logger interface {
	Info(txID string) *logrus.Entry
	Debug(txID string, message string) *logrus.Entry
	Error(txID string, errorCode string) *logrus.Entry
	Fatal(txID string, errorCode string) *logrus.Entry
	WithField(key string, value interface{}) *logrus.Entry
	WithFields(fields logrus.Fields) *logrus.Entry
	Logger() *logrus.Logger
}
