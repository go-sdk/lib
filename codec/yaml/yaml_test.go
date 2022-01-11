package yaml

import (
	"fmt"
	"testing"

	"github.com/go-sdk/lib/testx"
)

var (
	Data = map[string]interface{}{
		"foo": map[string]interface{}{"foo": "bar"},
		"bar": map[string]interface{}{"bar": "foo"},
		"baz": map[string]interface{}{"bar": "foo", "baz": "abc"},
	}

	DataRaw = `foo:
  foo: bar
bar:
  &bar
  bar: foo
baz:
  <<: *bar
  baz: abc
`

	DataTiled = `bar:
  bar: foo
baz:
  bar: foo
  baz: abc
foo:
  foo: bar
`
)

func TestMustMarshal(t *testing.T) {
	testx.AssertEqual(t, DataTiled, string(MustMarshal(Data)))
}

func TestMustMarshalToString(t *testing.T) {
	testx.AssertEqual(t, DataTiled, MustMarshalToString(Data))
}

func TestUnmarshalFromString(t *testing.T) {
	t.Run("default", func(t *testing.T) {
		var v interface{}
		testx.AssertNoError(t, UnmarshalFromString(DataRaw, &v))
		testx.AssertEqual(t, DataTiled, MustMarshalToString(v))
	})

	t.Run("with cleanup", func(t *testing.T) {
		var v interface{}
		testx.AssertNoError(t, UnmarshalFromString(DataRaw, &v, WithCleanup(true)))
		testx.AssertNotContains(t, fmt.Sprintf("%#v", v), "map[interface {}]interface {}")
		testx.AssertEqual(t, DataTiled, MustMarshalToString(v))
	})
}

func TestPrint(t *testing.T) {
	Print(Data)
}
