package log

import (
	"testing"
)

func TestInit(t *testing.T) {
	AttachFields(Fields{"hashcode": "123456", "version": "v0.0.1"})
	Debug("debug")
	Info("info")
	WithFields(Fields{"method": "test"}).Warn("warn")
	WithFields(Fields{"span": "test"}).Error("error")
}
