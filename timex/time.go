package timex

import (
	"time"
)

const WeekStartDay = time.Monday

type Time struct {
	time.Time
}

func New(t time.Time) *Time {
	return &Time{Time: t}
}

func (now *Time) BeginningOfMinute() time.Time {
	return now.Truncate(time.Minute)
}

func (now *Time) BeginningOfHour() time.Time {
	y, m, d := now.Date()
	return time.Date(y, m, d, now.Time.Hour(), 0, 0, 0, now.Time.Location())
}

func (now *Time) BeginningOfDay() time.Time {
	y, m, d := now.Date()
	return time.Date(y, m, d, 0, 0, 0, 0, now.Time.Location())
}

func (now *Time) BeginningOfWeek() time.Time {
	return now.BeginningOfWeekday(WeekStartDay)
}

func (now *Time) BeginningOfWeekday(weekStartDay time.Weekday) time.Time {
	b := now.BeginningOfDay()
	weekday := int(b.Weekday())
	if weekStartDay != time.Sunday {
		weekStartDayInt := int(weekStartDay)
		if weekday < weekStartDayInt {
			weekday = weekday + 7 - weekStartDayInt
		} else {
			weekday = weekday - weekStartDayInt
		}
	}
	return b.AddDate(0, 0, -weekday)
}

func (now *Time) BeginningOfMonth() time.Time {
	y, m, _ := now.Date()
	return time.Date(y, m, 1, 0, 0, 0, 0, now.Location())
}

func (now *Time) BeginningOfQuarter() time.Time {
	month := now.BeginningOfMonth()
	offset := (int(month.Month()) - 1) % 3
	return month.AddDate(0, -offset, 0)
}

func (now *Time) BeginningOfHalf() time.Time {
	month := now.BeginningOfMonth()
	offset := (int(month.Month()) - 1) % 6
	return month.AddDate(0, -offset, 0)
}

func (now *Time) BeginningOfYear() time.Time {
	y, _, _ := now.Date()
	return time.Date(y, time.January, 1, 0, 0, 0, 0, now.Location())
}

func (now *Time) EndOfMinute() time.Time {
	return now.BeginningOfMinute().Add(time.Minute - time.Nanosecond)
}

func (now *Time) EndOfHour() time.Time {
	return now.BeginningOfHour().Add(time.Hour - time.Nanosecond)
}

func (now *Time) EndOfDay() time.Time {
	y, m, d := now.Date()
	return time.Date(y, m, d, 23, 59, 59, int(time.Second-time.Nanosecond), now.Location())
}

func (now *Time) EndOfWeek() time.Time {
	return now.BeginningOfWeek().AddDate(0, 0, 7).Add(-time.Nanosecond)
}

func (now *Time) EndOfWeekday(weekStartDay time.Weekday) time.Time {
	return now.BeginningOfWeekday(weekStartDay).AddDate(0, 0, 7).Add(-time.Nanosecond)
}

func (now *Time) EndOfMonth() time.Time {
	return now.BeginningOfMonth().AddDate(0, 1, 0).Add(-time.Nanosecond)
}

func (now *Time) EndOfQuarter() time.Time {
	return now.BeginningOfQuarter().AddDate(0, 3, 0).Add(-time.Nanosecond)
}

func (now *Time) EndOfHalf() time.Time {
	return now.BeginningOfHalf().AddDate(0, 6, 0).Add(-time.Nanosecond)
}

func (now *Time) EndOfYear() time.Time {
	return now.BeginningOfYear().AddDate(1, 0, 0).Add(-time.Nanosecond)
}

func (now *Time) Yesterday() time.Time {
	return now.BeginningOfDay().AddDate(0, 0, -1)
}

func (now *Time) Today() time.Time {
	return now.BeginningOfDay()
}

func (now *Time) Tomorrow() time.Time {
	return now.BeginningOfDay().AddDate(0, 0, 1)
}

func (now *Time) Monday() time.Time {
	b := now.BeginningOfDay()
	weekday := int(b.Weekday())
	if weekday == 0 {
		weekday = 7
	}
	return b.AddDate(0, 0, -weekday+1)
}

func (now *Time) Sunday() time.Time {
	b := now.BeginningOfDay()
	weekday := int(b.Weekday())
	if weekday == 0 {
		return b
	}
	return b.AddDate(0, 0, 7-weekday)
}

func (now *Time) Quarter() uint {
	return (uint(now.Month())-1)/3 + 1
}
