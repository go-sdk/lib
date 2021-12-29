package db

import (
	"context"
	"errors"
	"time"

	"gorm.io/gorm/logger"

	"github.com/go-sdk/lib/log"
)

const (
	gormLogSilent = logger.Silent
	gormLogError  = logger.Error
	gormLogWarn   = logger.Warn
	gormLogInfo   = logger.Info
)

type gormLogger struct {
	e *log.Entry

	UseInfoSQL        bool
	ShowNotFoundError bool
	SlowThreshold     time.Duration
	LogLevel          logger.LogLevel
}

func (l *gormLogger) LogMode(level logger.LogLevel) logger.Interface {
	n := *l
	n.LogLevel = level
	return &n
}

func (l *gormLogger) Info(ctx context.Context, s string, i ...interface{}) {
	if l.LogLevel >= gormLogInfo {
		l.e.WithContext(ctx).Infof(s, i...)
	}
}

func (l *gormLogger) Warn(ctx context.Context, s string, i ...interface{}) {
	if l.LogLevel >= gormLogWarn {
		l.e.WithContext(ctx).Warnf(s, i...)
	}
}

func (l *gormLogger) Error(ctx context.Context, s string, i ...interface{}) {
	if l.LogLevel >= gormLogError {
		l.e.WithContext(ctx).Errorf(s, i...)
	}
}

func (l *gormLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	if l.LogLevel <= gormLogSilent {
		return
	}

	elapsed := time.Since(begin)

	x := l.e.WithContext(ctx).WithField("took", elapsed.String())

	sql, rows := fc()
	if rows != -1 {
		x = x.WithField("rows", rows)
	}

	switch {
	case err != nil && l.LogLevel >= gormLogError && (!errors.Is(err, logger.ErrRecordNotFound) || l.ShowNotFoundError):
		x.WithField("err", err).Error(sql)
	case l.LogLevel >= gormLogWarn && l.SlowThreshold != 0 && elapsed > l.SlowThreshold:
		x.Warnf("slow sql: %s", sql)
	case l.LogLevel == gormLogInfo:
		if l.UseInfoSQL {
			x.Info(sql)
		} else {
			x.Debug(sql)
		}
	}
}
