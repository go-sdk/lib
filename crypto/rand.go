package crypto

import (
	"crypto/rand"
	"math/big"
)

const (
	CharsetLetterLower = "abcdefghijklmnopqrstuvwxyz"
	CharsetLetterUpper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	CharsetNumber      = "0123456789"
	CharsetSymbol      = "!@#$%^&*"

	CharsetLetter          = CharsetLetterLower + CharsetLetterUpper
	CharsetLetterAndNumber = CharsetLetter + CharsetNumber
	CharsetAscii           = CharsetLetterAndNumber + CharsetSymbol
)

// RandString generate with CharsetLetterAndNumber
func RandString(length int, charsets ...string) string {
	charset := ""

	if len(charsets) != 0 {
		for i := 0; i < len(charsets); i++ {
			charset += charsets[i]
		}
	}

	if charset == "" {
		charset = CharsetLetterAndNumber
	}

	randString := make([]byte, length)
	charsetLength := big.NewInt(int64(len(charset)))

	for i := 0; i < length; i++ {
		b, err := rand.Int(rand.Reader, charsetLength)
		if err != nil {
			panic(err)
		}
		randString[i] = charset[b.Int64()]
	}

	return string(randString)
}

// RandInt returns a uniform random value in [min, max). It panics if max <= 0 and return 0 if min > max.
func RandInt(min, max int) int {
	switch {
	case min > max:
		return 0
	case min == max:
		return min
	default:
		diff := max - min
		b, err := rand.Int(rand.Reader, big.NewInt(int64(diff)))
		if err != nil {
			panic(err)
		}
		return min + int(b.Int64())
	}
}
