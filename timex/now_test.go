package timex

import (
	"testing"
	"time"
)

func TestNow(t *testing.T) {
	t.Log(BeginningOfMinute())
	t.Log(BeginningOfHour())
	t.Log(BeginningOfDay())
	t.Log(BeginningOfWeek())
	t.Log(BeginningOfWeekday(time.Sunday))
	t.Log(BeginningOfMonth())
	t.Log(BeginningOfQuarter())
	t.Log(BeginningOfHalf())
	t.Log(BeginningOfYear())

	t.Log("----------------------------------------")

	t.Log(EndOfMinute())
	t.Log(EndOfHour())
	t.Log(EndOfDay())
	t.Log(EndOfWeek())
	t.Log(EndOfWeekday(time.Sunday))
	t.Log(EndOfMonth())
	t.Log(EndOfQuarter())
	t.Log(EndOfHalf())
	t.Log(EndOfYear())

	t.Log("----------------------------------------")

	t.Log(Yesterday())
	t.Log(Today())
	t.Log(Tomorrow())
	t.Log(Monday())
	t.Log(Sunday())
	t.Log(Quarter())
}
