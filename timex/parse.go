package timex

import (
	"time"

	"github.com/go-sdk/lib/internal/dateparse"
)

func Parse(str string) (time.Time, error) {
	return ParseIn(str, time.Local)
}

func MustParse(str string) time.Time {
	t, err := Parse(str)
	if err != nil {
		panic(err)
	}
	return t
}

func ParseIn(str string, loc *time.Location) (time.Time, error) {
	return dateparse.ParseIn(str, loc)
}

func MustParseIn(str string, loc *time.Location) time.Time {
	t, err := ParseIn(str, loc)
	if err != nil {
		panic(err)
	}
	return t
}
