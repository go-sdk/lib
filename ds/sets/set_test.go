package sets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSet(t *testing.T) {
	s := NewSet()

	assert.Equal(t, true, s.IsEmpty())

	s.Add(1)
	s.Add(2)
	s.Adds("2", "5", "8", true)

	assert.Equal(t, 6, s.Size())

	assert.Equal(t, true, s.Contains("2", "3"))
	assert.Equal(t, false, s.Contains("3", false))

	assert.Equal(t, 1, s.Removes(3, "2", false))

	assert.Equal(t, []interface{}{1, 2, "5", "8", true}, s.Values())

	s.ForEach(func(value interface{}) bool {
		switch value.(type) {
		case string:
			s.Remove(value)

		}
		return true
	})

	s.Clear()

	assert.Equal(t, true, s.IsEmpty())
}
