package app

import (
	"fmt"
	"testing"

	"github.com/go-sdk/lib/log"
)

func TestNew(t *testing.T) {
	a := New("test")
	defer a.Recover()

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
