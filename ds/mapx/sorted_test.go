package mapx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSorted(t *testing.T) {
	m := NewSorted()

	assert.Equal(t, true, m.IsEmpty())

	AssertGet := func(t *testing.T, key interface{}, value interface{}, exist bool) {
		t.Helper()

		v, x := m.Get(key)
		assert.Equal(t, exist, x)
		assert.Equal(t, value, v)
	}

	m.Put("a", 1)
	m.Put("b", true)
	m.Put("c", "n")
	m.Put(1, "n")
	m.Put(1, "a")

	assert.Equal(t, 4, m.Size())

	AssertGet(t, 1, "n", true)
	AssertGet(t, "d", nil, false)

	assert.Equal(t, true, m.ContainsKey(1))

	m.Put(1, "a", true)
	AssertGet(t, 1, "a", true)

	m.Replace(1, "n")
	AssertGet(t, 1, "n", true)

	m.Replace(2, "n")
	AssertGet(t, 2, nil, false)

	m.Replace(2, "n", true)
	AssertGet(t, 2, "n", true)

	assert.Equal(t, []interface{}{"a", "b", "c", 1, 2}, m.Keys())
	assert.Equal(t, []interface{}{1, true, "n", "n", "n"}, m.Values())

	m.Remove("b")

	assert.Equal(t, []interface{}{"a", "c", 1, 2}, m.Keys())

	m.ForEach(func(key, value interface{}) bool {
		switch key.(type) {
		case string:
			m.Remove(key)
		}
		return true
	})

	assert.Equal(t, []interface{}{1, 2}, m.Keys())

	m.Clear()

	assert.Equal(t, true, m.IsEmpty())
}
