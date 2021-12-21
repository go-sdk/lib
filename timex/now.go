package timex

import (
	"time"
)

var Now = time.Now

func BeginningOfMinute() time.Time {
	return New(Now()).BeginningOfMinute()
}

func BeginningOfHour() time.Time {
	return New(Now()).BeginningOfHour()
}

func BeginningOfDay() time.Time {
	return New(Now()).BeginningOfDay()
}

func BeginningOfWeek() time.Time {
	return New(Now()).BeginningOfWeek()
}

func BeginningOfWeekday(weekStartDay time.Weekday) time.Time {
	return New(Now()).BeginningOfWeekday(weekStartDay)
}

func BeginningOfMonth() time.Time {
	return New(Now()).BeginningOfMonth()
}

func BeginningOfQuarter() time.Time {
	return New(Now()).BeginningOfQuarter()
}

func BeginningOfHalf() time.Time {
	return New(Now()).BeginningOfHalf()
}

func BeginningOfYear() time.Time {
	return New(Now()).BeginningOfYear()
}

func EndOfMinute() time.Time {
	return New(Now()).EndOfMinute()
}

func EndOfHour() time.Time {
	return New(Now()).EndOfHour()
}

func EndOfDay() time.Time {
	return New(Now()).EndOfDay()
}

func EndOfWeek() time.Time {
	return New(Now()).EndOfWeek()
}

func EndOfWeekday(weekStartDay time.Weekday) time.Time {
	return New(Now()).EndOfWeekday(weekStartDay)
}

func EndOfMonth() time.Time {
	return New(Now()).EndOfMonth()
}

func EndOfQuarter() time.Time {
	return New(Now()).EndOfQuarter()
}

func EndOfHalf() time.Time {
	return New(Now()).EndOfHalf()
}

func EndOfYear() time.Time {
	return New(Now()).EndOfYear()
}

func Yesterday() time.Time {
	return New(Now()).Yesterday()
}

func Today() time.Time {
	return New(Now()).Today()
}

func Tomorrow() time.Time {
	return New(Now()).Tomorrow()
}

func Monday() time.Time {
	return New(Now()).Monday()
}

func Sunday() time.Time {
	return New(Now()).Sunday()
}

func Quarter() uint {
	return New(Now()).Quarter()
}
