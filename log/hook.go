package log

import (
	"io"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

const timeFormatter = "2006-01-02T15:04:05.999Z07:00"

type (
	Hook       = logrus.Hook
	LevelHooks = logrus.LevelHooks
	Formatter  = logrus.Formatter

	JSONFormatter = logrus.JSONFormatter
	TextFormatter = logrus.TextFormatter

	Fields = logrus.Fields
)

type hook struct {
	w io.Writer
	f Formatter

	ls []logrus.Level
}

func (h *hook) Levels() []logrus.Level {
	return h.ls
}

func (h *hook) Fire(entry *logrus.Entry) error {
	bs, err := h.f.Format(entry)
	if err != nil {
		return err
	}
	_, err = h.w.Write(bs)
	return err
}

func (h *hook) SetLevel(level Level) {
	h.ls = []logrus.Level{}
	for i := 0; i < len(AllLevels); i++ {
		l := AllLevels[i]
		if level >= l {
			h.ls = append(h.ls, logrus.Level(l))
		}
	}
}

type ConsoleHookConfig struct {
	Level       Level
	ForceColors bool
}

func NewConsoleHook(configs ...*ConsoleHookConfig) *hook {
	config := &ConsoleHookConfig{
		Level:       InfoLevel,
		ForceColors: true,
	}
	if len(configs) > 0 && configs[0] != nil {
		config = configs[0]
	}

	h := &hook{}
	h.w = os.Stdout
	h.f = &TextFormatter{
		ForceColors:     config.ForceColors,
		DisableQuote:    true,
		FullTimestamp:   true,
		TimestampFormat: timeFormatter,
	}
	h.SetLevel(config.Level)
	return h
}

type FileHookConfig struct {
	Level       Level
	ForceColors bool
	ForceJSON   bool
	Filename    string
	MaxSize     int
	MaxAge      int
	MaxBackups  int
	LocalTime   bool
	Compress    bool
}

func NewFileHook(configs ...*FileHookConfig) *hook {
	path, err := os.Executable()
	if err != nil {
		path = os.TempDir()
	}
	path = filepath.Dir(path)
	name := filepath.Base(os.Args[0]) + ".log"

	config := &FileHookConfig{
		Level:     InfoLevel,
		Filename:  filepath.Join(path, name),
		LocalTime: true,
	}
	if len(configs) > 0 && configs[0] != nil {
		config = configs[0]
		if config.Filename == "" {
			config.Filename = filepath.Join(path, name)
		}
	}

	h := &hook{}
	h.w = &lumberjack.Logger{
		Filename:   config.Filename,
		MaxSize:    config.MaxSize,
		MaxAge:     config.MaxAge,
		MaxBackups: config.MaxBackups,
		LocalTime:  config.LocalTime,
		Compress:   config.Compress,
	}
	if config.ForceJSON {
		h.f = &JSONFormatter{
			TimestampFormat:   timeFormatter,
			DisableHTMLEscape: true,
		}
	} else {
		h.f = &TextFormatter{
			ForceColors:     config.ForceColors,
			DisableQuote:    true,
			FullTimestamp:   true,
			TimestampFormat: timeFormatter,
		}
	}
	h.SetLevel(config.Level)
	return h
}
