package flag

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlice(t *testing.T) {
	IntSlice("int-slice", []int{1}, "")
	Int64Slice("int64-slice", []int64{1}, "")
	StringSlice("string-slice", []string{"1"}, "")
	Float64Slice("float64-slice", []float64{1}, "")

	x1 := []int{1}
	IntSliceVar(&x1, "int-slice-var", []int{2}, "")
	x2 := []int64{1}
	Int64SliceVar(&x2, "int64-slice-var", []int64{2}, "")
	x3 := []string{"1"}
	StringSliceVar(&x3, "string-slice-var", []string{"2"}, "")
	x4 := []float64{1}
	Float64SliceVar(&x4, "float64-slice-var", []float64{2}, "")

	VisitAll(func(flag *Flag) { t.Logf("%s, val: %s, def: %s", flag.Name, flag.Value.String(), flag.DefValue) })

	assert.NoError(t, Set("int-slice", "2"))
	assert.NoError(t, Set("int64-slice", "2"))
	assert.NoError(t, Set("string-slice", "2"))
	assert.NoError(t, Set("float64-slice", "2"))

	assert.NoError(t, Set("int-slice-var", "3"))
	assert.NoError(t, Set("int64-slice-var", "3"))
	assert.NoError(t, Set("string-slice-var", "3"))
	assert.NoError(t, Set("float64-slice-var", "3"))

	assert.NoError(t, Set("int-slice-var", "4"))
	assert.NoError(t, Set("int64-slice-var", "4"))
	assert.NoError(t, Set("string-slice-var", "4"))
	assert.NoError(t, Set("float64-slice-var", "4"))

	VisitAll(func(flag *Flag) { t.Logf("%s, val: %s, def: %s", flag.Name, flag.Value.String(), flag.DefValue) })

	t.Logf("%#v %#v %#v %#v", x1, x2, x3, x4)
}
