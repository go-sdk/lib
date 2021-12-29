package cron

import (
	"github.com/go-sdk/lib/log"
)

type logger struct {
	e *log.Entry
}

func newLogger() *logger {
	return &logger{e: log.DefaultLogger().WithField("span", "cron")}
}

func (l *logger) Debug(msg string, keysAndValues ...interface{}) {
	l.e.WithFields(log.ToFields(keysAndValues...)).Debug(msg)
}

func (l *logger) Info(msg string, keysAndValues ...interface{}) {
	l.e.WithFields(log.ToFields(keysAndValues...)).Info(msg)
}

func (l *logger) Error(err error, msg string, keysAndValues ...interface{}) {
	l.e.WithField("err", err).WithFields(log.ToFields(keysAndValues...)).Error(msg)
}
