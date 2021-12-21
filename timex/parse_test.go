package timex

import (
	"testing"
	"time"
)

func TestParse(t *testing.T) {
	now := time.Now()
	t.Log(MustParse(now.Format(DateTime)))
	t.Log(MustParseIn(now.Format(DateTime), time.UTC))
}
