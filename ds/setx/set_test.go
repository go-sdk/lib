package setx

import (
	"testing"

	"github.com/go-sdk/lib/testx"
)

func TestNewSet(t *testing.T) {
	s := NewSet()

	testx.AssertEqual(t, true, s.IsEmpty())

	s.Add(1)
	s.Add(2)
	s.Adds("2", "5", "8", true)

	testx.AssertEqual(t, 6, s.Size())

	testx.AssertEqual(t, true, s.Contains("2", "3"))
	testx.AssertEqual(t, false, s.Contains("3", false))

	testx.AssertEqual(t, 1, s.Removes(3, "2", false))

	testx.AssertEqual(t, []interface{}{1, 2, "5", "8", true}, s.Values())

	s.ForEach(func(value interface{}) bool {
		switch value.(type) {
		case string:
			s.Remove(value)

		}
		return true
	})

	s.Clear()

	testx.AssertEqual(t, true, s.IsEmpty())
}
