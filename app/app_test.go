package app

import (
	"fmt"
	"testing"
	"time"

	"github.com/go-sdk/lib/log"
	"github.com/go-sdk/lib/testx"
)

func TestNew(t *testing.T) {
	a := New("test")
	defer a.Recover()

	a.Add()

	a.Add(
		func() error {
			log.Info("1")
			return nil
		},
		func() error {
			return fmt.Errorf("2")
		},
		func() error {
			log.Info("3")
			return nil
		},
	)

	_ = a.Run()
}

func TestNewAddAfterRun(t *testing.T) {
	a := New("test")
	defer a.Recover()

	a.Add(
		func() error {
			log.Info("1")
			return nil
		},
	)

	a.Start()

	time.Sleep(50 * time.Millisecond)

	a.Add(
		func() error {
			log.Info("2")
			return nil
		},
	)

	testx.AssertEqual(t, 1, len(a.ss))

	a.Stop()

	time.Sleep(50 * time.Millisecond)
}

func TestRecover(t *testing.T) {
	a := New("test")
	defer a.Recover()

	panic("panic")
}

func TestNewOnce(t *testing.T) {
	a := New("test")
	defer a.Recover()

	a.Add(
		func() error {
			log.Info("1")
			return nil
		},
		func() error {
			log.Info("3")
			return nil
		},
	)

	_ = a.Once()
}
