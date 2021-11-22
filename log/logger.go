package log

import (
	"context"
	"fmt"
	"io"
	"sync"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	l *logrus.Logger

	mu     sync.Mutex
	fields Fields

	skip     int
	skipOnce sync.Once
}

func New() *Logger {
	return NewWithLevel(InfoLevel)
}

func NewWithLevel(level Level) *Logger {
	log := &Logger{
		l:      logrus.New(),
		mu:     sync.Mutex{},
		fields: make(Fields),
	}
	log.SetOutput(io.Discard)
	log.SetLevel(level)
	log.AddHook(NewConsoleHook(&ConsoleHookConfig{Level: level, ForceColors: true}))
	return log
}

func (l *Logger) Debug(v ...interface{}) {
	l.l.WithFields(l.fields).Debug(v...)
}

func (l *Logger) Info(v ...interface{}) {
	l.l.WithFields(l.fields).Info(v...)
}

func (l *Logger) Warn(v ...interface{}) {
	l.l.WithFields(l.fields).Warn(v...)
}

func (l *Logger) Error(v ...interface{}) {
	l.l.WithFields(l.fields).Error(v...)
}

func (l *Logger) Fatal(v ...interface{}) {
	l.l.WithFields(l.fields).Fatal(v...)
}

func (l *Logger) Panic(v ...interface{}) {
	l.l.WithFields(l.fields).Panic(v...)
}

func (l *Logger) Debugf(s string, v ...interface{}) {
	l.l.WithFields(l.fields).Debugf(s, v...)
}

func (l *Logger) Infof(s string, v ...interface{}) {
	l.l.WithFields(l.fields).Infof(s, v...)
}

func (l *Logger) Warnf(s string, v ...interface{}) {
	l.l.WithFields(l.fields).Warnf(s, v...)
}

func (l *Logger) Errorf(s string, v ...interface{}) {
	l.l.WithFields(l.fields).Errorf(s, v...)
}

func (l *Logger) Fatalf(s string, v ...interface{}) {
	l.l.WithFields(l.fields).Fatalf(s, v...)
}

func (l *Logger) Panicf(s string, v ...interface{}) {
	l.l.WithFields(l.fields).Panicf(s, v...)
}

func (l *Logger) WithContext(ctx context.Context) *Entry {
	return NewEntry(l).WithContext(ctx)
}

func (l *Logger) WithField(k string, v interface{}) *Entry {
	return NewEntry(l).WithField(k, v)
}

func (l *Logger) WithFields(kv Fields) *Entry {
	return NewEntry(l).WithFields(kv)
}

func (l *Logger) Caller(skip ...int) *Entry {
	return NewEntry(l).Caller(skip...)
}

func (l *Logger) AttachField(k string, v interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.fields[k] = v
}

func (l *Logger) AttachFields(kv Fields) {
	l.mu.Lock()
	defer l.mu.Unlock()
	for k, v := range kv {
		l.fields[k] = v
	}
}

func (l *Logger) SetLevel(level Level) {
	l.l.SetLevel(logrus.Level(level))
}

func (l *Logger) GetLevel() Level {
	return Level(l.l.GetLevel())
}

func (l *Logger) AddHook(hook Hook) {
	l.l.AddHook(hook)
}

func (l *Logger) AddHooks(hooks ...Hook) {
	for i := 0; i < len(hooks); i++ {
		l.AddHook(hooks[i])
	}
}

func (l *Logger) ReplaceHooks(hooks ...Hook) {
	l.l.ReplaceHooks(LevelHooks{})
	l.AddHooks(hooks...)
}

func (l *Logger) SetFormatter(formatter Formatter) {
	l.l.SetFormatter(formatter)
}

func (l *Logger) SetOutput(output io.Writer) {
	l.l.SetOutput(output)
}

func (l *Logger) SetExitFunc(f func(int)) {
	l.l.ExitFunc = f
}

func ToError(i interface{}) error {
	switch x := i.(type) {
	case *logrus.Entry:
		return fmt.Errorf(x.Message)
	case error:
		return x
	default:
		return nil
	}
}
