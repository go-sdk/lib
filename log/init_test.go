package log

import (
	"context"
	"testing"
)

func TestInit(t *testing.T) {
	AttachField("app", "test")

	AttachFields(Fields{"hashcode": "123456", "version": "v0.0.1"})

	SetLevel(InfoLevel)

	DefaultLogger().SetExitFunc(func(i int) { t.Log(i) })

	Debug("Debug")
	Info("Info")
	Warn("Warn")
	Error("Error")
	Fatal("Fatal")
	// Panic("Panic")

	Debugf("Debugf")
	Infof("Infof")
	Warnf("Warnf")
	Errorf("Errorf")
	Fatalf("Fatalf")
	// Panicf("Panicf")

	WithContext(context.Background()).Info("Info")

	WithField("span", "test").Info("Info")

	WithFields(Fields{"span": "test"}).Info("Info")

	Caller().Info("Info")

	SetLevel(ErrorLevel)

	Info("Info")

	t.Log(GetLevel())
}

func TestPanic(t *testing.T) {
	defer func() {
		if err := ToError(recover()); err != nil {
			t.Log(err)
		}
	}()

	Panic("Panic")
}

func TestPanicf(t *testing.T) {
	defer func() {
		if err := ToError(recover()); err != nil {
			t.Log(err)
		}
	}()

	Panicf("Panicf")
}
