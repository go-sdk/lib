package val

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type Value interface {
	BoolE() (bool, error)
	Int64E() (int64, error)
	Float64E() (float64, error)
	StringE() (string, error)
	DurationE() (time.Duration, error)
	SliceE() ([]interface{}, error)
	SliceStringE() ([]string, error)

	Bool() bool
	Int64() int64
	Float64() float64
	String() string
	Duration() time.Duration
	Slice() []interface{}
	SliceString() []string

	BoolD(d bool) bool
	Int64D(d int64) int64
	Float64D(d float64) float64
	StringD(d string) string
	DurationD(d time.Duration) time.Duration
	SliceD(d []interface{}) []interface{}
	SliceStringD(d []string) []string
}

func New(v interface{}) Value {
	return val{v: v}
}

type val struct {
	v interface{}
}

func (v val) BoolE() (bool, error) {
	x := indirect(v.v)

	switch y := x.(type) {
	case bool:
		return y, nil
	case uint, uint8, uint16, uint32, uint64, int, int8, int16, int32, int64, float32, float64:
		if v.Float64() != 0 {
			return true, nil
		}
		return false, nil
	case string:
		return strconv.ParseBool(y)
	default:
		return false, fmt.Errorf("unable to cast %#v of type %T to bool", x, x)
	}
}

func (v val) Int64E() (int64, error) {
	x := indirect(v.v)

	switch y := x.(type) {
	case bool:
		if y {
			return 1, nil
		}
		return 0, nil
	case uint, uint8, uint16, uint32, uint64, int, int8, int16, int32, int64, float32, float64:
		return int64(v.Float64()), nil
	case string:
		i, err := strconv.ParseInt(y, 10, 64)
		if err == nil {
			return i, nil
		}
		return 0, fmt.Errorf("unable to cast %#v of type %T to int64", x, x)
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to int64", x, x)
	}
}

func (v val) Float64E() (float64, error) {
	x := indirect(v.v)

	switch y := x.(type) {
	case bool:
		if y {
			return 1, nil
		}
		return 0, nil
	case uint:
		return float64(y), nil
	case uint8:
		return float64(y), nil
	case uint16:
		return float64(y), nil
	case uint32:
		return float64(y), nil
	case uint64:
		return float64(y), nil
	case int:
		return float64(y), nil
	case int8:
		return float64(y), nil
	case int16:
		return float64(y), nil
	case int32:
		return float64(y), nil
	case int64:
		return float64(y), nil
	case float32:
		return float64(y), nil
	case float64:
		return y, nil
	case string:
		i, err := strconv.ParseFloat(y, 64)
		if err == nil {
			return i, nil
		}
		return 0, fmt.Errorf("unable to cast %#v of type %T to float64", x, x)
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to float64", x, x)
	}
}

func (v val) StringE() (string, error) {
	x := indirectToStringerOrError(v.v)

	switch y := x.(type) {
	case bool:
		return strconv.FormatBool(y), nil
	case uint:
		return strconv.FormatUint(uint64(y), 10), nil
	case uint8:
		return strconv.FormatUint(uint64(y), 10), nil
	case uint16:
		return strconv.FormatUint(uint64(y), 10), nil
	case uint32:
		return strconv.FormatUint(uint64(y), 10), nil
	case uint64:
		return strconv.FormatUint(y, 10), nil
	case int:
		return strconv.Itoa(y), nil
	case int8:
		return strconv.FormatInt(int64(y), 10), nil
	case int16:
		return strconv.FormatInt(int64(y), 10), nil
	case int32:
		return strconv.FormatInt(int64(y), 10), nil
	case int64:
		return strconv.FormatInt(y, 10), nil
	case float32:
		return strconv.FormatFloat(float64(y), 'f', -1, 32), nil
	case float64:
		return strconv.FormatFloat(y, 'f', -1, 64), nil
	case string:
		return y, nil
	case []byte:
		return string(y), nil
	case fmt.Stringer:
		return y.String(), nil
	case error:
		return y.Error(), nil
	default:
		return "", fmt.Errorf("unable to cast %#v of type %T to string", x, x)
	}
}

func (v val) DurationE() (time.Duration, error) {
	x := indirect(v.v)

	switch y := x.(type) {
	case time.Duration:
		return y, nil
	case uint, uint8, uint16, uint32, uint64, int, int8, int16, int32, int64, float32, float64:
		return time.Duration(v.Float64()), nil
	case string:
		if strings.ContainsAny(y, "nsuµmh") {
			return time.ParseDuration(y)
		}
		return time.ParseDuration(y + "ns")
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to Duration", x, x)
	}
}

func (v val) SliceE() ([]interface{}, error) {
	x := v.v

	switch y := v.v.(type) {
	case []interface{}:
		return y, nil
	case []map[string]interface{}:
		vs := make([]interface{}, len(y))
		for i, u := range y {
			vs[i] = u
		}
		return vs, nil
	default:
		return []interface{}{}, fmt.Errorf("unable to cast %#v of type %T to []interface{}", x, x)
	}
}

func (v val) SliceStringE() ([]string, error) {
	x := v.v

	switch y := v.v.(type) {
	case []bool, []uint, []uint8, []uint16, []uint32, []uint64, []int, []int8, []int16, []int32, []int64, []float32, []float64, []interface{}, []error:
		yv := reflect.ValueOf(y)
		vs := make([]string, yv.Len())
		for i := 0; i < yv.Len(); i++ {
			vs[i] = New(yv.Index(i).Interface()).String()
		}
		return vs, nil
	case []string:
		return y, nil
	case string:
		return strings.Fields(y), nil
	case interface{}:
		s, err := New(y).StringE()
		if err != nil {
			return []string{}, fmt.Errorf("unable to cast %#v of type %T to []string", x, x)
		}
		return []string{s}, nil
	default:
		return []string{}, fmt.Errorf("unable to cast %#v of type %T to []string", x, x)
	}
}

func (v val) Bool() bool {
	x, _ := v.BoolE()
	return x
}

func (v val) Int64() int64 {
	x, _ := v.Int64E()
	return x
}

func (v val) Float64() float64 {
	x, _ := v.Float64E()
	return x
}

func (v val) String() string {
	x, _ := v.StringE()
	return x
}

func (v val) Duration() time.Duration {
	x, _ := v.DurationE()
	return x
}

func (v val) Slice() []interface{} {
	x, _ := v.SliceE()
	return x
}

func (v val) SliceString() []string {
	x, _ := v.SliceStringE()
	return x
}

func (v val) BoolD(d bool) bool {
	x, err := v.BoolE()
	if err != nil {
		return d
	}
	return x
}

func (v val) Int64D(d int64) int64 {
	x, err := v.Int64E()
	if err != nil {
		return d
	}
	return x
}

func (v val) Float64D(d float64) float64 {
	x, err := v.Float64E()
	if err != nil {
		return d
	}
	return x
}

func (v val) StringD(d string) string {
	x, err := v.StringE()
	if err != nil {
		return d
	}
	return x
}

func (v val) DurationD(d time.Duration) time.Duration {
	x, err := v.DurationE()
	if err != nil {
		return d
	}
	return x
}

func (v val) SliceD(d []interface{}) []interface{} {
	x, err := v.SliceE()
	if err != nil {
		return d
	}
	return x
}

func (v val) SliceStringD(d []string) []string {
	x, err := v.SliceStringE()
	if err != nil {
		return d
	}
	return x
}
