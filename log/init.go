package log

import (
	"context"

	"github.com/go-sdk/lib/conf"
)

func init() {
	var hooks []Hook

	cl, err := ParseLevel(conf.Get("log.console.level").String())
	if err == nil {
		h := NewConsoleHook(&ConsoleHookConfig{
			Level:       cl,
			ForceColors: conf.Get("log.console.color").BoolD(true),
		})
		hooks = append(hooks, h)
	}

	fl, err := ParseLevel(conf.Get("log.file.level").String())
	if err == nil {
		h := NewFileHook(&FileHookConfig{
			Level:       fl,
			ForceColors: conf.Get("log.file.color").BoolD(false),
			ForceJSON:   conf.Get("log.file.json").BoolD(false),
			Filename:    conf.Get("log.file.path").String(),
			MaxSize:     int(conf.Get("log.file.max_size").Int64()),
			MaxAge:      int(conf.Get("log.file.max_age").Int64()),
			MaxBackups:  int(conf.Get("log.file.max_backups").Int64()),
			LocalTime:   conf.Get("log.file.local_time").Bool(),
			Compress:    conf.Get("log.file.compress").Bool(),
		})
		hooks = append(hooks, h)
	}

	if len(hooks) > 0 {
		logger = New()
		logger.ReplaceHooks(hooks...)
	}
}

var logger = NewWithLevel(DebugLevel)

func DefaultLogger() *Logger {
	return logger
}

func Debug(v ...interface{}) {
	logger.Debug(v...)
}

func Info(v ...interface{}) {
	logger.Info(v...)
}

func Warn(v ...interface{}) {
	logger.Warn(v...)
}

func Error(v ...interface{}) {
	logger.Error(v...)
}

func Fatal(v ...interface{}) {
	logger.Fatal(v...)
}

func Panic(v ...interface{}) {
	logger.Panic(v...)
}

func Debugf(s string, v ...interface{}) {
	logger.Debugf(s, v...)
}

func Infof(s string, v ...interface{}) {
	logger.Infof(s, v...)
}

func Warnf(s string, v ...interface{}) {
	logger.Warnf(s, v...)
}

func Errorf(s string, v ...interface{}) {
	logger.Errorf(s, v...)
}

func Fatalf(s string, v ...interface{}) {
	logger.Fatalf(s, v...)
}

func Panicf(s string, v ...interface{}) {
	logger.Panicf(s, v...)
}

func WithContext(ctx context.Context) *Entry {
	return logger.WithContext(ctx)
}

func WithField(k string, v interface{}) *Entry {
	return logger.WithField(k, v)
}

func WithFields(kv Fields) *Entry {
	return logger.WithFields(kv)
}

func Caller(skip ...int) *Entry {
	return logger.Caller(skip...)
}

func AttachField(k string, v interface{}) {
	logger.AttachField(k, v)
}

func AttachFields(kv Fields) {
	logger.AttachFields(kv)
}

func SetLevel(level Level) {
	logger.SetLevel(level)
}

func GetLevel() Level {
	return logger.GetLevel()
}

func AddHook(hook Hook) {
	logger.AddHook(hook)
}

func AddHooks(hooks ...Hook) {
	logger.AddHooks(hooks...)
}
