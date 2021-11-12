package crypto

import (
	"testing"
)

func TestRandString(t *testing.T) {
	t.Log(RandString(64))
	t.Log(RandString(64, CharsetLetterLower))
	t.Log(RandString(64, CharsetLetterUpper))
	t.Log(RandString(64, CharsetNumber))
	t.Log(RandString(64, CharsetSymbol))
	t.Log(RandString(64, CharsetAscii))
}

func TestRandInt(t *testing.T) {
	t.Log(RandInt(10, 20))
	t.Log(RandInt(-10, 10))
	t.Log(RandInt(-10, 0))
}
