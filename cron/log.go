package cron

import (
	"fmt"

	"github.com/go-sdk/lib/log"
)

type logger struct {
	l *log.Logger
}

func newLogger() *logger {
	fmt.Println(log.DefaultLogger().GetLevel())
	return &logger{l: log.DefaultLogger()}
}

func (l *logger) Debug(msg string, keysAndValues ...interface{}) {
	l.l.WithFields(log.ToFields(append(keysAndValues, "span", "cron")...)).Debug(msg)
}

func (l *logger) Info(msg string, keysAndValues ...interface{}) {
	l.l.WithFields(log.ToFields(append(keysAndValues, "span", "cron")...)).Info(msg)
}

func (l *logger) Error(err error, msg string, keysAndValues ...interface{}) {
	l.l.WithFields(log.ToFields(append(keysAndValues, "err", err, "span", "cron")...)).Error(msg)
}
