package slicex

import (
	"reflect"

	"github.com/go-sdk/lib/val"
)

func ToInterface(x interface{}) []interface{} {
	v := reflect.ValueOf(x)
	if v.Kind() != reflect.Slice || v.IsNil() {
		return nil
	}
	vs := make([]interface{}, v.Len())
	for i := 0; i < v.Len(); i++ {
		vs[i] = v.Index(i).Interface()
	}
	return vs
}

func ToString(vs []interface{}) []string {
	ss := make([]string, len(vs))
	for i := 0; i < len(vs); i++ {
		ss[i] = val.New(vs[i]).String()
	}
	return ss
}

func ToInt64(vs []interface{}) []int64 {
	is := make([]int64, len(vs))
	for i := 0; i < len(vs); i++ {
		is[i] = val.New(vs[i]).Int64()
	}
	return is
}
