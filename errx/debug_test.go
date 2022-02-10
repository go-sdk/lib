package errx

import (
	"errors"
	"fmt"
	"testing"
)

func init() {
	SetTrack(true)
	SetNotEmptyExit(false)
}

func TestNotEmpty(t *testing.T) {
	err := errors.New("xyz")
	if NotEmpty(err) {
		t.Log("not empty")
	}
}

func TestNotEmpty2(t *testing.T) {
	defer func() { NotEmpty(fmt.Errorf("%v", recover()), 4) }()

	panic("xyz")
}
