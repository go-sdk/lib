package slicex

import (
	"testing"
)

func TestToInterface(t *testing.T) {
	var v interface{}
	v = []string{"1", "2", "3"}
	t.Logf("%#v", ToInterface(v))
}

func TestToString(t *testing.T) {
	var v interface{}
	v = []int{1, 2, 3}
	t.Logf("%#v", ToInterface(v))
	t.Logf("%#v", ToString(ToInterface(v)))
}

func TestToInt64(t *testing.T) {
	var v interface{}
	v = []int{1, 2, 3}
	t.Logf("%#v", ToInterface(v))
	t.Logf("%#v", ToInt64(ToInterface(v)))
}
