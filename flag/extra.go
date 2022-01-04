package flag

import (
	"strconv"
	"strings"
)

const sep = ","

var overwrite = map[interface{}]struct{}{}

// -- int slice Value
type intSliceValue []int

func newIntSliceValue(val []int, p *[]int) *intSliceValue {
	*p = val
	return (*intSliceValue)(p)
}

func (is *intSliceValue) Set(s string) error {
	if _, ok := overwrite[is]; !ok {
		overwrite[is] = struct{}{}
		*is = []int{}
	}
	v, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		*is = append(*is, int(v))
	}
	return numError(err)
}

func (is *intSliceValue) Get() interface{} {
	return []int(*is)
}

func (is *intSliceValue) String() string {
	sb := strings.Builder{}
	for i := 0; i < len(*is); i++ {
		if i > 0 {
			sb.WriteString(sep)
		}
		sb.WriteString(strconv.FormatInt(int64((*is)[i]), 10))
	}
	return sb.String()
}

func (f *FlagSet) IntSliceVar(p *[]int, name string, value []int, usage string) {
	f.Var(newIntSliceValue(value, p), name, usage)
}

func IntSliceVar(p *[]int, name string, value []int, usage string) {
	CommandLine.Var(newIntSliceValue(value, p), name, usage)
}

func (f *FlagSet) IntSlice(name string, value []int, usage string) *[]int {
	p := new([]int)
	f.IntSliceVar(p, name, value, usage)
	return p
}

func IntSlice(name string, value []int, usage string) *[]int {
	return CommandLine.IntSlice(name, value, usage)
}

// -- int64 slice Value
type int64SliceValue []int64

func newInt64SliceValue(val []int64, p *[]int64) *int64SliceValue {
	*p = val
	return (*int64SliceValue)(p)
}

func (is *int64SliceValue) Set(s string) error {
	if _, ok := overwrite[is]; !ok {
		overwrite[is] = struct{}{}
		*is = []int64{}
	}
	v, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		*is = append(*is, v)
	}
	return numError(err)
}

func (is *int64SliceValue) Get() interface{} {
	return []int64(*is)
}

func (is *int64SliceValue) String() string {
	sb := strings.Builder{}
	for i := 0; i < len(*is); i++ {
		if i > 0 {
			sb.WriteString(sep)
		}
		sb.WriteString(strconv.FormatInt((*is)[i], 10))
	}
	return sb.String()
}

func (f *FlagSet) Int64SliceVar(p *[]int64, name string, value []int64, usage string) {
	f.Var(newInt64SliceValue(value, p), name, usage)
}

func Int64SliceVar(p *[]int64, name string, value []int64, usage string) {
	CommandLine.Var(newInt64SliceValue(value, p), name, usage)
}

func (f *FlagSet) Int64Slice(name string, value []int64, usage string) *[]int64 {
	p := new([]int64)
	f.Int64SliceVar(p, name, value, usage)
	return p
}

func Int64Slice(name string, value []int64, usage string) *[]int64 {
	return CommandLine.Int64Slice(name, value, usage)
}

// -- string slice Value
type stringSliceValue []string

func newStringSliceValue(val []string, p *[]string) *stringSliceValue {
	*p = val
	return (*stringSliceValue)(p)
}

func (ss *stringSliceValue) Set(val string) error {
	if _, ok := overwrite[ss]; !ok {
		overwrite[ss] = struct{}{}
		*ss = []string{}
	}
	*ss = append(*ss, val)
	return nil
}

func (ss *stringSliceValue) Get() interface{} {
	return []string(*ss)
}

func (ss *stringSliceValue) String() string {
	sb := strings.Builder{}
	for i := 0; i < len(*ss); i++ {
		if i > 0 {
			sb.WriteString(sep)
		}
		sb.WriteString((*ss)[i])
	}
	return sb.String()
}

func (f *FlagSet) StringSliceVar(p *[]string, name string, value []string, usage string) {
	f.Var(newStringSliceValue(value, p), name, usage)
}

func StringSliceVar(p *[]string, name string, value []string, usage string) {
	CommandLine.Var(newStringSliceValue(value, p), name, usage)
}

func (f *FlagSet) StringSlice(name string, value []string, usage string) *[]string {
	p := new([]string)
	f.StringSliceVar(p, name, value, usage)
	return p
}

func StringSlice(name string, value []string, usage string) *[]string {
	return CommandLine.StringSlice(name, value, usage)
}

// -- float64 slice Value
type float64SliceValue []float64

func newFloat64SliceValue(val []float64, p *[]float64) *float64SliceValue {
	*p = val
	return (*float64SliceValue)(p)
}

func (fs *float64SliceValue) Set(s string) error {
	if _, ok := overwrite[fs]; !ok {
		overwrite[fs] = struct{}{}
		*fs = []float64{}
	}
	v, err := strconv.ParseFloat(s, 64)
	if err == nil {
		*fs = append(*fs, v)
	}
	return numError(err)
}

func (fs *float64SliceValue) Get() interface{} {
	return []float64(*fs)
}

func (fs *float64SliceValue) String() string {
	sb := strings.Builder{}
	for i := 0; i < len(*fs); i++ {
		if i > 0 {
			sb.WriteString(sep)
		}
		sb.WriteString(strconv.FormatFloat((*fs)[i], 'f', -1, 64))
	}
	return sb.String()
}

func (f *FlagSet) Float64SliceVar(p *[]float64, name string, value []float64, usage string) {
	f.Var(newFloat64SliceValue(value, p), name, usage)
}

func Float64SliceVar(p *[]float64, name string, value []float64, usage string) {
	CommandLine.Var(newFloat64SliceValue(value, p), name, usage)
}

func (f *FlagSet) Float64Slice(name string, value []float64, usage string) *[]float64 {
	p := new([]float64)
	f.Float64SliceVar(p, name, value, usage)
	return p
}

func Float64Slice(name string, value []float64, usage string) *[]float64 {
	return CommandLine.Float64Slice(name, value, usage)
}
