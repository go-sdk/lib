package log

import (
	"testing"
)

func TestInit(t *testing.T) {
	AttachFields(Fields{"hashcode": "123456", "version": "v0.0.1"})
	Debug("Debug")
	Info("Info")
	WithFields(Fields{"method": "test"}).Warn("Warn")
	Caller().Warnf("Warnf")
	WithFields(Fields{"span": "test"}).Error("Error")
	Caller().Errorf("Errorf")
}
