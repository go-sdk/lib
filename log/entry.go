package log

import (
	"context"
	"runtime"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

type Entry struct {
	l *Logger
	e *logrus.Entry
}

func NewEntry(l *Logger) *Entry {
	return &Entry{
		l: l,
		e: logrus.NewEntry(l.l),
	}
}

func (e *Entry) Debug(v ...interface{}) {
	e.e.WithFields(e.l.fields).Debug(v...)
}

func (e *Entry) Info(v ...interface{}) {
	e.e.WithFields(e.l.fields).Info(v...)
}

func (e *Entry) Warn(v ...interface{}) {
	e.e.WithFields(e.l.fields).Warn(v...)
}

func (e *Entry) Error(v ...interface{}) {
	e.e.WithFields(e.l.fields).Error(v...)
}

func (e *Entry) Fatal(v ...interface{}) {
	e.e.WithFields(e.l.fields).Fatal(v...)
}

func (e *Entry) Panic(v ...interface{}) {
	e.e.WithFields(e.l.fields).Panic(v...)
}

func (e *Entry) Debugf(s string, v ...interface{}) {
	e.e.WithFields(e.l.fields).Debugf(s, v...)
}

func (e *Entry) Infof(s string, v ...interface{}) {
	e.e.WithFields(e.l.fields).Infof(s, v...)
}

func (e *Entry) Warnf(s string, v ...interface{}) {
	e.e.WithFields(e.l.fields).Warnf(s, v...)
}

func (e *Entry) Errorf(s string, v ...interface{}) {
	e.e.WithFields(e.l.fields).Errorf(s, v...)
}

func (e *Entry) Fatalf(s string, v ...interface{}) {
	e.e.WithFields(e.l.fields).Fatalf(s, v...)
}

func (e *Entry) Panicf(s string, v ...interface{}) {
	e.e.WithFields(e.l.fields).Panicf(s, v...)
}

func (e *Entry) WithContext(ctx context.Context) *Entry {
	return &Entry{l: e.l, e: logrus.NewEntry(e.l.l).WithContext(ctx)}
}

func (e *Entry) WithField(k string, v interface{}) *Entry {
	return &Entry{l: e.l, e: logrus.NewEntry(e.l.l).WithField(k, v)}
}

func (e *Entry) WithFields(kv Fields) *Entry {
	return &Entry{l: e.l, e: logrus.NewEntry(e.l.l).WithFields(kv)}
}

func (e *Entry) Caller(skip ...int) *Entry {
	caller := e.getCaller(skip...)
	if caller != "" {
		return e.WithField("caller", caller)
	}
	return e
}

const (
	maxSkip = 20
	pkgName = "starudream/lib/log"
)

func (e *Entry) getCaller(skip ...int) string {
	e.l.skipOnce.Do(func() {
		ls, le := 0, 0
		for i := 1; i < maxSkip; i++ {
			_, f, _, ok := runtime.Caller(i)
			if !ok {
				break
			}
			if strings.Contains(f, pkgName) && !strings.Contains(f, "_test.go") {
				if ls == 0 {
					ls = i
				}
				le = i
			}
		}
		e.l.skip = le + 1 - ls
	})
	if len(skip) == 0 || skip[0] < 0 {
		skip = []int{0}
	}
	pc, f, l, ok := runtime.Caller(e.l.skip + skip[0])
	if !ok {
		return ""
	}
	fn := runtime.FuncForPC(pc).Name()
	return f + ":" + strconv.Itoa(l) + " " + fn
}
