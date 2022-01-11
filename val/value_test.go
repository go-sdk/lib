package val

import (
	"fmt"
	"testing"
	"time"

	"github.com/go-sdk/lib/testx"
)

type V struct {
	v  interface{} // value
	ev interface{} // expected value
	we bool        // with error
}

func TestVal_Bool(t *testing.T) {
	list := []V{
		{v: true, ev: true},
		{v: false, ev: false},
		{v: uint(10), ev: true},
		{v: "T", ev: true},
		{v: "TT", ev: false, we: true},
	}

	for i := 0; i < len(list); i++ {
		v := list[i]
		y, e := New(v.v).BoolE()
		t.Logf("%v %v", y, e)
		testx.AssertEqual(t, v.we, e != nil)
		testx.AssertEqual(t, v.ev, y)
	}
}

func TestVal_Int64(t *testing.T) {
	list := []V{
		{v: true, ev: int64(1)},
		{v: uint(10), ev: int64(10)},
		{v: 12.3456, ev: int64(12)},
		{v: "3.1415", ev: int64(0), we: true},
		{v: "12", ev: int64(12)},
	}

	for i := 0; i < len(list); i++ {
		v := list[i]
		y, e := New(v.v).Int64E()
		t.Logf("%v %v", y, e)
		testx.AssertEqual(t, v.we, e != nil)
		testx.AssertEqual(t, v.ev, y)
	}
}

func TestVal_Float64(t *testing.T) {
	list := []V{
		{v: true, ev: float64(1)},
		{v: uint(10), ev: float64(10)},
		{v: 12.3456, ev: 12.3456},
		{v: "3.1415", ev: 3.1415},
		{v: "12", ev: float64(12)},
	}

	for i := 0; i < len(list); i++ {
		v := list[i]
		y, e := New(v.v).Float64E()
		t.Logf("%v %v", y, e)
		testx.AssertEqual(t, v.we, e != nil)
		testx.AssertEqual(t, v.ev, y)
	}
}

func TestVal_String(t *testing.T) {
	n := time.Now()

	list := []V{
		{v: true, ev: "true"},
		{v: false, ev: "false"},
		{v: uint(10), ev: "10"},
		{v: 12.3456, ev: "12.3456"},
		{v: "T", ev: "T"},
		{v: []byte("Hello"), ev: "Hello"},
		{v: n, ev: n.String()},
		{v: fmt.Errorf("error"), ev: "error"},
	}

	for i := 0; i < len(list); i++ {
		v := list[i]
		y, e := New(v.v).StringE()
		t.Logf("%v %v", y, e)
		testx.AssertEqual(t, v.we, e != nil)
		testx.AssertEqual(t, v.ev, y)
	}
}

func TestVal_Duration(t *testing.T) {
	list := []V{
		{v: true, ev: time.Duration(0), we: true},
		{v: uint(10), ev: time.Duration(10)},
		{v: 123.456, ev: time.Duration(123)},
		{v: "30s", ev: 30 * time.Second},
	}

	for i := 0; i < len(list); i++ {
		v := list[i]
		y, e := New(v.v).DurationE()
		t.Logf("%v %v", y, e)
		testx.AssertEqual(t, v.we, e != nil)
		testx.AssertEqual(t, v.ev, y)
	}
}

func TestVal_Slice(t *testing.T) {
	list := []V{
		{v: true, ev: []interface{}{}, we: true},
		{v: []interface{}{"1", "2"}, ev: []interface{}{"1", "2"}},
	}

	for i := 0; i < len(list); i++ {
		v := list[i]
		y, e := New(v.v).SliceE()
		t.Logf("%v %v", y, e)
		testx.AssertEqual(t, v.we, e != nil)
		testx.AssertEqual(t, v.ev, y)
	}
}

func TestVal_SliceString(t *testing.T) {
	list := []V{
		{v: true, ev: []string{"true"}},
		{v: []interface{}{"1", "2"}, ev: []string{"1", "2"}},
		{v: []bool{true, false}, ev: []string{"true", "false"}},
		{v: []float64{12.34, 67.89}, ev: []string{"12.34", "67.89"}},
		{v: []error{fmt.Errorf("1"), fmt.Errorf("2")}, ev: []string{"1", "2"}},
		{v: "a b c", ev: []string{"a", "b", "c"}},
	}

	for i := 0; i < len(list); i++ {
		v := list[i]
		y, e := New(v.v).SliceStringE()
		t.Logf("%v %v", y, e)
		testx.AssertEqual(t, v.we, e != nil)
		testx.AssertEqual(t, v.ev, y)
	}
}
